# ✅ Setup Checklist

Quick reference checklist for setting up and using the WebApp Bootstrap Project.

## 📋 Pre-Setup Checklist

### Prerequisites Installation
- [ ] **Go 1.21+** installed and in PATH
  ```bash
  go version  # Should show go1.21.x or higher
  ```

- [ ] **Node.js 18+** installed
  ```bash
  node --version  # Should show v18.x.x or higher
  ```

- [ ] **npm 9+** installed
  ```bash
  npm --version  # Should show 9.x.x or higher
  ```

- [ ] **Git** installed
  ```bash
  git --version  # Should show git version 2.x.x
  ```

- [ ] **Terminal/Command Line** access
- [ ] **Text Editor** (VS Code recommended)
- [ ] **Web Browser** (Chrome/Firefox/Safari)

### Optional Tools
- [ ] **SQLite Browser** for database inspection
- [ ] **jq** for JSON processing
- [ ] **curl** for API testing

## 🚀 Initial Setup Checklist

### 1. Project Download
- [ ] Clone or download project files
- [ ] Navigate to project directory
- [ ] Verify all files are present
  ```bash
  ls -la  # Should see bootstrap.sh, Makefile, etc.
  ```

### 2. Scripts Preparation
- [ ] Make scripts executable
  ```bash
  chmod +x bootstrap.sh test-bootstrap.sh
  ```

### 3. System Validation
- [ ] Run test suite
  ```bash
  ./test-bootstrap.sh
  ```
- [ ] All tests pass ✅
- [ ] No error messages

### 4. Choose Starting Variant
- [ ] **Beginner?** Choose TODO: `./bootstrap.sh todo`
- [ ] **Intermediate?** Choose Forum: `./bootstrap.sh forum`  
- [ ] **Advanced?** Choose Wiki: `./bootstrap.sh wiki`

### 5. First Run
- [ ] Start development server
  ```bash
  make dev
  ```
- [ ] Backend starts on port 8080
- [ ] Frontend starts on port 5173
- [ ] Both servers running without errors

### 6. Browser Verification
- [ ] Open http://localhost:5173
- [ ] Application loads successfully
- [ ] No console errors (F12)
- [ ] Basic functionality works

## 🎯 Variant-Specific Setup

### TODO List Setup
- [ ] Switch to TODO variant: `./bootstrap.sh todo`
- [ ] Application starts successfully
- [ ] Dashboard shows statistics
- [ ] Can create new todo
- [ ] Can view projects
- [ ] Can filter todos
- [ ] Can mark todos complete

**Test Checklist:**
- [ ] Create 3 different todos
- [ ] Assign to different projects
- [ ] Set different priorities
- [ ] Add due dates
- [ ] Mark one as complete
- [ ] View dashboard statistics

### Forum Setup
- [ ] Switch to Forum variant: `./bootstrap.sh forum`
- [ ] Application starts successfully
- [ ] Can view boards
- [ ] Can create new topic
- [ ] Can reply to topics
- [ ] Tripcode system works
- [ ] Read status tracking works

**Test Checklist:**
- [ ] Create topic in General board
- [ ] Reply to your own topic
- [ ] Create topic with tripcode
- [ ] Navigate between boards
- [ ] Check read status indicators

### Wiki Setup
- [ ] Switch to Wiki variant: `./bootstrap.sh wiki`
- [ ] Application starts successfully
- [ ] Can view welcome page
- [ ] Can edit existing pages
- [ ] Can create new pages
- [ ] Wiki links work
- [ ] Search functionality works

**Test Checklist:**
- [ ] Edit the welcome page
- [ ] Create a new page with wiki links
- [ ] View revision history
- [ ] Search for content
- [ ] Navigate between namespaces

## 🔧 Development Environment Setup

### IDE Configuration
- [ ] **VS Code** (recommended)
  - [ ] Go extension installed
  - [ ] TypeScript extension installed
  - [ ] React snippets extension
  - [ ] Tailwind CSS extension
  - [ ] SQLite viewer extension

### Git Configuration
- [ ] Initialize git repository (if not already)
  ```bash
  git init
  git add .
  git commit -m "Initial commit"
  ```

### Environment Variables
- [ ] Check .env file exists after variant selection
- [ ] Verify VARIANT setting matches chosen variant
- [ ] APP_NAME reflects current variant

## 📱 Application Testing Checklist

### Basic Functionality
- [ ] Application loads without errors
- [ ] All navigation links work
- [ ] Forms submit successfully
- [ ] Data persists between sessions
- [ ] Responsive design works on mobile

### API Testing
- [ ] Backend API responds to requests
- [ ] CRUD operations work
- [ ] Error handling works
- [ ] Rate limiting functions (if applicable)

### Database Testing
- [ ] Database file created (minibb.db)
- [ ] Tables exist with correct schema
- [ ] Sample data loaded
- [ ] Queries execute successfully

### Frontend Testing
- [ ] React components render
- [ ] State management works
- [ ] Form validation works
- [ ] UI interactions work
- [ ] Styling applied correctly

## 🔄 Switching Variants Checklist

### Before Switching
- [ ] Current work is saved
- [ ] No pending changes
- [ ] Database backed up (automatic)

### During Switch
- [ ] Run bootstrap command: `./bootstrap.sh [variant]`
- [ ] Wait for completion message
- [ ] No error messages during switch

### After Switch
- [ ] New variant loads successfully
- [ ] Database schema matches variant
- [ ] Frontend matches variant
- [ ] All features work as expected

## 🚨 Troubleshooting Checklist

### Common Issues
- [ ] **Port conflicts?** Kill processes: `lsof -ti:8080 | xargs kill -9`
- [ ] **Database locked?** Remove file: `rm minibb.db`
- [ ] **Dependencies missing?** Reinstall: `go mod tidy && cd web && npm install`
- [ ] **Permission errors?** Fix scripts: `chmod +x bootstrap.sh test-bootstrap.sh`

### Diagnostic Commands
- [ ] Run health check: `./test-bootstrap.sh`
- [ ] Check logs: `make tail-log`
- [ ] Verify processes: `ps aux | grep minibb`
- [ ] Check ports: `lsof -i:8080 && lsof -i:5173`

## 📚 Learning Readiness Checklist

### Knowledge Prerequisites
- [ ] **Basic programming concepts** (variables, functions, loops)
- [ ] **Command line familiarity** (cd, ls, mkdir)
- [ ] **Text editor proficiency** (create, edit, save files)
- [ ] **Web browser basics** (inspect element, console)

### Learning Materials Ready
- [ ] **Documentation** files read
- [ ] **Code editor** configured
- [ ] **Note-taking** system prepared
- [ ] **Practice environment** set up

### Learning Goals Set
- [ ] **Skill level** assessed
- [ ] **Learning path** chosen
- [ ] **Timeline** established
- [ ] **Success metrics** defined

## 🎓 Next Steps Checklist

### Immediate Actions
- [ ] **Explore** chosen variant thoroughly
- [ ] **Complete** basic exercises
- [ ] **Understand** code structure
- [ ] **Make** first small changes

### Short-term Goals (Week 1)
- [ ] **Master** basic functionality
- [ ] **Complete** learning exercises
- [ ] **Understand** architecture
- [ ] **Build** simple features

### Medium-term Goals (Week 2-3)
- [ ] **Add** new features
- [ ] **Customize** application
- [ ] **Optimize** performance
- [ ] **Implement** advanced features

### Long-term Goals (Week 4+)
- [ ] **Deploy** application
- [ ] **Add** authentication
- [ ] **Scale** application
- [ ] **Build** from scratch

## 📊 Success Metrics

### Technical Metrics
- [ ] **Application runs** without errors
- [ ] **All tests pass** consistently
- [ ] **Features work** as expected
- [ ] **Performance** is acceptable

### Learning Metrics
- [ ] **Understand** codebase structure
- [ ] **Can modify** existing features
- [ ] **Can add** new features
- [ ] **Can debug** issues

### Completion Metrics
- [ ] **All exercises** completed
- [ ] **All features** explored
- [ ] **All documentation** read
- [ ] **Ready for** next variant

## 🔗 Quick Reference

### Essential Commands
```bash
# Start development
make dev

# View logs
make tail-log

# Run tests
./test-bootstrap.sh

# Switch variants
./bootstrap.sh [todo|forum|wiki]

# Reset database
rm minibb.db && make dev
```

### Important URLs
- **Application:** http://localhost:5173
- **API:** http://localhost:8080/api/
- **Database:** sqlite3 minibb.db

### Help Resources
- **Installation:** [INSTALLATION_GUIDE.md](INSTALLATION_GUIDE.md)
- **Usage:** [USAGE_GUIDE.md](USAGE_GUIDE.md)
- **Learning:** [LEARNING_PATHS.md](LEARNING_PATHS.md)
- **Troubleshooting:** [TROUBLESHOOTING.md](TROUBLESHOOTING.md)

---

**🎉 Congratulations!** If you've checked all the boxes, you're ready to start coding!

Remember: Take your time with setup. A properly configured environment will save you hours of debugging later.

Happy coding! 🚀