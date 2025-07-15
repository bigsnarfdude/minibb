# 🗣️ Claude Code web starter project - Example Simple Bulletin Board System

A lightweight, modern CC project. Built with Go backend and React frontend, featuring tripcode authentication and real-time discussions.

## ✨ Features

- **🔐 Tripcode Authentication** - Anonymous posting with optional identity verification
- **📋 Multiple Boards** - Organize discussions by topic
- **📝 Markdown Support** - Rich text formatting in posts
- **👁️ Read Status Tracking** - Keep track of what you've read
- **⚡ Real-time Updates** - Live discussions without page refreshes
- **📱 Responsive Design** - Works on desktop and mobile

## 🚀 Quick Start

### Prerequisites

Before you begin, make sure you have:
- **Go 1.21+** installed ([Download Go](https://golang.org/dl/))
- **Node.js 18+** installed ([Download Node.js](https://nodejs.org/))
- **Git** installed ([Download Git](https://git-scm.com/))

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/minibb.git
   cd minibb
   ```

2. **Start the development server**
   ```bash
   make dev
   ```
   
   This command will:
   - Install all dependencies automatically
   - Start the Go backend server on port 8080
   - Start the React frontend development server on port 5173
   - Set up hot-reloading for both backend and frontend

3. **Open your browser**
   
   Navigate to [http://localhost:5173](http://localhost:5173) and start using MiniBB!

That's it! 🎉 You now have a fully functional bulletin board system running locally.

## 📱 Using MiniBB

### Creating Your First Post

1. Visit the main page and click on a board (like "General Discussion")
2. Click "New Topic" to start a discussion
3. Fill in:
   - **Subject** - What your topic is about
   - **Name** (optional) - How you want to appear
   - **Tripcode** (optional) - Add `#password` to your name for identity verification
   - **Message** - Your post content (supports Markdown)
4. Click "Post" to publish

### Understanding Tripcodes

Tripcodes let you prove your identity without registration:
- Type your name as `Username#password`
- Example: `Alice#mysecret123`
- This creates a unique identifier that others can't fake
- Same password always generates the same tripcode

### Markdown Formatting

You can format your posts with Markdown:
- **Bold text**: `**bold**`
- *Italic text*: `*italic*`
- `Code`: `` `code` ``
- Links: `[text](url)`
- Lists, headers, and more!

## 🛠️ Development Commands

```bash
# Start development server (most common)
make dev

# View application logs
make tail-log

# Check code quality (linting & type checking)
make check

# Format all code
make format

# Build for production
make build

# Clean build artifacts
make clean
```

**💡 Tip**: Keep `make dev` running while you develop. It automatically reloads when you make changes!

## 📁 Project Structure

```
minibb/
├── cmd/minibb/           # Go application entry point
│   └── main.go
├── internal/             # Backend Go code
│   ├── db/              # Database layer
│   ├── handlers/        # API request handlers
│   ├── models/          # Data models
│   ├── server/          # HTTP server setup
│   └── utils/           # Utility functions
├── web/                 # Frontend React application
│   ├── src/
│   │   ├── routes/      # Page components
│   │   ├── lib/         # API client & utilities
│   │   └── main.tsx     # App entry point
│   ├── package.json
│   └── vite.config.ts
├── migrations/          # Database schema
├── Makefile            # Development commands
└── README.md           # This file
```

## 🗄️ Database

MiniBB uses SQLite for simplicity. The database file (`minibb.db`) is created automatically when you first run the application.

To explore the database:
```bash
sqlite3 minibb.db
.schema  # View table structure
.quit    # Exit
```

## 🔧 Configuration

### Backend (Go)
- **Port**: Set `PORT` environment variable (default: 8080)
- **Database**: SQLite file at `minibb.db`
- **Logs**: Written to `dev.log` during development

### Frontend (React)
- **Development server**: Port 5173
- **API proxy**: Automatically proxies `/api/*` to backend
- **Build output**: Embedded in Go binary for production

## 🌐 API Endpoints

The backend provides a REST API:

- `GET /api/boards` - List all boards
- `GET /api/boards/{board}/topics` - List topics in a board
- `GET /api/boards/{board}/topics/{topic}` - Get topic with posts
- `POST /api/boards/{board}/topics` - Create new topic
- `POST /api/boards/{board}/topics/{topic}/posts` - Reply to topic

## 🎯 Next Steps

### For New Developers
1. **Explore the code** - Start with `web/src/routes/index.tsx` for the frontend
2. **Make small changes** - Try modifying the welcome message
3. **Add a feature** - Maybe a post counter or new board?
4. **Learn the patterns** - See how data flows from database to UI

### For Experienced Developers
1. **Add features** - User avatars, post reactions, search functionality
2. **Improve performance** - Implement caching, optimize queries
3. **Add testing** - Unit tests, integration tests, e2e tests
4. **Deploy to production** - Set up CI/CD, containerization

## 🐛 Troubleshooting

### Common Issues

**Port already in use**
```bash
# Kill processes on ports 8080 or 5173
lsof -ti:8080 | xargs kill -9
lsof -ti:5173 | xargs kill -9
```

**Go dependencies issues**
```bash
go mod tidy
go mod download
```

**Frontend build issues**
```bash
cd web
npm install
npm run build
```

**Database issues**
```bash
# Delete and recreate database
rm minibb.db
make dev  # Will recreate automatically
```

### Getting Help

1. Check the logs: `make tail-log`
2. Ensure all prerequisites are installed
3. Try restarting: stop `make dev` and run it again
4. Check that ports 8080 and 5173 are available

## 🤝 Contributing

Contributions are welcome! Here's how:

1. Fork the repository
2. Create a feature branch: `git checkout -b my-feature`
3. Make your changes
4. Test your changes: `make check`
5. Commit your changes: `git commit -m "Add my feature"`
6. Push to your fork: `git push origin my-feature`
7. Open a Pull Request

### Development Guidelines
- Write clear commit messages
- Test your changes before submitting
- Follow the existing code style
- Update documentation for new features

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🌟 Acknowledgments

- Inspired by phpBB and 4chan
- Built with modern web technologies
- Designed for simplicity and ease of use

---

**Ready to start discussions?** Run `make dev` and visit [http://localhost:5173](http://localhost:5173)! 🚀
