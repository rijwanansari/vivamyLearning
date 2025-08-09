# VivaLearning Full Stack Integration Test

Write-Host "🚀 VivaLearning Full Stack Integration Test" -ForegroundColor Green
Write-Host "=========================================" -ForegroundColor Green
Write-Host ""

# Test Backend (Go API)
Write-Host "🔧 Testing Backend (Go API)..." -ForegroundColor Yellow

try {
    $backendHealth = Invoke-RestMethod -Uri "http://localhost:8080/ping" -Method GET
    Write-Host "✅ Backend Health Check: SUCCESS" -ForegroundColor Green
    Write-Host "   Response: $($backendHealth.message)" -ForegroundColor White
} catch {
    Write-Host "❌ Backend Health Check: FAILED" -ForegroundColor Red
    Write-Host "   Error: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "   Make sure the Go backend is running on port 8080" -ForegroundColor Yellow
    exit 1
}

# Test Frontend (React)
Write-Host ""
Write-Host "🎨 Testing Frontend (React)..." -ForegroundColor Yellow

try {
    $frontendHealth = Invoke-WebRequest -Uri "http://localhost:3000" -Method GET -UseBasicParsing
    Write-Host "✅ Frontend Health Check: SUCCESS" -ForegroundColor Green
    Write-Host "   Status Code: $($frontendHealth.StatusCode)" -ForegroundColor White
    Write-Host "   Content Length: $($frontendHealth.Content.Length) bytes" -ForegroundColor White
} catch {
    Write-Host "❌ Frontend Health Check: FAILED" -ForegroundColor Red
    Write-Host "   Error: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "   Make sure the React frontend is running on port 3000" -ForegroundColor Yellow
    exit 1
}

Write-Host ""
Write-Host "🎉 Both Frontend and Backend are Running Successfully!" -ForegroundColor Green
Write-Host ""
Write-Host "🌐 Access Points:" -ForegroundColor Cyan
Write-Host "   Frontend (React): http://localhost:3000" -ForegroundColor White
Write-Host "   Backend API (Go): http://localhost:8080/api/v1" -ForegroundColor White
Write-Host "   Health Check: http://localhost:8080/ping" -ForegroundColor White
