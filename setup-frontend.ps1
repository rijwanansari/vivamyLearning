# VivaLearning Frontend Setup Script

Write-Host "🚀 Setting up VivaLearning Frontend..." -ForegroundColor Green
Write-Host ""

# Check if Node.js is installed
try {
    $nodeVersion = node --version
    Write-Host "✅ Node.js version: $nodeVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Node.js is not installed. Please install Node.js v18 or higher." -ForegroundColor Red
    Write-Host "   Download from: https://nodejs.org/" -ForegroundColor Yellow
    exit 1
}

# Navigate to frontend directory
Set-Location "frontend"

Write-Host "📦 Installing dependencies..." -ForegroundColor Yellow
npm install

if ($LASTEXITCODE -eq 0) {
    Write-Host "✅ Dependencies installed successfully!" -ForegroundColor Green
} else {
    Write-Host "❌ Failed to install dependencies." -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "🎉 Frontend setup complete!" -ForegroundColor Green
Write-Host ""
Write-Host "📋 Available Commands:" -ForegroundColor Cyan
Write-Host "   npm run dev     - Start development server (http://localhost:3000)" -ForegroundColor White
Write-Host "   npm run build   - Build for production" -ForegroundColor White
Write-Host "   npm run preview - Preview production build" -ForegroundColor White
Write-Host "   npm run lint    - Run ESLint" -ForegroundColor White
Write-Host ""
Write-Host "🚦 To start development:" -ForegroundColor Yellow
Write-Host "   cd frontend" -ForegroundColor White
Write-Host "   npm run dev" -ForegroundColor White
Write-Host ""
Write-Host "⚠️  Make sure the backend is running on http://localhost:8080" -ForegroundColor Yellow
