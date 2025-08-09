# VivaLearning - Start All Services

Write-Host "🚀 Starting VivaLearning Full Stack Application..." -ForegroundColor Green
Write-Host "=================================================" -ForegroundColor Green
Write-Host ""

# Check if executable exists
if (-not (Test-Path "vivaLearning.exe")) {
    Write-Host "❌ vivaLearning.exe not found. Building..." -ForegroundColor Yellow
    go build -o vivaLearning.exe .
    if (-not (Test-Path "vivaLearning.exe")) {
        Write-Host "❌ Failed to build backend. Please check Go installation." -ForegroundColor Red
        exit 1
    }
}

# Check if frontend dependencies are installed
if (-not (Test-Path "frontend\node_modules")) {
    Write-Host "📦 Installing frontend dependencies..." -ForegroundColor Yellow
    Set-Location "frontend"
    npm install
    Set-Location ".."
    if (-not (Test-Path "frontend\node_modules")) {
        Write-Host "❌ Failed to install frontend dependencies." -ForegroundColor Red
        exit 1
    }
}

Write-Host "🔧 Starting Backend (Go API Server)..." -ForegroundColor Yellow
Start-Process -FilePath ".\vivaLearning.exe" -ArgumentList "serve" -WindowStyle Normal

Write-Host "⏱️  Waiting for backend to initialize..." -ForegroundColor Gray
Start-Sleep -Seconds 5

# Test if backend is responding
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/ping" -Method GET -TimeoutSec 5
    Write-Host "✅ Backend started successfully!" -ForegroundColor Green
    Write-Host "   Response: $($response.message)" -ForegroundColor White
} catch {
    Write-Host "⚠️  Backend may still be starting... continuing anyway" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "🎨 Starting Frontend (React Development Server)..." -ForegroundColor Yellow
Set-Location "frontend"
Start-Process -FilePath "npm" -ArgumentList "run", "dev" -WindowStyle Normal
Set-Location ".."

Write-Host ""
Write-Host "🎉 VivaLearning is starting up!" -ForegroundColor Green
Write-Host "================================" -ForegroundColor Green
Write-Host ""
Write-Host "🌐 Access your application at:" -ForegroundColor Cyan
Write-Host "   Frontend:    http://localhost:3000" -ForegroundColor White
Write-Host "   Backend API: http://localhost:8080/api/v1" -ForegroundColor White
Write-Host "   Health Check: http://localhost:8080/ping" -ForegroundColor White
Write-Host ""
Write-Host "📝 Test Credentials:" -ForegroundColor Cyan
Write-Host "   Email: test@example.com" -ForegroundColor White
Write-Host "   Password: password123" -ForegroundColor White
Write-Host ""
Write-Host "⏸️  To stop all services:" -ForegroundColor Yellow
Write-Host "   Press Ctrl+C in both terminal windows" -ForegroundColor White
Write-Host "   Or run: Get-Process | Where-Object {`$_.ProcessName -like '*node*' -or `$_.ProcessName -like '*vivaLearning*'} | Stop-Process" -ForegroundColor White
Write-Host ""
Write-Host "🚀 Happy Learning!" -ForegroundColor Green
