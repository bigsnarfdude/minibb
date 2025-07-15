# 🚀 Installation & Setup Guide

Complete step-by-step guide to install, set up, and start using the WebApp Bootstrap Project.

## 📋 Prerequisites

### Required Software

#### 1. Go Programming Language
**Version:** 1.21 or higher

**Installation:**
```bash
# macOS (using Homebrew)
brew install go

# Linux (Ubuntu/Debian)
sudo apt update
sudo apt install golang-go

# Windows (using Chocolatey)
choco install golang

# Or download from: https://golang.org/dl/
```

**Verify installation:**
```bash
go version
# Should output: go version go1.21.x
```

#### 2. Node.js & npm
**Version:** Node.js 18+ and npm 9+

**Installation:**
```bash
# macOS (using Homebrew)
brew install node

# Linux (using NodeSource)
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# Windows (using Chocolatey)
choco install nodejs

# Or download from: https://nodejs.org/
```

**Verify installation:**
```bash
node --version  # Should output: v18.x.x or higher
npm --version   # Should output: 9.x.x or higher
```

#### 3. Git
**For cloning the repository**

**Installation:**
```bash
# macOS (using Homebrew)
brew install git

# Linux (Ubuntu/Debian)
sudo apt update
sudo apt install git

# Windows (using Chocolatey)
choco install git

# Or download from: https://git-scm.com/
```

**Verify installation:**
```bash
git --version
# Should output: git version 2.x.x
```

### Optional Tools

#### 1. SQLite Browser (Recommended)
**For database inspection and debugging**

**Installation:**
```bash
# macOS (using Homebrew)
brew install --cask db-browser-for-sqlite

# Linux (Ubuntu/Debian)
sudo apt install sqlitebrowser

# Windows: Download from https://sqlitebrowser.org/
```

#### 2. VS Code (Recommended)
**With Go and TypeScript extensions**

**Installation:**
```bash
# macOS (using Homebrew)
brew install --cask visual-studio-code

# Or download from: https://code.visualstudio.com/
```

**Recommended Extensions:**
- Go (by Google)
- TypeScript and JavaScript Language Features
- ES7+ React/Redux/React-Native snippets
- Tailwind CSS IntelliSense
- SQLite Viewer

#### 3. jq (Optional)
**For JSON parsing and validation**

**Installation:**
```bash
# macOS (using Homebrew)
brew install jq

# Linux (Ubuntu/Debian)
sudo apt install jq

# Windows (using Chocolatey)
choco install jq
```

## 🛠️ Installation Steps

### Step 1: Clone the Repository

```bash
# Clone the repository
git clone https://github.com/your-username/minibb-bootstrap.git
cd minibb-bootstrap

# Or if you have the files locally
cd /path/to/minibb-main
```

### Step 2: Verify Prerequisites

```bash
# Check if all required tools are installed
go version
node --version
npm --version
git --version

# All commands should return version numbers
```

### Step 3: Run the Test Suite

```bash
# Make the test script executable
chmod +x test-bootstrap.sh

# Run comprehensive tests
./test-bootstrap.sh

# You should see "All tests passed! ✅"
```

### Step 4: Choose Your Variant

```bash
# View available variants
./bootstrap.sh list

# Choose your starting variant (recommended: todo for beginners)
./bootstrap.sh todo
```

**Output should look like:**
```
[INFO] Setting up todo variant...
[INFO] Database backed up to backups/20240115_143022/
[INFO] Applied todo database migrations
[INFO] Applied todo API handlers
[INFO] Applied todo frontend files
[INFO] Updated package.json for todo
[INFO] Created .env file for todo
[INFO] Installing Go dependencies...
[INFO] Installing Node.js dependencies...
[INFO] ✅ todo variant setup complete!

Next steps:
  1. Run 'make dev' to start development server
  2. Open http://localhost:5173 in your browser
  3. Start coding!
```

### Step 5: Start Development

```bash
# Start the development server
make dev

# This will start both backend (Go) and frontend (React)
# Backend runs on :8080
# Frontend runs on :5173 (with proxy to backend)
```

### Step 6: Verify Everything Works

1. **Open your browser** to `http://localhost:5173`
2. **You should see** the application running
3. **Test basic functionality** (create a todo, view dashboard, etc.)

## 🎯 Choosing Your Variant

### Decision Matrix

| Variant | Difficulty | Best For | Time Investment | Skills Learned |
|---------|------------|----------|-----------------|----------------|
| **TODO** | Beginner | First web app | 2-3 days | CRUD, Forms, State |
| **Forum** | Intermediate | Understanding discussions | 1-2 weeks | Auth, Content, Pagination |
| **Wiki** | Advanced | Complex systems | 2-3 weeks | Versioning, Search, Relations |

### Recommended Learning Paths

#### 🟢 Complete Beginner
**Start with TODO variant**
```bash
./bootstrap.sh todo
```

**What you'll learn:**
- Basic database design
- REST API creation
- React components
- Form handling
- State management

**Time:** 2-3 days
**Next:** Move to Forum variant

#### 🟡 Some Programming Experience
**Start with Forum variant**
```bash
./bootstrap.sh forum
```

**What you'll learn:**
- User authentication concepts
- Content management
- Real-time updates
- Advanced UI patterns

**Time:** 1-2 weeks
**Next:** Try Wiki variant or customize Forum

#### 🔴 Experienced Developer
**Start with Wiki variant**
```bash
./bootstrap.sh wiki
```

**What you'll learn:**
- Version control systems
- Full-text search
- Complex relationships
- Performance optimization

**Time:** 2-3 weeks
**Next:** Build your own variant

## 📚 Detailed Usage Instructions

### TODO Variant Usage

#### Starting the TODO App
```bash
# Switch to TODO variant
./bootstrap.sh todo

# Start development
make dev

# Open browser to http://localhost:5173
```

#### Key Features to Explore
1. **Dashboard** - View statistics and project overview
2. **Create Todo** - Add new tasks with priorities and due dates
3. **Projects** - Organize todos into categories
4. **Filtering** - Sort by status, priority, or search
5. **Comments** - Add notes to todos

#### Learning Exercises
1. **Create a personal project** with 5 todos
2. **Set different priorities** and due dates
3. **Add comments** to todos
4. **Complete some todos** and see statistics update
5. **Try the search functionality**

### Forum Variant Usage

#### Starting the Forum App
```bash
# Switch to Forum variant
./bootstrap.sh forum

# Start development
make dev

# Open browser to http://localhost:5173
```

#### Key Features to Explore
1. **Boards** - Different discussion categories
2. **Topics** - Create discussion threads
3. **Posts** - Reply to discussions
4. **Tripcode** - Anonymous identity verification
5. **Read Status** - Track what you've read

#### Learning Exercises
1. **Create a new topic** in different boards
2. **Reply to existing topics**
3. **Experiment with tripcode** authentication
4. **Navigate between boards**
5. **Check read status** functionality

### Wiki Variant Usage

#### Starting the Wiki App
```bash
# Switch to Wiki variant
./bootstrap.sh wiki

# Start development
make dev

# Open browser to http://localhost:5173
```

#### Key Features to Explore
1. **Pages** - Create and edit wiki pages
2. **Namespaces** - Organize content into categories
3. **Revisions** - View page history
4. **Links** - Create internal page links
5. **Search** - Full-text search across all pages

#### Learning Exercises
1. **Edit the welcome page**
2. **Create a new page** with wiki links
3. **View revision history**
4. **Create pages in different namespaces**
5. **Use the search functionality**

## 🔧 Development Commands

### Essential Commands

```bash
# Start development server (most important!)
make dev

# View application logs
make tail-log

# Check code quality
make check

# Format code
make format

# Build for production
make build

# Clean build artifacts
make clean
```

### Switching Variants

```bash
# List available variants
./bootstrap.sh list

# Switch to TODO variant
./bootstrap.sh todo

# Switch to Forum variant
./bootstrap.sh forum

# Switch to Wiki variant
./bootstrap.sh wiki

# Get help
./bootstrap.sh help
```

### Database Management

```bash
# View database contents (if SQLite browser installed)
open minibb.db

# Or use sqlite3 command line
sqlite3 minibb.db ".tables"
sqlite3 minibb.db ".schema"
```

## 🚨 Troubleshooting

### Common Issues & Solutions

#### 1. Port Already in Use
**Problem:** `address already in use :8080` or `:5173`

**Solution:**
```bash
# Kill processes on these ports
lsof -ti:8080 | xargs kill -9
lsof -ti:5173 | xargs kill -9

# Then restart
make dev
```

#### 2. Database Locked
**Problem:** `database is locked`

**Solution:**
```bash
# Stop the server (Ctrl+C)
# Remove database file
rm minibb.db

# Restart (this will recreate the database)
make dev
```

#### 3. Go Module Issues
**Problem:** `module not found` or Go errors

**Solution:**
```bash
# Clean and rebuild modules
go clean -modcache
go mod tidy
go mod download

# Try again
make dev
```

#### 4. Node.js Issues
**Problem:** `Cannot find module` or npm errors

**Solution:**
```bash
# Navigate to web directory
cd web

# Clean install
rm -rf node_modules package-lock.json
npm install

# Go back to root
cd ..

# Try again
make dev
```

#### 5. Bootstrap Script Issues
**Problem:** `Permission denied` or script errors

**Solution:**
```bash
# Make scripts executable
chmod +x bootstrap.sh
chmod +x test-bootstrap.sh

# Try again
./bootstrap.sh todo
```

### Getting Help

1. **Run the test suite:** `./test-bootstrap.sh`
2. **Check logs:** `make tail-log`
3. **Verify prerequisites:** All tools installed?
4. **Try a clean setup:** Remove database and restart
5. **Check the troubleshooting section** in BOOTSTRAP_README.md

## 🎓 Next Steps After Installation

### 1. Explore the Codebase
```bash
# Backend code
ls -la internal/
ls -la handlers/

# Frontend code
ls -la web/src/
ls -la web/src/variants/

# Database structure
sqlite3 minibb.db ".schema"
```

### 2. Make Your First Change
**TODO Variant:** Add a new priority level
**Forum Variant:** Add a new board
**Wiki Variant:** Create a new namespace

### 3. Learn the Architecture
- **Backend:** Go with Chi router
- **Frontend:** React with TypeScript
- **Database:** SQLite with migrations
- **Build:** Vite for frontend, Go for backend

### 4. Customize Your Project
- Add new features
- Modify the UI
- Create new database tables
- Add new API endpoints

## 📝 Quick Reference

### File Structure
```
minibb-main/
├── bootstrap.sh              # Project switcher
├── Makefile                  # Development commands
├── go.mod                    # Go dependencies
├── web/package.json          # Node.js dependencies
├── migrations/               # Database schemas
├── handlers/                 # API handlers
├── internal/                 # Core Go code
└── web/src/                  # Frontend code
```

### Key URLs
- **Main App:** http://localhost:5173
- **Backend API:** http://localhost:8080/api/
- **Database:** SQLite file: `minibb.db`

### Essential Commands
```bash
make dev          # Start development
make tail-log     # View logs
./bootstrap.sh    # Switch variants
```

---

**🎉 You're ready to start coding!**

Choose your variant, run `make dev`, and begin your web development journey. Remember: start with TODO if you're new to web development, or jump to Forum/Wiki if you have experience.

Happy coding! 🚀