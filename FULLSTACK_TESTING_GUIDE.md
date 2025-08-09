# VivaLearning Full Stack Testing Guide

## 🏗️ Architecture Overview

**Frontend:** React 18 + TypeScript + Vite + Tailwind CSS  
**Backend:** Go + Echo + GORM + PostgreSQL  
**Communication:** REST APIs with JWT Authentication

## 🚀 Complete Testing Workflow

### 1. **Backend Setup & Testing**

#### Start the Go Backend:
```powershell
# Terminal 1: Start Go backend
cd e:\Projects\vivaLearning
.\vivaLearning.exe serve
```

Expected output:
```
=== Configuration Loaded from .env ===
App Name: vivaLearning
App Port: 8080
DB Host: localhost
DB Port: 5432
...
Server started on port 8080
```

#### Test Backend APIs:
```powershell
# Test health endpoint
curl http://localhost:8080/ping

# Expected: {"message": "Pong from viva Learning"}
```

### 2. **Frontend Setup & Testing**

#### Install Dependencies:
```powershell
# Terminal 2: Setup React frontend
cd e:\Projects\vivaLearning\frontend
npm install
```

#### Start React Development Server:
```powershell
npm run dev
```

Expected output:
```
> viva-learning-frontend@0.0.0 dev
> vite

  VITE v4.4.5  ready in 500 ms

  ➜  Local:   http://localhost:3000/
  ➜  Network: use --host to expose
```

### 3. **Full Stack Integration Testing**

#### Test the Complete Flow:

1. **Open Browser:** Navigate to `http://localhost:3000`
2. **Registration Test:**
   - Click "Create new account"
   - Fill form with test data
   - Submit and verify token storage
3. **Login Test:**
   - Use credentials: `test@example.com` / `password123`
   - Verify redirect to dashboard
4. **Protected Routes:**
   - Access dashboard, courses, my courses
   - Verify navigation works
5. **API Integration:**
   - Create a course
   - Add lessons
   - Test enrollment

## 🔧 Development Tools

### **VS Code Extensions for React + TypeScript:**
```json
{
  "recommendations": [
    "bradlc.vscode-tailwindcss",
    "ms-vscode.vscode-typescript-next",
    "esbenp.prettier-vscode",
    "ms-vscode.vscode-eslint",
    "formulahendry.auto-rename-tag",
    "christian-kohler.path-intellisense"
  ]
}
```

### **Browser DevTools Setup:**
- **React Developer Tools** - Chrome/Firefox extension
- **Redux DevTools** - For state management debugging
- **Network Tab** - Monitor API calls to Go backend

## 🧪 Testing Scenarios

### **Authentication Flow:**
```typescript
// Test sequence:
1. Register new user → Backend: POST /api/v1/auth/register
2. Verify JWT token stored in cookies
3. Access protected route → Frontend: Check token validity
4. Logout → Clear tokens and redirect
```

### **Course Management Flow:**
```typescript
// Test sequence:
1. Create course → Backend: POST /api/v1/courses
2. View course list → Backend: GET /api/v1/courses  
3. Update course → Backend: PUT /api/v1/courses/:id
4. Add lessons → Backend: POST /api/v1/courses/:id/lessons
```

### **Learning Flow:**
```typescript
// Test sequence:
1. Browse courses → Backend: GET /api/v1/courses/search
2. Enroll in course → Backend: POST /api/v1/courses/:id/enroll
3. Watch lesson → Frontend: Video player + progress tracking
4. Update progress → Backend: POST /api/v1/lessons/progress
```

## 🔍 Debugging Guide

### **Frontend Debugging:**

#### Check React App Status:
```javascript
// Browser Console Commands:
console.log('React Version:', React.version);
console.log('Auth Token:', document.cookie);
console.log('User Data:', localStorage.getItem('user'));
```

#### Common Frontend Issues:
```typescript
// CORS Issues - Check Vite proxy configuration:
// vite.config.ts
server: {
  proxy: {
    '/api': 'http://localhost:8080'
  }
}

// Authentication Issues - Check token storage:
import Cookies from 'js-cookie';
console.log('Token:', Cookies.get('access_token'));
```

### **Backend Debugging:**

#### Check Go Server Status:
```powershell
# Check if server is running
netstat -an | findstr :8080

# Check database connection
# Look for database connection logs in terminal
```

#### API Response Testing:
```powershell
# Test with curl
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

## 📊 Performance Testing

### **Frontend Performance:**
```typescript
// React DevTools Profiler
// Check component render times
// Monitor bundle size with npm run build

// Lighthouse audit in Chrome
// Check Core Web Vitals
```

### **Backend Performance:**
```go
// Add logging middleware in Go
// Monitor API response times
// Check database query performance
```

## 🛠️ Development Workflow

### **Hot Reload Setup:**
```yaml
# Development environment:
Backend (Go):
  - File: .go files
  - Tool: go run main.go serve
  - Port: 8080
  - Auto-reload: Use 'air' for hot reload

Frontend (React):
  - Files: .tsx, .ts, .css files  
  - Tool: vite dev server
  - Port: 3000
  - Auto-reload: Built-in with Vite
```

### **Code Quality Tools:**
```json
// package.json scripts:
{
  "scripts": {
    "dev": "vite",
    "build": "tsc && vite build",
    "lint": "eslint . --ext ts,tsx",
    "preview": "vite preview",
    "type-check": "tsc --noEmit"
  }
}
```

## 🚦 Production Testing

### **Build Testing:**
```powershell
# Frontend production build
cd frontend
npm run build
npm run preview

# Backend production build
cd ..
go build -o vivaLearning.exe .
```

### **Environment Configuration:**
```env
# .env (Backend)
APP_PORT=8080
DB_HOST=localhost
JWT_ACCESS_TOKEN_SECRET=your-secret

# Frontend - API URL in production
VITE_API_URL=https://your-api-domain.com/api/v1
```

## 📱 Mobile Testing

### **Responsive Design Testing:**
```css
/* Test breakpoints in browser DevTools */
Mobile: 375px - 768px
Tablet: 768px - 1024px  
Desktop: 1024px+

/* Test touch interactions */
/* Verify mobile navigation */
```

## 🔐 Security Testing

### **Authentication Security:**
```typescript
// Test scenarios:
1. Invalid JWT tokens
2. Expired tokens  
3. CSRF protection
4. XSS prevention
5. Secure cookie settings
```

### **API Security:**
```go
// Backend security checks:
1. Input validation
2. SQL injection prevention  
3. Rate limiting
4. CORS configuration
```

## 📈 Monitoring & Analytics

### **Frontend Monitoring:**
```typescript
// Error tracking
window.addEventListener('error', (e) => {
  console.error('Frontend Error:', e);
});

// Performance monitoring  
performance.mark('app-start');
```

### **Backend Monitoring:**
```go
// Request logging
// Database connection monitoring
// Error rate tracking
```

## 🎯 Next Steps

1. **Complete Integration Testing** - Test all API endpoints with frontend
2. **Add Unit Tests** - Jest for React, Go testing for backend  
3. **E2E Testing** - Cypress or Playwright for full user flows
4. **Performance Optimization** - Bundle analysis, database indexing
5. **Security Hardening** - Penetration testing, dependency scanning
6. **Production Deployment** - Docker containers, CI/CD pipeline

## 📞 Troubleshooting

### **Common Issues:**

#### Frontend won't start:
```powershell
# Clear node_modules and reinstall
rm -rf node_modules package-lock.json
npm install
```

#### Backend connection issues:
```powershell
# Check .env file configuration
# Verify database is running
# Check firewall settings
```

#### CORS errors:
```typescript
// Update vite.config.ts proxy settings
// Check backend CORS middleware
```

---

**Ready to test your full stack application! 🚀**

Run both servers simultaneously and test the complete user journey from registration to course completion.
