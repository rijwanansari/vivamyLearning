# Simple API Test
$baseUrl = "http://localhost:8080"

Write-Host "Testing Health Endpoint..." -ForegroundColor Yellow

try {
    $response = Invoke-WebRequest -Uri "$baseUrl/ping" -Method GET -UseBasicParsing
    Write-Host "✅ Health Check Success!" -ForegroundColor Green
    Write-Host "Status: $($response.StatusCode)" -ForegroundColor Cyan
    Write-Host "Response: $($response.Content)" -ForegroundColor White
}
catch {
    Write-Host "❌ Health Check Failed!" -ForegroundColor Red
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`nTesting Registration..." -ForegroundColor Yellow

$registerData = @{
    name = "Test User"
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

try {
    $response = Invoke-WebRequest -Uri "$baseUrl/api/v1/auth/register" -Method POST -Body $registerData -ContentType "application/json" -UseBasicParsing
    Write-Host "✅ Registration Success!" -ForegroundColor Green
    Write-Host "Status: $($response.StatusCode)" -ForegroundColor Cyan
    Write-Host "Response: $($response.Content)" -ForegroundColor White
}
catch {
    Write-Host "❌ Registration Failed!" -ForegroundColor Red
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "Status: $($_.Exception.Response.StatusCode)" -ForegroundColor Red
}
