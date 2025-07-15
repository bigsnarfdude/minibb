// API functions for TODO list application

import { 
  Project, 
  Todo, 
  Comment, 
  TodoStats, 
  CreateTodoRequest, 
  UpdateTodoRequest, 
  CreateProjectRequest, 
  CreateCommentRequest,
  TodoFilter 
} from './types';

const API_BASE = '/api';

// Helper function to handle API responses
async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || `HTTP ${response.status}`);
  }
  return response.json();
}

// Project API functions
export async function fetchProjects(): Promise<Project[]> {
  const response = await fetch(`${API_BASE}/projects`);
  return handleResponse<Project[]>(response);
}

export async function createProject(project: CreateProjectRequest): Promise<Project> {
  const response = await fetch(`${API_BASE}/projects`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(project),
  });
  return handleResponse<Project>(response);
}

export async function fetchProject(slug: string): Promise<Project> {
  const response = await fetch(`${API_BASE}/projects/${slug}`);
  return handleResponse<Project>(response);
}

export async function updateProject(slug: string, project: Partial<CreateProjectRequest>): Promise<void> {
  const response = await fetch(`${API_BASE}/projects/${slug}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(project),
  });
  return handleResponse<void>(response);
}

export async function deleteProject(slug: string): Promise<void> {
  const response = await fetch(`${API_BASE}/projects/${slug}`, {
    method: 'DELETE',
  });
  return handleResponse<void>(response);
}

// Todo API functions
export async function fetchTodos(filter?: TodoFilter): Promise<Todo[]> {
  const params = new URLSearchParams();
  
  if (filter?.project_id) params.append('project_id', filter.project_id.toString());
  if (filter?.completed !== undefined) params.append('completed', filter.completed.toString());
  if (filter?.priority) params.append('priority', filter.priority);
  if (filter?.author) params.append('author', filter.author);
  if (filter?.search) params.append('search', filter.search);
  if (filter?.limit) params.append('limit', filter.limit.toString());
  if (filter?.offset) params.append('offset', filter.offset.toString());

  const response = await fetch(`${API_BASE}/todos?${params}`);
  return handleResponse<Todo[]>(response);
}

export async function createTodo(todo: CreateTodoRequest): Promise<Todo> {
  const response = await fetch(`${API_BASE}/todos`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(todo),
  });
  return handleResponse<Todo>(response);
}

export async function fetchTodo(id: number): Promise<Todo> {
  const response = await fetch(`${API_BASE}/todos/${id}`);
  return handleResponse<Todo>(response);
}

export async function updateTodo(id: number, todo: UpdateTodoRequest): Promise<void> {
  const response = await fetch(`${API_BASE}/todos/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(todo),
  });
  return handleResponse<void>(response);
}

export async function deleteTodo(id: number): Promise<void> {
  const response = await fetch(`${API_BASE}/todos/${id}`, {
    method: 'DELETE',
  });
  return handleResponse<void>(response);
}

export async function completeTodo(id: number): Promise<void> {
  const response = await fetch(`${API_BASE}/todos/${id}/complete`, {
    method: 'POST',
  });
  return handleResponse<void>(response);
}

export async function uncompleteTodo(id: number): Promise<void> {
  const response = await fetch(`${API_BASE}/todos/${id}/uncomplete`, {
    method: 'POST',
  });
  return handleResponse<void>(response);
}

// Comment API functions
export async function fetchComments(todoId: number): Promise<Comment[]> {
  const response = await fetch(`${API_BASE}/todos/${todoId}/comments`);
  return handleResponse<Comment[]>(response);
}

export async function createComment(comment: CreateCommentRequest): Promise<Comment> {
  const response = await fetch(`${API_BASE}/todos/${comment.todo_id}/comments`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(comment),
  });
  return handleResponse<Comment>(response);
}

export async function deleteComment(id: number): Promise<void> {
  const response = await fetch(`${API_BASE}/comments/${id}`, {
    method: 'DELETE',
  });
  return handleResponse<void>(response);
}

// Stats API functions
export async function fetchStats(): Promise<TodoStats> {
  const response = await fetch(`${API_BASE}/stats`);
  return handleResponse<TodoStats>(response);
}

export async function fetchProjectStats(slug: string): Promise<TodoStats> {
  const response = await fetch(`${API_BASE}/projects/${slug}/stats`);
  return handleResponse<TodoStats>(response);
}