# 🔧 Troubleshooting Guide

Comprehensive troubleshooting guide for common issues with the WebApp Bootstrap Project.

## 🚨 Quick Fix Checklist

**Before diving into specific issues, try these quick fixes:**

```bash
# 1. Kill all processes and restart
lsof -ti:8080 | xargs kill -9
lsof -ti:5173 | xargs kill -9
make dev

# 2. Reset database
rm minibb.db
make dev

# 3. Clean dependencies
go mod tidy
cd web && rm -rf node_modules package-lock.json && npm install && cd ..

# 4. Run tests
./test-bootstrap.sh

# 5. Check prerequisites
go version && node --version && npm --version
```

## 🔍 Common Issues

### 1. Application Won't Start

#### Issue: "Port already in use"
```
Error: listen tcp :8080: bind: address already in use
Error: Port 5173 is already in use
```

**Solution:**
```bash
# Find and kill processes using the ports
lsof -ti:8080 | xargs kill -9
lsof -ti:5173 | xargs kill -9

# Alternative: Find process IDs and kill manually
lsof -i:8080  # Shows process using port 8080
lsof -i:5173  # Shows process using port 5173
kill -9 [PID]

# Restart the application
make dev
```

#### Issue: "Command not found: make"
```
bash: make: command not found
```

**Solution:**
```bash
# macOS
xcode-select --install

# Linux (Ubuntu/Debian)
sudo apt update && sudo apt install build-essential

# Windows
# Install Make for Windows or use WSL
```

#### Issue: "Go command not found"
```
bash: go: command not found
```

**Solution:**
```bash
# Install Go (see INSTALLATION_GUIDE.md)
# Then add to PATH
export PATH=$PATH:/usr/local/go/bin

# Make permanent by adding to ~/.bashrc or ~/.zshrc
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
```

### 2. Database Issues

#### Issue: "Database is locked"
```
Error: database is locked
```

**Solution:**
```bash
# Stop all processes
pkill -f minibb
pkill -f npm

# Remove database file
rm minibb.db

# Restart (will recreate database)
make dev
```

#### Issue: "No such table: todos"
```
Error: no such table: todos
```

**Solution:**
```bash
# Check current variant
cat .env

# Re-run bootstrap to ensure correct schema
./bootstrap.sh [current-variant]

# Or manually remove database
rm minibb.db
make dev
```

#### Issue: "SQL syntax error"
```
Error: SQL syntax error near "..."
```

**Solution:**
```bash
# Check database schema
sqlite3 minibb.db ".schema"

# Compare with expected schema
cat migrations/[variant]/0001_initial_schema.sql

# Reset database if needed
rm minibb.db
./bootstrap.sh [variant]
```

### 3. Frontend Issues

#### Issue: "Cannot resolve module"
```
Error: Cannot resolve module '@/components/...'
```

**Solution:**
```bash
# Navigate to web directory
cd web

# Clean and reinstall dependencies
rm -rf node_modules package-lock.json
npm install

# Check TypeScript configuration
cat tsconfig.json

# Go back to root
cd ..
make dev
```

#### Issue: "TypeScript errors"
```
Error: Property 'xyz' does not exist on type...
```

**Solution:**
```bash
# Check TypeScript configuration
cd web
npx tsc --noEmit

# Fix type errors in the code
# Or temporarily disable strict mode in tsconfig.json
```

#### Issue: "React component not rendering"
```
Blank page or component not showing
```

**Solution:**
```bash
# Check browser console for errors
# Open browser dev tools (F12)

# Check for JavaScript errors
# Look for network errors in Network tab

# Check component imports
# Verify file paths are correct
```

### 4. Backend Issues

#### Issue: "Go module errors"
```
Error: go: cannot find module providing package...
```

**Solution:**
```bash
# Clean module cache
go clean -modcache

# Re-download dependencies
go mod tidy
go mod download

# Verify go.mod file
cat go.mod
```

#### Issue: "Handler not found"
```
Error: 404 Not Found for API endpoint
```

**Solution:**
```bash
# Check API routes
grep -r "api/" internal/server/

# Verify handler registration
cat internal/server/routes.go

# Check if correct variant is active
./bootstrap.sh list
cat .env
```

#### Issue: "CORS errors"
```
Error: Access to fetch at 'http://localhost:8080/api/...' from origin 'http://localhost:5173' has been blocked by CORS policy
```

**Solution:**
```bash
# Check CORS configuration in server
grep -r "cors" internal/

# Verify development proxy
cat web/vite.config.ts

# Ensure both servers are running
make dev
```

### 5. Bootstrap Script Issues

#### Issue: "Permission denied"
```
bash: ./bootstrap.sh: Permission denied
```

**Solution:**
```bash
# Make script executable
chmod +x bootstrap.sh
chmod +x test-bootstrap.sh

# Run again
./bootstrap.sh todo
```

#### Issue: "Variant not found"
```
Error: No migrations found for variant: xyz
```

**Solution:**
```bash
# Check available variants
./bootstrap.sh list

# Verify directory structure
ls -la migrations/
ls -la handlers/
ls -la web/src/variants/

# Use correct variant name
./bootstrap.sh todo  # or forum, wiki
```

#### Issue: "Backup failed"
```
Error: Cannot create backup directory
```

**Solution:**
```bash
# Create backups directory
mkdir -p backups

# Check permissions
ls -la backups/

# Run bootstrap again
./bootstrap.sh [variant]
```

### 6. Build and Development Issues

#### Issue: "Build failed"
```
Error: Build failed with errors
```

**Solution:**
```bash
# Check for TypeScript errors
cd web
npm run build

# Fix any reported errors
# Check for missing dependencies
npm install

# Try development build
cd ..
make dev
```

#### Issue: "Hot reload not working"
```
Changes not reflected in browser
```

**Solution:**
```bash
# Check if development server is running
ps aux | grep vite

# Restart development server
make dev

# Clear browser cache
# Hard refresh: Ctrl+Shift+R (or Cmd+Shift+R on Mac)
```

#### Issue: "CSS not loading"
```
Page appears unstyled
```

**Solution:**
```bash
# Check Tailwind configuration
cat web/tailwind.config.js

# Verify CSS imports
grep -r "tailwind" web/src/

# Check for build errors
cd web
npm run build
```

### 7. Performance Issues

#### Issue: "Application running slowly"
```
Slow response times, high CPU usage
```

**Solution:**
```bash
# Check for infinite loops in code
# Monitor CPU usage: top or htop

# Check database queries
# Look for N+1 query problems

# Optimize database with indexes
sqlite3 minibb.db "EXPLAIN QUERY PLAN SELECT * FROM todos;"

# Enable production mode
make build
```

#### Issue: "High memory usage"
```
Application consuming too much memory
```

**Solution:**
```bash
# Monitor memory usage
ps aux | grep minibb

# Check for memory leaks
# Profile Go application
go tool pprof http://localhost:8080/debug/pprof/heap

# Optimize database connections
# Check for unclosed connections
```

## 🔬 Debugging Techniques

### 1. Backend Debugging

#### Using Logs
```bash
# View real-time logs
make tail-log

# Search for specific errors
grep -i "error" dev.log
grep -i "panic" dev.log
```

#### Using Go Debugger
```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Run with debugger
dlv debug ./cmd/minibb

# Set breakpoints and investigate
(dlv) break main.main
(dlv) continue
```

#### Database Debugging
```bash
# Connect to database
sqlite3 minibb.db

# Check table contents
.tables
SELECT * FROM todos LIMIT 5;

# Check schema
.schema todos

# Exit
.exit
```

### 2. Frontend Debugging

#### Browser Developer Tools
```bash
# Open dev tools (F12)
# Check Console tab for errors
# Check Network tab for API calls
# Check Sources tab for debugging
```

#### React Developer Tools
```bash
# Install React DevTools browser extension
# Inspect component state and props
# Profile performance
```

### 3. API Debugging

#### Using curl
```bash
# Test API endpoints
curl http://localhost:8080/api/todos
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Test","project_id":1,"author":"Test"}'
```

#### Using Browser Network Tab
```bash
# Open browser dev tools
# Go to Network tab
# Perform actions in the app
# Check requests and responses
```

## 📋 Diagnostic Commands

### System Health Check
```bash
# Check all prerequisites
echo "Go version: $(go version)"
echo "Node version: $(node --version)"
echo "npm version: $(npm --version)"
echo "Git version: $(git --version)"

# Check ports
lsof -i:8080
lsof -i:5173

# Check processes
ps aux | grep minibb
ps aux | grep node
```

### Application Health Check
```bash
# Test bootstrap functionality
./test-bootstrap.sh

# Check database
sqlite3 minibb.db "SELECT COUNT(*) FROM sqlite_master WHERE type='table';"

# Check API endpoints
curl -s http://localhost:8080/api/health || echo "API not responding"

# Check frontend
curl -s http://localhost:5173 || echo "Frontend not responding"
```

### Performance Diagnostics
```bash
# Check system resources
top -p $(pgrep minibb)
df -h  # Disk usage
free -h  # Memory usage

# Check database performance
sqlite3 minibb.db "PRAGMA cache_size;"
sqlite3 minibb.db "PRAGMA journal_mode;"
```

## 🛠️ Advanced Troubleshooting

### 1. Environment Issues

#### Check Environment Variables
```bash
# View current environment
env | grep -i go
env | grep -i node
env | grep -i npm

# Check project environment
cat .env
```

#### Path Issues
```bash
# Check PATH
echo $PATH

# Add missing paths
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/usr/local/bin
```

### 2. Network Issues

#### Check Firewall
```bash
# macOS
sudo pfctl -s all

# Linux
sudo iptables -L
sudo ufw status

# Windows
netsh advfirewall show allprofiles
```

#### Check DNS
```bash
# Test DNS resolution
nslookup localhost
ping localhost
```

### 3. File System Issues

#### Check Permissions
```bash
# Check file permissions
ls -la bootstrap.sh
ls -la minibb.db

# Fix permissions
chmod +x bootstrap.sh
chmod 644 minibb.db
```

#### Check Disk Space
```bash
# Check available space
df -h
du -sh minibb-main/

# Clean up if needed
make clean
rm -rf web/node_modules
```

## 🆘 Getting Help

### 1. Self-Diagnosis Steps
1. **Run test suite:** `./test-bootstrap.sh`
2. **Check logs:** `make tail-log`
3. **Try quick fixes:** Reset database, clean dependencies
4. **Search this guide:** Use Ctrl+F to find specific errors
5. **Check documentation:** Review relevant guide sections

### 2. Information to Gather
When seeking help, provide:
- **Operating system:** macOS, Linux, Windows
- **Software versions:** Go, Node.js, npm versions
- **Error messages:** Full error text
- **Steps to reproduce:** What you did before the error
- **Environment:** Output of `env` command
- **Test results:** Output of `./test-bootstrap.sh`

### 3. Community Resources
- **GitHub Issues:** Report bugs and feature requests
- **Documentation:** Check all guide files
- **Stack Overflow:** Search for similar issues
- **Discord/Slack:** Join development communities

## 📚 Prevention Tips

### 1. Best Practices
- **Regular backups:** Database is backed up automatically
- **Version control:** Commit changes frequently
- **Clean environment:** Keep dependencies updated
- **Test before switching:** Run tests before changing variants

### 2. Maintenance
```bash
# Weekly maintenance
go mod tidy
cd web && npm audit fix && cd ..
./test-bootstrap.sh

# Monthly maintenance
go clean -modcache
cd web && rm -rf node_modules && npm install && cd ..
```

### 3. Monitoring
```bash
# Check application health
curl -s http://localhost:8080/api/health

# Monitor resources
top -p $(pgrep minibb)

# Check logs regularly
tail -f dev.log
```

---

**Still having issues?** Follow the self-diagnosis steps and gather the required information before seeking help. Most issues can be resolved with the solutions provided in this guide.

Remember: The best debugging approach is systematic. Start with the quick fixes, then work through the specific issue sections methodically.

Happy debugging! 🚀