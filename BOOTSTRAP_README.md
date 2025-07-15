# WebApp Bootstrap Project

A comprehensive bootstrap project for teaching intro web programming, featuring three different application variants: TODO list, Forum/Bulletin Board, and Wiki system.

## 🎯 Overview

This project provides a solid foundation for learning full-stack web development using modern technologies:
- **Backend**: Go with Chi router and SQLite database
- **Frontend**: React with TypeScript, TanStack Query/Router, and Tailwind CSS
- **Development**: Hot reloading, linting, formatting, and testing tools

## 📋 Available Variants

### 1. TODO List Application (Beginner)
**Perfect for learning basic CRUD operations**

**Features:**
- Create, edit, and delete todos
- Organize todos into projects
- Set priorities and due dates
- Track completion status
- Add comments to todos
- Statistics dashboard

**Learning Objectives:**
- Database design and relationships
- REST API development
- Form handling and validation
- State management
- Basic UI components

### 2. Forum/Bulletin Board (Intermediate)
**Great for understanding user interactions and content management**

**Features:**
- Multiple discussion boards
- Create topics and replies
- Tripcode authentication
- Read status tracking
- Markdown formatting
- Pagination and filtering

**Learning Objectives:**
- User authentication concepts
- Content management systems
- Real-time updates
- Advanced UI patterns
- Security considerations

### 3. Wiki System (Advanced)
**Ideal for learning complex data relationships and versioning**

**Features:**
- Page creation and editing
- Revision history and diffs
- Cross-page linking
- Full-text search
- Namespaces/categories
- Collaborative editing

**Learning Objectives:**
- Version control systems
- Full-text search implementation
- Complex data relationships
- Advanced UI components
- Performance optimization

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Node.js 18+
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd minibb-main
   ```

2. **Choose your variant**
   ```bash
   # For TODO list application
   ./bootstrap.sh todo
   
   # For forum/bulletin board
   ./bootstrap.sh forum
   
   # For wiki system
   ./bootstrap.sh wiki
   ```

3. **Start development**
   ```bash
   make dev
   ```

4. **Open your browser**
   - Navigate to `http://localhost:5173`
   - Start coding!

## 📁 Project Structure

```
minibb-main/
├── bootstrap.sh              # Project switcher script
├── bootstrap.config.json     # Configuration and learning paths
├── Makefile                  # Development commands
├── 
├── cmd/minibb/              # Go application entry point
├── internal/                # Go backend code
│   ├── db/                  # Database connection and migrations
│   ├── handlers/            # HTTP handlers
│   ├── models/              # Data models
│   ├── server/              # Server setup
│   └── utils/               # Utility functions
├── 
├── migrations/              # Database migrations for each variant
│   ├── todo/               # TODO list schema
│   ├── forum/              # Forum schema
│   └── wiki/               # Wiki schema
├── 
├── handlers/               # API handlers for each variant
│   ├── todo/              # TODO list handlers
│   ├── forum/             # Forum handlers
│   └── wiki/              # Wiki handlers
├── 
├── web/                   # Frontend application
│   ├── src/
│   │   ├── lib/           # API clients and utilities
│   │   ├── routes/        # Route definitions
│   │   └── variants/      # Variant-specific components
│   │       ├── todo/      # TODO list frontend
│   │       ├── forum/     # Forum frontend
│   │       └── wiki/      # Wiki frontend
│   ├── package.json
│   └── vite.config.ts
└── 
└── backups/               # Automatic backups when switching variants
```

## 🛠️ Development Commands

```bash
# Start development server (auto-reload)
make dev

# Run linting and type checking
make check

# Format code
make format

# View application logs
make tail-log

# Build for production
make build

# Clean build artifacts
make clean
```

## 📚 Learning Path

### Step 1: Environment Setup
- Install Go and Node.js
- Understand the project structure
- Run the development server
- Explore the codebase

### Step 2: Database Fundamentals
- Learn SQLite basics
- Understand database migrations
- Practice SQL queries
- Learn about database relationships

### Step 3: Backend Development
- Understand Go basics
- Learn about HTTP handlers
- Implement CRUD operations
- Practice API design

### Step 4: Frontend Development
- Learn React and TypeScript
- Understand component architecture
- Practice state management
- Learn about routing

### Step 5: Advanced Features
- Implement authentication
- Add form validation
- Optimize performance
- Style with Tailwind CSS

## 🎨 Customization

### Switching Variants
Use the bootstrap script to switch between variants:
```bash
./bootstrap.sh todo    # Switch to TODO list
./bootstrap.sh forum   # Switch to forum
./bootstrap.sh wiki    # Switch to wiki
./bootstrap.sh list    # Show available variants
```

### Adding Custom Features
1. **Database**: Add migrations in `migrations/{variant}/`
2. **Backend**: Add handlers in `handlers/{variant}/`
3. **Frontend**: Add components in `web/src/variants/{variant}/`

### Configuration
Edit `bootstrap.config.json` to customize:
- Variant descriptions
- Learning paths
- Feature lists
- Difficulty levels

## 🔧 API Endpoints

### TODO List Application
```
GET    /api/projects           # List all projects
POST   /api/projects           # Create new project
GET    /api/projects/{slug}    # Get project details
PUT    /api/projects/{slug}    # Update project
DELETE /api/projects/{slug}    # Delete project

GET    /api/todos              # List todos (with filters)
POST   /api/todos              # Create new todo
GET    /api/todos/{id}         # Get todo details
PUT    /api/todos/{id}         # Update todo
DELETE /api/todos/{id}         # Delete todo
POST   /api/todos/{id}/complete   # Mark as complete
POST   /api/todos/{id}/uncomplete # Mark as incomplete

GET    /api/stats              # Get statistics
```

### Forum Application
```
GET    /api/boards             # List all boards
POST   /api/boards             # Create new board
GET    /api/boards/{slug}      # Get board details

GET    /api/topics             # List topics
POST   /api/topics             # Create new topic
GET    /api/topics/{id}        # Get topic details
PUT    /api/topics/{id}        # Update topic

GET    /api/posts              # List posts
POST   /api/posts              # Create new post
GET    /api/posts/{id}         # Get post details
```

### Wiki Application
```
GET    /api/namespaces         # List all namespaces
POST   /api/namespaces         # Create new namespace
GET    /api/namespaces/{slug}  # Get namespace details

GET    /api/pages              # List pages
POST   /api/pages              # Create new page
GET    /api/pages/{id}         # Get page details
PUT    /api/pages/{id}         # Update page
DELETE /api/pages/{id}         # Delete page

GET    /api/revisions          # List revisions
GET    /api/revisions/{id}     # Get revision details

GET    /api/search             # Search pages
```

## 🎓 Educational Benefits

### For Students
- **Real-world Technologies**: Learn with production-ready tools
- **Progressive Complexity**: Start simple, add features incrementally
- **Best Practices**: Follow modern development patterns
- **Full-stack Understanding**: See how frontend and backend connect

### For Instructors
- **Flexible Curriculum**: Choose variant based on class level
- **Comprehensive Examples**: Complete applications, not just snippets
- **Assessment Ready**: Clear learning objectives and milestones
- **Extensible**: Easy to add custom requirements

## 🔍 Troubleshooting

### Common Issues

**Port already in use:**
```bash
# Kill processes using the port
lsof -ti:8080 | xargs kill -9
lsof -ti:5173 | xargs kill -9
```

**Database locked:**
```bash
# Remove database and restart
rm minibb.db
make dev
```

**Node modules issues:**
```bash
# Clean and reinstall
cd web && rm -rf node_modules package-lock.json
npm install
```

**Go module issues:**
```bash
# Clean Go modules
go clean -modcache
go mod tidy
```

## 📈 Performance Tips

### Development
- Use `make dev` for hot reloading
- Monitor logs with `make tail-log`
- Use browser dev tools for frontend debugging

### Production
- Run `make build` for optimized builds
- Use appropriate database indexes
- Implement caching strategies
- Monitor application performance

## 🤝 Contributing

1. **Fork the repository**
2. **Create a feature branch**
3. **Make your changes**
4. **Add tests if applicable**
5. **Submit a pull request**

### Development Guidelines
- Follow Go and TypeScript best practices
- Use meaningful commit messages
- Document complex logic
- Test your changes thoroughly

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- **MiniBB**: Original bulletin board inspiration
- **Go Community**: For excellent web development tools
- **React Team**: For the fantastic frontend framework
- **TanStack**: For Query and Router libraries
- **Tailwind CSS**: For the utility-first CSS framework

---

**Happy coding! 🚀**

Start with the TODO list variant if you're new to web development, or jump to the wiki system if you're ready for a challenge. Each variant is designed to teach different aspects of full-stack development while maintaining the same core architecture.

Remember: The best way to learn is by doing. Pick a variant, start coding, and don't be afraid to experiment!