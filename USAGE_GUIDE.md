# 📖 Usage Guide

Detailed instructions for using each variant of the WebApp Bootstrap Project.

## 🎯 Overview

This guide covers how to use each variant effectively, including step-by-step tutorials, feature explanations, and learning exercises.

## 📋 TODO List Application

### Getting Started with TODO

#### 1. Switch to TODO Variant
```bash
./bootstrap.sh todo
make dev
```

#### 2. First Look at the Dashboard
Navigate to http://localhost:5173

**What you'll see:**
- Statistics cards showing total/completed/pending todos
- Completion rate progress bar
- Project overview with color-coded projects
- Quick action buttons
- Getting started tips

#### 3. Create Your First Todo

**Step-by-step:**
1. Click "New Todo" button (top-right or quick actions)
2. Fill out the form:
   - **Title:** "Learn React basics"
   - **Project:** Personal Tasks
   - **Priority:** High
   - **Due Date:** Tomorrow
   - **Description:** "Complete React tutorial and build first component"
   - **Author:** Your name

3. Click "Create Todo"

#### 4. Managing Todos

**View All Todos:**
- Click "All Todos" in navigation
- Use filters to sort by status, priority, or search
- Check off completed todos
- Edit existing todos
- Delete todos you no longer need

**Project Organization:**
- View projects in the dashboard
- Create new projects with custom colors
- Filter todos by project
- See todo counts per project

### TODO Features Deep Dive

#### Dashboard Analytics
```
📊 Statistics Overview
- Total todos across all projects
- Completion percentage
- High-priority items
- Overdue tasks
```

#### Todo Management
```
✅ Create todos with:
- Title and description
- Project assignment
- Priority levels (Low/Medium/High)
- Due dates
- Author tracking

📝 Edit todos:
- Update any field
- Change completion status
- Add/edit comments
- Modify due dates
```

#### Project Organization
```
📁 Project Features:
- Color-coded projects
- Custom project names
- Project descriptions
- Todo counts per project
- Project-specific statistics
```

### TODO Learning Exercises

#### Exercise 1: Personal Task Manager
**Goal:** Create a personal productivity system

**Steps:**
1. Create projects for different life areas:
   - Work (blue)
   - Personal (green)
   - Learning (purple)
   - Health (red)

2. Add 3-5 todos to each project
3. Set realistic due dates
4. Use appropriate priority levels
5. Track completion over a week

#### Exercise 2: Team Project Simulation
**Goal:** Simulate a team project workflow

**Steps:**
1. Create a "Team Project" project
2. Add todos for different team members
3. Use comments to simulate team communication
4. Track progress through completion
5. Analyze completion statistics

#### Exercise 3: Advanced Features
**Goal:** Explore all TODO features

**Steps:**
1. Use search functionality to find specific todos
2. Filter by different criteria
3. Add detailed comments to todos
4. Create todos with various due dates
5. Monitor overdue tasks

### TODO API Reference

#### Key Endpoints
```
GET    /api/projects        # List all projects
POST   /api/projects        # Create new project
GET    /api/todos           # List todos (with filters)
POST   /api/todos           # Create new todo
PUT    /api/todos/{id}      # Update todo
POST   /api/todos/{id}/complete # Mark complete
GET    /api/stats           # Get statistics
```

#### Example API Calls
```bash
# Create a new todo
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "title": "Learn Go",
    "description": "Complete Go tutorial",
    "priority": "high",
    "author": "Student"
  }'

# Get all todos
curl http://localhost:8080/api/todos

# Mark todo as complete
curl -X POST http://localhost:8080/api/todos/1/complete
```

## 🏛️ Forum Application

### Getting Started with Forum

#### 1. Switch to Forum Variant
```bash
./bootstrap.sh forum
make dev
```

#### 2. Explore the Forum Structure
Navigate to http://localhost:5173

**What you'll see:**
- List of discussion boards
- Recent topics and activity
- Board descriptions and post counts
- Navigation between boards

#### 3. Create Your First Topic

**Step-by-step:**
1. Click on a board (e.g., "General Discussion")
2. Click "New Topic" button
3. Fill out the form:
   - **Title:** "Welcome to the Forum!"
   - **Author:** Your name (or use tripcode)
   - **Content:** Write your first post with Markdown
4. Click "Create Topic"

#### 4. Participate in Discussions

**Replying to Topics:**
- Click on any topic to view the full thread
- Scroll to the bottom for the reply form
- Write your response using Markdown
- Submit your reply

**Using Tripcodes:**
- Add `#password` to your name for a tripcode
- Example: `John#mypassword` becomes `John !!AbCdEf`
- This creates a verifiable identity without registration

### Forum Features Deep Dive

#### Board System
```
📋 Multiple Boards:
- General Discussion
- Technical Help
- Project Showcase
- Random Topics
- Each with description and stats
```

#### Topic Management
```
🧵 Topic Features:
- Create discussion threads
- Markdown formatting support
- Reply to existing topics
- View post counts and dates
- Track read/unread status
```

#### Authentication
```
👤 Tripcode System:
- Anonymous posting
- Optional identity verification
- No user registration required
- 4chan-style tripcode algorithm
```

### Forum Learning Exercises

#### Exercise 1: Community Building
**Goal:** Understand forum dynamics

**Steps:**
1. Create topics in different boards
2. Write meaningful posts with Markdown
3. Reply to create conversation threads
4. Use tripcodes for identity consistency
5. Observe how discussions develop

#### Exercise 2: Content Creation
**Goal:** Practice content formatting

**Steps:**
1. Create a technical help topic
2. Use Markdown formatting:
   - Headers and lists
   - Code blocks
   - Links and emphasis
3. Include images and structured content
4. Create a showcase topic with project details

#### Exercise 3: Forum Moderation
**Goal:** Understand forum management

**Steps:**
1. Create multiple boards for different topics
2. Post in appropriate boards
3. Observe post ordering and pagination
4. Track read status across sessions
5. Understand rate limiting

### Forum API Reference

#### Key Endpoints
```
GET    /api/boards          # List all boards
POST   /api/boards          # Create new board
GET    /api/topics          # List topics
POST   /api/topics          # Create new topic
GET    /api/posts           # List posts
POST   /api/posts           # Create new post
```

## 📚 Wiki Application

### Getting Started with Wiki

#### 1. Switch to Wiki Variant
```bash
./bootstrap.sh wiki
make dev
```

#### 2. Explore the Wiki Structure
Navigate to http://localhost:5173

**What you'll see:**
- Welcome page with wiki navigation
- Namespace organization (Main, Help, Project, Templates)
- Internal page linking
- Search functionality
- Recent changes

#### 3. Edit Your First Page

**Step-by-step:**
1. Click "Edit" on the welcome page
2. Modify the content using Markdown
3. Add some wiki links: `[[New Page]]`
4. Save your changes
5. View the revision history

#### 4. Create New Pages

**Creating Pages:**
- Click on red wiki links to create new pages
- Use the page creation form
- Choose appropriate namespace
- Write content with internal links
- Save and view your new page

### Wiki Features Deep Dive

#### Page Management
```
📄 Page Features:
- Create and edit pages
- Markdown content support
- Internal wiki linking
- Page history and revisions
- Cross-namespace linking
```

#### Version Control
```
🔄 Revision System:
- Full edit history
- Compare revisions
- Revert to previous versions
- Edit summaries
- Author tracking
```

#### Organization
```
📚 Namespace System:
- Main - primary content
- Help - documentation
- Project - project-specific pages
- Templates - reusable content
- Custom namespaces
```

### Wiki Learning Exercises

#### Exercise 1: Personal Knowledge Base
**Goal:** Create a personal wiki

**Steps:**
1. Create pages for different topics you're learning
2. Link related pages together
3. Use different namespaces for organization
4. Add comprehensive content with formatting
5. Create a personal homepage with navigation

#### Exercise 2: Project Documentation
**Goal:** Document a project comprehensively

**Steps:**
1. Create pages for:
   - Project overview
   - Technical specifications
   - User guides
   - API documentation
   - Meeting notes
2. Link all pages together
3. Use templates for consistency
4. Track changes over time

#### Exercise 3: Collaborative Wiki
**Goal:** Simulate team collaboration

**Steps:**
1. Create pages with different authors
2. Edit existing pages to add content
3. Use edit summaries to explain changes
4. View revision history
5. Create discussion pages for feedback

### Wiki API Reference

#### Key Endpoints
```
GET    /api/namespaces      # List all namespaces
POST   /api/namespaces      # Create new namespace
GET    /api/pages           # List pages
POST   /api/pages           # Create new page
PUT    /api/pages/{id}      # Update page
GET    /api/revisions       # List revisions
GET    /api/search          # Search pages
```

## 🔄 Switching Between Variants

### Saving Your Work
```bash
# Your work is automatically backed up when switching
# Backups are stored in backups/ directory
ls backups/
```

### Switching Process
```bash
# Current work is backed up automatically
./bootstrap.sh [new-variant]

# Database is reset to new variant
# Frontend is updated to new variant
# API handlers are updated to new variant
```

### Restoring Previous Work
```bash
# View available backups
ls backups/

# Manually restore database if needed
cp backups/20240115_143022/minibb.db ./
```

## 🎓 Learning Progression

### Week 1: Basic Usage
- Set up and explore chosen variant
- Use basic features
- Understand user interface
- Complete simple tasks

### Week 2: Feature Exploration
- Use all available features
- Understand data relationships
- Explore API endpoints
- Practice with different scenarios

### Week 3: Customization
- Modify existing features
- Add new functionality
- Explore code structure
- Make UI improvements

### Week 4: Advanced Usage
- Build complex workflows
- Integrate with external tools
- Optimize performance
- Deploy your application

## 🔗 Additional Resources

### Documentation
- [Installation Guide](INSTALLATION_GUIDE.md) - Setup instructions
- [Bootstrap README](BOOTSTRAP_README.md) - Complete project overview
- [Quick Start](QUICK_START.md) - Fast setup guide

### Code Examples
- Check `web/src/variants/` for frontend code
- Check `handlers/` for backend API code
- Check `migrations/` for database schemas

### Community
- Create issues for bugs or feature requests
- Share your customizations
- Help others learn web development

---

**Ready to dive deeper?** Choose your variant and start exploring! Each application offers unique learning opportunities and real-world development experience.

Happy coding! 🚀