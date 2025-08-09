# 🚀 VivaLearning - How to Run Frontend & Backend

This guide will help you start both the Go backend and React frontend for the VivaLearning platform.

## 📋 Prerequisites

Before starting, make sure you have:

- **Go 1.24+** installed
- **Node.js 18+** installed  
- **PostgreSQL** database running
- **Git** installed

## 🗄️ Database Setup

1. **Start PostgreSQL** (if not already running)
2. **Create Database** (if not exists):
   ```sql
   CREATE DATABASE Mysample;
   ```
3. **Update .env file** with your database credentials:
   ```bash
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASS=admin
   DB_SCHEMA=Mysample
   ```

## 🔧 Backend Setup (Go)

### Option 1: Quick Start Script
```powershell
# Run the automated setup
.\setup-backend.ps1
```

### Option 2: Manual Setup

1. **Navigate to project root**:
   ```powershell
   cd e:\Projects\vivaLearning
   ```

2. **Install Go dependencies**:
   ```powershell
   go mod download
   ```

3. **Build the application**:
   ```powershell
   go build -o vivaLearning.exe .
   ```

4. **Start the backend server**:
   ```powershell
   .\vivaLearning.exe serve
   ```

5. **Verify backend is running**:
   - Open browser: `http://localhost:8080/ping`
   - Should see: `{"message": "Pong from viva Learning"}`

## 🎨 Frontend Setup (React)

### Option 1: Quick Start Script
```powershell
# Run the automated setup
.\setup-frontend.ps1
```

### Option 2: Manual Setup

1. **Navigate to frontend directory**:
   ```powershell
   cd frontend
   ```

2. **Install dependencies**:
   ```powershell
   npm install
   ```

3. **Start the development server**:
   ```powershell
   npm run dev
   ```

4. **Verify frontend is running**:
   - Open browser: `http://localhost:3000`
   - Should see the VivaLearning login page

## 🚦 Running Both Services

### Method 1: Two Terminal Windows

**Terminal 1 (Backend)**:
```powershell
cd e:\Projects\vivaLearning
.\vivaLearning.exe serve
```

**Terminal 2 (Frontend)**:
```powershell
cd e:\Projects\vivaLearning\frontend
npm run dev
```

### Method 2: PowerShell Background Jobs

```powershell
# Start backend in background
Start-Job -ScriptBlock { 
    Set-Location "e:\Projects\vivaLearning"
    .\vivaLearning.exe serve 
} -Name "Backend"

# Start frontend in background
Start-Job -ScriptBlock { 
    Set-Location "e:\Projects\vivaLearning\frontend"
    npm run dev 
} -Name "Frontend"

# Check job status
Get-Job
```

### Method 3: All-in-One Start Script

Create `start-all.ps1`:
```powershell
Write-Host "🚀 Starting VivaLearning Full Stack..." -ForegroundColor Green

# Start Backend
Write-Host "Starting Go Backend..." -ForegroundColor Yellow
Start-Process -FilePath ".\vivaLearning.exe" -ArgumentList "serve" -WindowStyle Normal

# Wait for backend to start
Start-Sleep -Seconds 3

# Start Frontend  
Write-Host "Starting React Frontend..." -ForegroundColor Yellow
Set-Location "frontend"
Start-Process -FilePath "npm" -ArgumentList "run", "dev" -WindowStyle Normal

Write-Host "✅ Both services starting..." -ForegroundColor Green
Write-Host "Backend: http://localhost:8080" -ForegroundColor Cyan
Write-Host "Frontend: http://localhost:3000" -ForegroundColor Cyan
```

## 🔍 Verification Steps

### 1. Check Backend Health
```powershell
curl http://localhost:8080/ping
# OR
Invoke-RestMethod -Uri "http://localhost:8080/ping"
```

### 2. Check Frontend
```powershell
curl http://localhost:3000
# OR open browser at http://localhost:3000
```

### 3. Test API Integration
```powershell
# Run the comprehensive test
.\test-api.ps1
```

## 🌐 Access Points

Once both services are running:

| Service | URL | Purpose |
|---------|-----|---------|
| **Frontend** | `http://localhost:3000` | Main web application |
| **Backend API** | `http://localhost:8080/api/v1` | REST API endpoints |
| **Health Check** | `http://localhost:8080/ping` | Backend status |

## 📱 Using the Application

1. **Open Frontend**: Navigate to `http://localhost:3000`
2. **Register**: Create a new account
3. **Login**: Sign in with your credentials
4. **Explore**: Browse courses, create content, track progress

## 🛠️ Development Workflow

### Hot Reload Development
- **Frontend**: Changes auto-reload in browser
- **Backend**: Restart server after Go code changes

### Build for Production

**Backend**:
```powershell
go build -o vivaLearning.exe .
```

**Frontend**:
```powershell
cd frontend
npm run build
```

## 🔧 Troubleshooting

### Backend Issues

**Port 8080 already in use**:
```powershell
netstat -ano | findstr :8080
# Kill the process if needed
taskkill /PID <PID> /F
```

**Database connection issues**:
- Check PostgreSQL is running
- Verify credentials in `.env` file
- Check database exists

### Frontend Issues

**Port 3000 already in use**:
- Vite will automatically use next available port
- Or manually specify: `npm run dev -- --port 3001`

**Module not found errors**:
```powershell
rm -rf node_modules package-lock.json
npm install
```

### Integration Issues

**CORS errors**:
- Backend has CORS enabled for frontend
- Check if both services are running on correct ports

**API calls failing**:
- Verify backend is accessible at `http://localhost:8080`
- Check network tab in browser dev tools

## 🚀 Quick Start Commands

**Start Everything** (in project root):
```powershell
# Terminal 1: Backend
.\vivaLearning.exe serve

# Terminal 2: Frontend
cd frontend; npm run dev
```

**Stop Everything**:
```powershell
# Ctrl+C in both terminals
# OR kill background jobs
Get-Job | Stop-Job
```

## 📞 Support

If you encounter issues:

1. **Check logs** in both terminal windows
2. **Verify prerequisites** are installed
3. **Check ports** 8080 and 3000 are available
4. **Review .env file** for correct database settings
5. **Test individual components** first

---

**🎉 Happy Learning with VivaLearning!**
