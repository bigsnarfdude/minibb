// Types for TODO list application

export interface Project {
  id: number;
  slug: string;
  name: string;
  description: string;
  color: string;
  created_at: string;
  todo_count?: number;
}

export interface Todo {
  id: number;
  project_id: number;
  title: string;
  description: string;
  completed: boolean;
  priority: 'low' | 'medium' | 'high';
  due_date?: string;
  created_at: string;
  updated_at: string;
  author: string;
  project?: Project;
  comments?: Comment[];
}

export interface Comment {
  id: number;
  todo_id: number;
  content: string;
  author: string;
  created_at: string;
}

export interface TodoStats {
  total: number;
  completed: number;
  pending: number;
  high_priority: number;
  overdue: number;
}

export interface CreateTodoRequest {
  project_id: number;
  title: string;
  description: string;
  priority: 'low' | 'medium' | 'high';
  due_date?: string;
  author: string;
}

export interface UpdateTodoRequest {
  title?: string;
  description?: string;
  completed?: boolean;
  priority?: 'low' | 'medium' | 'high';
  due_date?: string;
}

export interface CreateProjectRequest {
  slug: string;
  name: string;
  description: string;
  color: string;
}

export interface CreateCommentRequest {
  todo_id: number;
  content: string;
  author: string;
}

export interface TodoFilter {
  project_id?: number;
  completed?: boolean;
  priority?: 'low' | 'medium' | 'high';
  author?: string;
  search?: string;
  limit?: number;
  offset?: number;
}

export const PRIORITY_COLORS = {
  high: 'bg-red-100 text-red-800',
  medium: 'bg-yellow-100 text-yellow-800',
  low: 'bg-green-100 text-green-800',
} as const;

export const PRIORITY_LABELS = {
  high: 'High Priority',
  medium: 'Medium Priority',
  low: 'Low Priority',
} as const;