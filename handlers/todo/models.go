package handlers

import (
	"database/sql"
	"time"
)

// Project represents a todo project/category
type Project struct {
	ID          int       `json:"id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	TodoCount   int       `json:"todo_count,omitempty"`
}

// Todo represents a single todo item
type Todo struct {
	ID          int                `json:"id"`
	ProjectID   int                `json:"project_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
	Priority    string             `json:"priority"`
	DueDate     *time.Time         `json:"due_date,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Author      string             `json:"author"`
	Project     *Project           `json:"project,omitempty"`
	Comments    []Comment          `json:"comments,omitempty"`
}

// Comment represents a comment on a todo item
type Comment struct {
	ID        int       `json:"id"`
	TodoID    int       `json:"todo_id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoStats represents todo statistics
type TodoStats struct {
	Total       int `json:"total"`
	Completed   int `json:"completed"`
	Pending     int `json:"pending"`
	HighPriority int `json:"high_priority"`
	Overdue     int `json:"overdue"`
}

// CreateTodoRequest represents the request to create a new todo
type CreateTodoRequest struct {
	ProjectID   int        `json:"project_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	Author      string     `json:"author"`
}

// UpdateTodoRequest represents the request to update a todo
type UpdateTodoRequest struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Completed   *bool      `json:"completed,omitempty"`
	Priority    *string    `json:"priority,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

// CreateProjectRequest represents the request to create a new project
type CreateProjectRequest struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

// CreateCommentRequest represents the request to create a new comment
type CreateCommentRequest struct {
	TodoID  int    `json:"todo_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// TodoFilter represents filtering options for todos
type TodoFilter struct {
	ProjectID *int    `json:"project_id,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
	Priority  *string `json:"priority,omitempty"`
	Author    *string `json:"author,omitempty"`
	Search    *string `json:"search,omitempty"`
	Limit     int     `json:"limit,omitempty"`
	Offset    int     `json:"offset,omitempty"`
}

// Helper function to scan a nullable time
func scanNullableTime(src interface{}) *time.Time {
	if src == nil {
		return nil
	}
	
	switch v := src.(type) {
	case time.Time:
		return &v
	case string:
		if t, err := time.Parse("2006-01-02 15:04:05", v); err == nil {
			return &t
		}
	}
	return nil
}

// Helper function to scan a nullable string to time
func scanNullableTimeString(ns sql.NullString) *time.Time {
	if !ns.Valid {
		return nil
	}
	if t, err := time.Parse("2006-01-02 15:04:05", ns.String); err == nil {
		return &t
	}
	return nil
}