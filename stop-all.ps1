# VivaLearning - Stop All Services

Write-Host "🛑 Stopping VivaLearning Services..." -ForegroundColor Red
Write-Host "====================================" -ForegroundColor Red
Write-Host ""

# Stop Node.js processes (frontend)
Write-Host "🎨 Stopping Frontend (React)..." -ForegroundColor Yellow
try {
    $nodeProcesses = Get-Process | Where-Object {$_.ProcessName -like "*node*"}
    if ($nodeProcesses) {
        $nodeProcesses | Stop-Process -Force
        Write-Host "✅ Frontend stopped" -ForegroundColor Green
    } else {
        Write-Host "ℹ️  No frontend processes found" -ForegroundColor Gray
    }
} catch {
    Write-Host "⚠️  Error stopping frontend: $($_.Exception.Message)" -ForegroundColor Yellow
}

# Stop Go backend processes
Write-Host "🔧 Stopping Backend (Go)..." -ForegroundColor Yellow
try {
    $goProcesses = Get-Process | Where-Object {$_.ProcessName -like "*vivaLearning*"}
    if ($goProcesses) {
        $goProcesses | Stop-Process -Force
        Write-Host "✅ Backend stopped" -ForegroundColor Green
    } else {
        Write-Host "ℹ️  No backend processes found" -ForegroundColor Gray
    }
} catch {
    Write-Host "⚠️  Error stopping backend: $($_.Exception.Message)" -ForegroundColor Yellow
}

# Stop any npm processes
Write-Host "📦 Stopping npm processes..." -ForegroundColor Yellow
try {
    $npmProcesses = Get-Process | Where-Object {$_.ProcessName -like "*npm*"}
    if ($npmProcesses) {
        $npmProcesses | Stop-Process -Force
        Write-Host "✅ npm processes stopped" -ForegroundColor Green
    } else {
        Write-Host "ℹ️  No npm processes found" -ForegroundColor Gray
    }
} catch {
    Write-Host "⚠️  Error stopping npm: $($_.Exception.Message)" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "🔍 Checking if ports are free..." -ForegroundColor Yellow

# Check if ports are free
$port8080 = netstat -an | findstr ":8080"
$port3000 = netstat -an | findstr ":3000"

if (-not $port8080) {
    Write-Host "✅ Port 8080 is free" -ForegroundColor Green
} else {
    Write-Host "⚠️  Port 8080 may still be in use" -ForegroundColor Yellow
}

if (-not $port3000) {
    Write-Host "✅ Port 3000 is free" -ForegroundColor Green
} else {
    Write-Host "⚠️  Port 3000 may still be in use" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "🎉 All VivaLearning services have been stopped!" -ForegroundColor Green
Write-Host ""
Write-Host "🚀 To start again, run: .\start-all.ps1" -ForegroundColor Cyan
