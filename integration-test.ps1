# VivaLearning Full Stack Integration Test

Write-Host "🚀 Testing VivaLearning Full Stack Application" -ForegroundColor Green
Write-Host "=============================================" -ForegroundColor Green
Write-Host ""

# Check Backend
Write-Host "1. 🔧 Go Backend Health Check..." -ForegroundColor Yellow
try {
    $backendResponse = Invoke-WebRequest -Uri "http://localhost:8080/ping" -UseBasicParsing
    Write-Host "   ✅ Backend Running: $($backendResponse.StatusCode)" -ForegroundColor Green
    Write-Host "   📍 URL: http://localhost:8080" -ForegroundColor Cyan
} catch {
    Write-Host "   ❌ Backend Error: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

# Check Frontend  
Write-Host "2. ⚛️  React Frontend Health Check..." -ForegroundColor Yellow
try {
    $frontendResponse = Invoke-WebRequest -Uri "http://localhost:3000" -UseBasicParsing
    Write-Host "   ✅ Frontend Running: $($frontendResponse.StatusCode)" -ForegroundColor Green
    Write-Host "   📍 URL: http://localhost:3000" -ForegroundColor Cyan
} catch {
    Write-Host "   ❌ Frontend Error: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

# Test API Integration
Write-Host "3. 🔗 API Integration Test..." -ForegroundColor Yellow

# Register a test user
$testUser = @{
    name = "Integration Test User"
    email = "integration$(Get-Random)@test.com"
    password = "password123"
} | ConvertTo-Json

try {
    # Test registration via frontend proxy
    $regResponse = Invoke-WebRequest -Uri "http://localhost:3000/api/v1/auth/register" -Method POST -Body $testUser -ContentType "application/json" -UseBasicParsing
    Write-Host "   ✅ User Registration: $($regResponse.StatusCode)" -ForegroundColor Green
    
    $regData = $regResponse.Content | ConvertFrom-Json
    $token = $regData.data.access_token
    Write-Host "   🎫 Token acquired: $($token.Substring(0,20))..." -ForegroundColor Cyan
    
    # Test authenticated endpoint
    $authHeaders = @{ "Authorization" = "Bearer $token" }
    $coursesResponse = Invoke-WebRequest -Uri "http://localhost:3000/api/v1/my/courses" -Headers $authHeaders -UseBasicParsing
    Write-Host "   ✅ Authenticated API: $($coursesResponse.StatusCode)" -ForegroundColor Green
    
} catch {
    Write-Host "   ❌ API Error: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host ""
Write-Host "🎉 Full Stack Test Complete!" -ForegroundColor Green
Write-Host "📱 Open http://localhost:3000 in your browser" -ForegroundColor Yellow
