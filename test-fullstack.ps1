# VivaLearning Full Stack Test Script
# Tests both React frontend and Go backend integration

Write-Host "🚀 VivaLearning Full Stack Testing" -ForegroundColor Green
Write-Host "====================================" -ForegroundColor Green
Write-Host ""

# Test 1: Check if backend is running
Write-Host "🔍 Testing Backend (Go)..." -ForegroundColor Yellow
try {
    $backendResponse = Invoke-RestMethod -Uri "http://localhost:8080/ping" -Method GET -TimeoutSec 5
    Write-Host "✅ Backend is running!" -ForegroundColor Green
    Write-Host "   Response: $($backendResponse.message)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Backend not running on port 8080" -ForegroundColor Red
    Write-Host "   Please start: .\vivaLearning.exe serve" -ForegroundColor Yellow
    Write-Host ""
}

# Test 2: Check if frontend is accessible
Write-Host ""
Write-Host "🔍 Testing Frontend (React)..." -ForegroundColor Yellow
try {
    $frontendResponse = Invoke-WebRequest -Uri "http://localhost:3000" -Method GET -TimeoutSec 5 -UseBasicParsing
    Write-Host "✅ Frontend is running!" -ForegroundColor Green
    Write-Host "   Status: $($frontendResponse.StatusCode)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Frontend not running on port 3000" -ForegroundColor Red
    Write-Host "   Please start: cd frontend && npm run dev" -ForegroundColor Yellow
    Write-Host ""
}

# Test 3: Test API integration via frontend proxy
Write-Host ""
Write-Host "🔍 Testing API Integration..." -ForegroundColor Yellow
try {
    $apiResponse = Invoke-RestMethod -Uri "http://localhost:3000/api/ping" -Method GET -TimeoutSec 5
    Write-Host "✅ API proxy working!" -ForegroundColor Green
    Write-Host "   Frontend can communicate with backend" -ForegroundColor Cyan
} catch {
    Write-Host "⚠️  API proxy may need configuration" -ForegroundColor Yellow
    Write-Host "   Check vite.config.ts proxy settings" -ForegroundColor Gray
}

# Test 4: Check development environment
Write-Host ""
Write-Host "🔍 Development Environment Check..." -ForegroundColor Yellow

# Check Node.js
try {
    $nodeVersion = node --version
    Write-Host "✅ Node.js: $nodeVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Node.js not found" -ForegroundColor Red
}

# Check Go
try {
    $goVersion = go version
    Write-Host "✅ Go: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Go not found" -ForegroundColor Red
}

# Test 5: Database connection (through backend API)
Write-Host ""
Write-Host "🔍 Testing Database Connection..." -ForegroundColor Yellow
try {
    # Try to register a test user to check database connectivity
    $testUser = @{
        name = "Test User $(Get-Date -Format 'HHmmss')"
        email = "test$(Get-Date -Format 'HHmmss')@example.com"
        password = "password123"
    } | ConvertTo-Json

    $registerResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/register" -Method POST -Body $testUser -ContentType "application/json" -TimeoutSec 10
    Write-Host "✅ Database connection working!" -ForegroundColor Green
    Write-Host "   Successfully created test user" -ForegroundColor Cyan
} catch {
    Write-Host "⚠️  Database connection issue" -ForegroundColor Yellow
    Write-Host "   Error: $($_.Exception.Message)" -ForegroundColor Gray
}

Write-Host ""
Write-Host "📋 Testing Summary:" -ForegroundColor Magenta
Write-Host "===================" -ForegroundColor Magenta
Write-Host "• Backend (Go):        http://localhost:8080" -ForegroundColor White
Write-Host "• Frontend (React):    http://localhost:3000" -ForegroundColor White
Write-Host "• API Integration:     /api proxy to backend" -ForegroundColor White
Write-Host "• Database:            PostgreSQL via GORM" -ForegroundColor White
Write-Host ""

Write-Host "🎯 Next Steps:" -ForegroundColor Cyan
Write-Host "1. Open browser: http://localhost:3000" -ForegroundColor White
Write-Host "2. Test user registration and login" -ForegroundColor White
Write-Host "3. Create a course and add lessons" -ForegroundColor White
Write-Host "4. Test enrollment and progress tracking" -ForegroundColor White
Write-Host ""

Write-Host "🛠️  If services aren't running:" -ForegroundColor Yellow
Write-Host "Backend:  .\vivaLearning.exe serve" -ForegroundColor Gray
Write-Host "Frontend: cd frontend && npm run dev" -ForegroundColor Gray
Write-Host ""

Write-Host "🎉 Full Stack Testing Complete!" -ForegroundColor Green
