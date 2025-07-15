import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { fetchTodos, completeTodo, uncompleteTodo, deleteTodo } from '../api';
import { Todo, TodoFilter, PRIORITY_COLORS, PRIORITY_LABELS } from '../types';

interface TodoListProps {
  filter?: TodoFilter;
  onEdit?: (todo: Todo) => void;
  onView?: (todo: Todo) => void;
}

export function TodoList({ filter, onEdit, onView }: TodoListProps) {
  const queryClient = useQueryClient();
  const [localFilter, setLocalFilter] = useState<TodoFilter>(filter || {});

  const { data: todos, isLoading, error } = useQuery({
    queryKey: ['todos', localFilter],
    queryFn: () => fetchTodos(localFilter),
  });

  const toggleComplete = useMutation({
    mutationFn: ({ id, completed }: { id: number; completed: boolean }) => 
      completed ? uncompleteTodo(id) : completeTodo(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
      queryClient.invalidateQueries({ queryKey: ['stats'] });
    },
  });

  const deleteTodoMutation = useMutation({
    mutationFn: deleteTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
      queryClient.invalidateQueries({ queryKey: ['stats'] });
    },
  });

  const handleToggleComplete = (todo: Todo) => {
    toggleComplete.mutate({ id: todo.id, completed: todo.completed });
  };

  const handleDelete = (todo: Todo) => {
    if (confirm(`Are you sure you want to delete "${todo.title}"?`)) {
      deleteTodoMutation.mutate(todo.id);
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString();
  };

  const isOverdue = (todo: Todo) => {
    if (!todo.due_date || todo.completed) return false;
    return new Date(todo.due_date) < new Date();
  };

  if (isLoading) {
    return (
      <div className="flex justify-center items-center py-8">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-red-600 text-center py-8">
        Error loading todos: {error.message}
      </div>
    );
  }

  if (!todos || todos.length === 0) {
    return (
      <div className="text-gray-500 text-center py-8">
        No todos found. Create your first todo to get started!
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {/* Filter Controls */}
      <div className="flex flex-wrap gap-4 p-4 bg-gray-50 rounded-lg">
        <div className="flex items-center gap-2">
          <label className="text-sm font-medium">Status:</label>
          <select
            value={localFilter.completed?.toString() || ''}
            onChange={(e) => {
              const value = e.target.value;
              setLocalFilter(prev => ({
                ...prev,
                completed: value === '' ? undefined : value === 'true'
              }));
            }}
            className="px-3 py-1 border rounded-md text-sm"
          >
            <option value="">All</option>
            <option value="false">Pending</option>
            <option value="true">Completed</option>
          </select>
        </div>

        <div className="flex items-center gap-2">
          <label className="text-sm font-medium">Priority:</label>
          <select
            value={localFilter.priority || ''}
            onChange={(e) => {
              const value = e.target.value;
              setLocalFilter(prev => ({
                ...prev,
                priority: value === '' ? undefined : value as 'low' | 'medium' | 'high'
              }));
            }}
            className="px-3 py-1 border rounded-md text-sm"
          >
            <option value="">All</option>
            <option value="high">High</option>
            <option value="medium">Medium</option>
            <option value="low">Low</option>
          </select>
        </div>

        <div className="flex items-center gap-2">
          <label className="text-sm font-medium">Search:</label>
          <input
            type="text"
            value={localFilter.search || ''}
            onChange={(e) => {
              setLocalFilter(prev => ({
                ...prev,
                search: e.target.value || undefined
              }));
            }}
            placeholder="Search todos..."
            className="px-3 py-1 border rounded-md text-sm"
          />
        </div>
      </div>

      {/* Todo Items */}
      <div className="space-y-3">
        {todos.map((todo) => (
          <div
            key={todo.id}
            className={`p-4 border rounded-lg transition-all duration-200 ${
              todo.completed
                ? 'bg-gray-50 border-gray-200 opacity-75'
                : 'bg-white border-gray-300 hover:shadow-md'
            } ${isOverdue(todo) ? 'border-red-300 bg-red-50' : ''}`}
          >
            <div className="flex items-start gap-3">
              {/* Checkbox */}
              <input
                type="checkbox"
                checked={todo.completed}
                onChange={() => handleToggleComplete(todo)}
                className="mt-1 w-5 h-5 text-blue-600 rounded focus:ring-blue-500"
                disabled={toggleComplete.isPending}
              />

              {/* Main Content */}
              <div className="flex-1 min-w-0">
                <div className="flex items-start justify-between">
                  <div className="flex-1">
                    <h3
                      className={`text-lg font-medium ${
                        todo.completed ? 'line-through text-gray-500' : 'text-gray-900'
                      }`}
                    >
                      {todo.title}
                    </h3>
                    {todo.description && (
                      <p className={`mt-1 text-sm ${
                        todo.completed ? 'text-gray-400' : 'text-gray-600'
                      }`}>
                        {todo.description}
                      </p>
                    )}
                  </div>

                  {/* Actions */}
                  <div className="flex items-center gap-2 ml-4">
                    {onView && (
                      <button
                        onClick={() => onView(todo)}
                        className="text-blue-600 hover:text-blue-800 text-sm"
                      >
                        View
                      </button>
                    )}
                    {onEdit && (
                      <button
                        onClick={() => onEdit(todo)}
                        className="text-gray-600 hover:text-gray-800 text-sm"
                      >
                        Edit
                      </button>
                    )}
                    <button
                      onClick={() => handleDelete(todo)}
                      className="text-red-600 hover:text-red-800 text-sm"
                      disabled={deleteTodoMutation.isPending}
                    >
                      Delete
                    </button>
                  </div>
                </div>

                {/* Meta Information */}
                <div className="flex items-center gap-4 mt-3">
                  {/* Priority Badge */}
                  <span
                    className={`px-2 py-1 rounded-full text-xs font-medium ${
                      PRIORITY_COLORS[todo.priority]
                    }`}
                  >
                    {PRIORITY_LABELS[todo.priority]}
                  </span>

                  {/* Project */}
                  {todo.project && (
                    <span className="text-xs text-gray-500">
                      Project: {todo.project.name}
                    </span>
                  )}

                  {/* Due Date */}
                  {todo.due_date && (
                    <span
                      className={`text-xs ${
                        isOverdue(todo) ? 'text-red-600 font-medium' : 'text-gray-500'
                      }`}
                    >
                      Due: {formatDate(todo.due_date)}
                      {isOverdue(todo) && ' (Overdue)'}
                    </span>
                  )}

                  {/* Author */}
                  <span className="text-xs text-gray-500">
                    by {todo.author}
                  </span>

                  {/* Created Date */}
                  <span className="text-xs text-gray-400">
                    Created: {formatDate(todo.created_at)}
                  </span>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Load More Button */}
      {todos.length >= (localFilter.limit || 50) && (
        <div className="text-center pt-4">
          <button
            onClick={() => {
              setLocalFilter(prev => ({
                ...prev,
                offset: (prev.offset || 0) + (prev.limit || 50)
              }));
            }}
            className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
          >
            Load More
          </button>
        </div>
      )}
    </div>
  );
}