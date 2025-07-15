-- Default projects for TODO list application

INSERT INTO projects (slug, name, description, color) VALUES
('personal', 'Personal Tasks', 'Personal todos and reminders', '#10b981'),
('work', 'Work Projects', 'Work-related tasks and deadlines', '#3b82f6'),
('shopping', 'Shopping List', 'Items to buy and errands to run', '#f59e0b'),
('learning', 'Learning Goals', 'Educational tasks and skill development', '#8b5cf6');

-- Sample todos for demonstration
INSERT INTO todos (project_id, title, description, priority, author) VALUES
(1, 'Buy groceries', 'Milk, bread, eggs, and vegetables', 'medium', 'Student'),
(1, 'Exercise routine', 'Go for a 30-minute walk or gym session', 'high', 'Student'),
(2, 'Complete project proposal', 'Draft and submit the Q1 project proposal', 'high', 'Student'),
(2, 'Team meeting preparation', 'Review agenda and prepare presentation slides', 'medium', 'Student'),
(3, 'New laptop battery', 'Research and purchase replacement battery', 'low', 'Student'),
(4, 'Learn React hooks', 'Complete tutorial on useEffect and useState', 'medium', 'Student');

-- Sample comments
INSERT INTO comments (todo_id, content, author) VALUES
(1, 'Don''t forget to check for sales on organic produce', 'Student'),
(3, 'Deadline is next Friday - need to prioritize this', 'Student'),
(6, 'Found a great tutorial on YouTube - https://example.com/react-hooks', 'Student');