package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

// TodoAPI holds the database connection and provides API handlers
type TodoAPI struct {
	db *sql.DB
}

// NewTodoAPI creates a new TodoAPI instance
func NewTodoAPI(db *sql.DB) *TodoAPI {
	return &TodoAPI{db: db}
}

// Routes sets up the API routes for todo application
func (api *TodoAPI) Routes() chi.Router {
	r := chi.NewRouter()

	// Project routes
	r.Get("/projects", api.ListProjects)
	r.Post("/projects", api.CreateProject)
	r.Get("/projects/{slug}", api.GetProject)
	r.Put("/projects/{slug}", api.UpdateProject)
	r.Delete("/projects/{slug}", api.DeleteProject)

	// Todo routes
	r.Get("/todos", api.ListTodos)
	r.Post("/todos", api.CreateTodo)
	r.Get("/todos/{id}", api.GetTodo)
	r.Put("/todos/{id}", api.UpdateTodo)
	r.Delete("/todos/{id}", api.DeleteTodo)
	r.Post("/todos/{id}/complete", api.CompleteTodo)
	r.Post("/todos/{id}/uncomplete", api.UncompleteTodo)

	// Comment routes
	r.Get("/todos/{id}/comments", api.ListComments)
	r.Post("/todos/{id}/comments", api.CreateComment)
	r.Delete("/comments/{id}", api.DeleteComment)

	// Stats routes
	r.Get("/stats", api.GetStats)
	r.Get("/projects/{slug}/stats", api.GetProjectStats)

	return r
}

// ListProjects returns all projects with todo counts
func (api *TodoAPI) ListProjects(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT p.id, p.slug, p.name, p.description, p.color, p.created_at,
			   COUNT(t.id) as todo_count
		FROM projects p
		LEFT JOIN todos t ON p.id = t.project_id
		GROUP BY p.id, p.slug, p.name, p.description, p.color, p.created_at
		ORDER BY p.created_at DESC
	`

	rows, err := api.db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		err := rows.Scan(&p.ID, &p.Slug, &p.Name, &p.Description, &p.Color, &p.CreatedAt, &p.TodoCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		projects = append(projects, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// CreateProject creates a new project
func (api *TodoAPI) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Slug == "" || req.Name == "" {
		http.Error(w, "Slug and name are required", http.StatusBadRequest)
		return
	}

	// Set default color if not provided
	if req.Color == "" {
		req.Color = "#3b82f6"
	}

	query := `
		INSERT INTO projects (slug, name, description, color)
		VALUES (?, ?, ?, ?)
	`

	result, err := api.db.Exec(query, req.Slug, req.Name, req.Description, req.Color)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			http.Error(w, "Project with this slug already exists", http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	project := Project{
		ID:          int(id),
		Slug:        req.Slug,
		Name:        req.Name,
		Description: req.Description,
		Color:       req.Color,
		CreatedAt:   time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

// GetProject returns a specific project by slug
func (api *TodoAPI) GetProject(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	query := `
		SELECT id, slug, name, description, color, created_at
		FROM projects
		WHERE slug = ?
	`

	var p Project
	err := api.db.QueryRow(query, slug).Scan(&p.ID, &p.Slug, &p.Name, &p.Description, &p.Color, &p.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// ListTodos returns todos with filtering and pagination
func (api *TodoAPI) ListTodos(w http.ResponseWriter, r *http.Request) {
	filter := TodoFilter{
		Limit:  50, // Default limit
		Offset: 0,  // Default offset
	}

	// Parse query parameters
	if projectID := r.URL.Query().Get("project_id"); projectID != "" {
		if id, err := strconv.Atoi(projectID); err == nil {
			filter.ProjectID = &id
		}
	}

	if completed := r.URL.Query().Get("completed"); completed != "" {
		if comp, err := strconv.ParseBool(completed); err == nil {
			filter.Completed = &comp
		}
	}

	if priority := r.URL.Query().Get("priority"); priority != "" {
		filter.Priority = &priority
	}

	if search := r.URL.Query().Get("search"); search != "" {
		filter.Search = &search
	}

	if limit := r.URL.Query().Get("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= 100 {
			filter.Limit = l
		}
	}

	if offset := r.URL.Query().Get("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil && o >= 0 {
			filter.Offset = o
		}
	}

	// Build query
	query := `
		SELECT t.id, t.project_id, t.title, t.description, t.completed, t.priority,
			   t.due_date, t.created_at, t.updated_at, t.author,
			   p.slug, p.name, p.color
		FROM todos t
		LEFT JOIN projects p ON t.project_id = p.id
		WHERE 1=1
	`

	var args []interface{}
	argIndex := 0

	if filter.ProjectID != nil {
		query += " AND t.project_id = ?"
		args = append(args, *filter.ProjectID)
		argIndex++
	}

	if filter.Completed != nil {
		query += " AND t.completed = ?"
		args = append(args, *filter.Completed)
		argIndex++
	}

	if filter.Priority != nil {
		query += " AND t.priority = ?"
		args = append(args, *filter.Priority)
		argIndex++
	}

	if filter.Search != nil {
		query += " AND (t.title LIKE ? OR t.description LIKE ?)"
		searchTerm := "%" + *filter.Search + "%"
		args = append(args, searchTerm, searchTerm)
		argIndex += 2
	}

	query += " ORDER BY t.completed ASC, t.priority DESC, t.created_at DESC LIMIT ? OFFSET ?"
	args = append(args, filter.Limit, filter.Offset)

	rows, err := api.db.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		var dueDate sql.NullString
		var project Project

		err := rows.Scan(
			&t.ID, &t.ProjectID, &t.Title, &t.Description, &t.Completed, &t.Priority,
			&dueDate, &t.CreatedAt, &t.UpdatedAt, &t.Author,
			&project.Slug, &project.Name, &project.Color,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.DueDate = scanNullableTimeString(dueDate)
		t.Project = &project

		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo creates a new todo
func (api *TodoAPI) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Title == "" || req.Author == "" {
		http.Error(w, "Title and author are required", http.StatusBadRequest)
		return
	}

	// Set default priority if not provided
	if req.Priority == "" {
		req.Priority = "medium"
	}

	query := `
		INSERT INTO todos (project_id, title, description, priority, due_date, author)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := api.db.Exec(query, req.ProjectID, req.Title, req.Description, req.Priority, req.DueDate, req.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	todo := Todo{
		ID:          int(id),
		ProjectID:   req.ProjectID,
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		Author:      req.Author,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// GetTodo returns a specific todo with comments
func (api *TodoAPI) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Get todo details
	query := `
		SELECT t.id, t.project_id, t.title, t.description, t.completed, t.priority,
			   t.due_date, t.created_at, t.updated_at, t.author,
			   p.slug, p.name, p.color
		FROM todos t
		LEFT JOIN projects p ON t.project_id = p.id
		WHERE t.id = ?
	`

	var t Todo
	var dueDate sql.NullString
	var project Project

	err := api.db.QueryRow(query, id).Scan(
		&t.ID, &t.ProjectID, &t.Title, &t.Description, &t.Completed, &t.Priority,
		&dueDate, &t.CreatedAt, &t.UpdatedAt, &t.Author,
		&project.Slug, &project.Name, &project.Color,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.DueDate = scanNullableTimeString(dueDate)
	t.Project = &project

	// Get comments
	commentQuery := `
		SELECT id, todo_id, content, author, created_at
		FROM comments
		WHERE todo_id = ?
		ORDER BY created_at ASC
	`

	commentRows, err := api.db.Query(commentQuery, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer commentRows.Close()

	var comments []Comment
	for commentRows.Next() {
		var c Comment
		err := commentRows.Scan(&c.ID, &c.TodoID, &c.Content, &c.Author, &c.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		comments = append(comments, c)
	}

	t.Comments = comments

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

// UpdateTodo updates a todo
func (api *TodoAPI) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Build update query dynamically
	var setParts []string
	var args []interface{}

	if req.Title != nil {
		setParts = append(setParts, "title = ?")
		args = append(args, *req.Title)
	}

	if req.Description != nil {
		setParts = append(setParts, "description = ?")
		args = append(args, *req.Description)
	}

	if req.Completed != nil {
		setParts = append(setParts, "completed = ?")
		args = append(args, *req.Completed)
	}

	if req.Priority != nil {
		setParts = append(setParts, "priority = ?")
		args = append(args, *req.Priority)
	}

	if req.DueDate != nil {
		setParts = append(setParts, "due_date = ?")
		args = append(args, *req.DueDate)
	}

	if len(setParts) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("UPDATE todos SET %s WHERE id = ?", strings.Join(setParts, ", "))
	args = append(args, id)

	_, err := api.db.Exec(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CompleteTodo marks a todo as completed
func (api *TodoAPI) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query := "UPDATE todos SET completed = 1 WHERE id = ?"
	_, err := api.db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UncompleteTodo marks a todo as not completed
func (api *TodoAPI) UncompleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query := "UPDATE todos SET completed = 0 WHERE id = ?"
	_, err := api.db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteTodo deletes a todo
func (api *TodoAPI) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query := "DELETE FROM todos WHERE id = ?"
	_, err := api.db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateComment creates a comment on a todo
func (api *TodoAPI) CreateComment(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "id")

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Content == "" || req.Author == "" {
		http.Error(w, "Content and author are required", http.StatusBadRequest)
		return
	}

	// Convert todoID to int
	todoIDInt, err := strconv.Atoi(todoID)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO comments (todo_id, content, author)
		VALUES (?, ?, ?)
	`

	result, err := api.db.Exec(query, todoIDInt, req.Content, req.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	comment := Comment{
		ID:        int(id),
		TodoID:    todoIDInt,
		Content:   req.Content,
		Author:    req.Author,
		CreatedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// GetStats returns overall todo statistics
func (api *TodoAPI) GetStats(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT 
			COUNT(*) as total,
			SUM(CASE WHEN completed = 1 THEN 1 ELSE 0 END) as completed,
			SUM(CASE WHEN completed = 0 THEN 1 ELSE 0 END) as pending,
			SUM(CASE WHEN priority = 'high' AND completed = 0 THEN 1 ELSE 0 END) as high_priority,
			SUM(CASE WHEN due_date < datetime('now') AND completed = 0 THEN 1 ELSE 0 END) as overdue
		FROM todos
	`

	var stats TodoStats
	err := api.db.QueryRow(query).Scan(&stats.Total, &stats.Completed, &stats.Pending, &stats.HighPriority, &stats.Overdue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}