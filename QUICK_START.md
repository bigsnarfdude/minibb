# ⚡ Quick Start Guide

Get up and running with the WebApp Bootstrap Project in 5 minutes!

## 🏃‍♂️ Super Quick Setup

### 1. One-Command Setup
```bash
# Copy and paste this entire block
git clone https://github.com/your-username/minibb-bootstrap.git
cd minibb-bootstrap
chmod +x bootstrap.sh test-bootstrap.sh
./test-bootstrap.sh
```

If tests pass, continue to step 2. If not, see [Full Installation Guide](INSTALLATION_GUIDE.md).

### 2. Choose Your Path
```bash
# For beginners (recommended)
./bootstrap.sh todo

# For intermediate developers
./bootstrap.sh forum

# For advanced developers
./bootstrap.sh wiki
```

### 3. Start Coding
```bash
make dev
```

Open http://localhost:5173 - you're ready to code! 🎉

## 🎯 Choose Your Adventure

### 🟢 New to Web Development?
**Start with TODO List**
```bash
./bootstrap.sh todo
make dev
```

**What you'll build:**
- Task management app
- Create, edit, complete todos
- Project organization
- Due dates and priorities

**Time to working app:** 5 minutes
**Learning time:** 2-3 days

### 🟡 Some Programming Experience?
**Start with Forum**
```bash
./bootstrap.sh forum
make dev
```

**What you'll build:**
- Discussion boards
- Topic creation and replies
- User authentication
- Read status tracking

**Time to working app:** 5 minutes
**Learning time:** 1-2 weeks

### 🔴 Experienced Developer?
**Start with Wiki**
```bash
./bootstrap.sh wiki
make dev
```

**What you'll build:**
- Collaborative wiki
- Page editing and versioning
- Internal linking
- Full-text search

**Time to working app:** 5 minutes
**Learning time:** 2-3 weeks

## 📱 What Each App Looks Like

### TODO App Features
- ✅ Create and complete todos
- 📊 Statistics dashboard
- 📁 Project organization
- 🔍 Search and filtering
- 📝 Comments on todos
- 📅 Due dates and priorities

### Forum App Features
- 💬 Discussion boards
- 🧵 Topic threads
- 👤 Tripcode authentication
- 📖 Read status tracking
- ✏️ Markdown formatting
- 🔢 Pagination

### Wiki App Features
- 📄 Page creation and editing
- 🔄 Revision history
- 🔗 Internal page linking
- 🔍 Full-text search
- 📚 Namespaces
- 👥 Collaborative editing

## 🛠️ Essential Commands

### Daily Development
```bash
make dev          # Start development server
make tail-log     # View application logs
make check        # Run linting and type checking
make format       # Format all code
```

### Switching Between Apps
```bash
./bootstrap.sh list    # Show available apps
./bootstrap.sh todo    # Switch to TODO app
./bootstrap.sh forum   # Switch to Forum app
./bootstrap.sh wiki    # Switch to Wiki app
```

### Database Management
```bash
sqlite3 minibb.db ".tables"     # View database tables
sqlite3 minibb.db ".schema"     # View database structure
```

## 🚨 Quick Troubleshooting

### App Won't Start?
```bash
# Kill existing processes
lsof -ti:8080 | xargs kill -9
lsof -ti:5173 | xargs kill -9

# Clean and restart
rm minibb.db
make dev
```

### Missing Dependencies?
```bash
# Check prerequisites
go version    # Should be 1.21+
node --version # Should be 18+
npm --version  # Should be 9+

# Reinstall if needed
go mod tidy
cd web && npm install && cd ..
```

### Database Issues?
```bash
# Reset database
rm minibb.db
make dev  # Will recreate database
```

## 📚 Learning Resources

### Code Structure
```
minibb-main/
├── internal/           # Go backend code
├── web/src/           # React frontend code
├── migrations/        # Database schemas
├── handlers/          # API endpoints
└── Makefile          # Development commands
```

### Key Files to Explore
- `internal/handlers/api.go` - API endpoints
- `web/src/routes/` - Frontend routes
- `web/src/components/` - React components
- `migrations/*/` - Database structure

### Next Steps
1. **Make your first change** - Edit a component
2. **Add a new feature** - Create a new todo field
3. **Explore the database** - View tables and data
4. **Read the code** - Understand the architecture

## 🎓 Learning Path Recommendations

### Week 1: Get Comfortable
- Set up your development environment
- Explore the chosen variant
- Make small UI changes
- Understand the file structure

### Week 2: Add Features
- Add new database fields
- Create new API endpoints
- Build new frontend components
- Connect frontend to backend

### Week 3: Advanced Features
- Implement user authentication
- Add form validation
- Optimize performance
- Deploy your application

### Week 4: Make It Your Own
- Design your own features
- Customize the UI completely
- Add new variants
- Share your project

## 🔗 Quick Links

- **[Full Installation Guide](INSTALLATION_GUIDE.md)** - Detailed setup instructions
- **[Bootstrap README](BOOTSTRAP_README.md)** - Complete project documentation
- **[Project Structure](CLAUDE.md)** - Technical details and architecture

## 🎉 You're Ready!

Pick your variant, run `make dev`, and start building!

```bash
# Choose your adventure
./bootstrap.sh [todo|forum|wiki]
make dev
```

Open http://localhost:5173 and start coding! 🚀

---

**Need help?** Check the [troubleshooting section](INSTALLATION_GUIDE.md#troubleshooting) or run `./test-bootstrap.sh` to validate your setup.