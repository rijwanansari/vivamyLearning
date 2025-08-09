# Simple Backend API Response Test

Write-Host "🔍 Testing Backend API Response Format" -ForegroundColor Cyan
Write-Host "=====================================" -ForegroundColor Cyan

$testLogin = @{
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

Write-Host "Testing with payload: $testLogin" -ForegroundColor Yellow

try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/login" -Method POST -Body $testLogin -ContentType "application/json"
    
    Write-Host ""
    Write-Host "✅ Raw Response:" -ForegroundColor Green
    Write-Host ($response | ConvertTo-Json -Depth 4) -ForegroundColor White
    
    Write-Host ""
    Write-Host "🔍 Response Analysis:" -ForegroundColor Cyan
    Write-Host "- Response Type: $($response.GetType().Name)" -ForegroundColor White
    Write-Host "- Has 'success' property: $($response.PSObject.Properties.Name -contains 'success')" -ForegroundColor White
    Write-Host "- Has 'message' property: $($response.PSObject.Properties.Name -contains 'message')" -ForegroundColor White
    Write-Host "- Has 'data' property: $($response.PSObject.Properties.Name -contains 'data')" -ForegroundColor White
    
    if ($response.data) {
        Write-Host "- Data has 'user' property: $($response.data.PSObject.Properties.Name -contains 'user')" -ForegroundColor White
        Write-Host "- Data has 'access_token' property: $($response.data.PSObject.Properties.Name -contains 'access_token')" -ForegroundColor White
        Write-Host "- Data has 'refresh_token' property: $($response.data.PSObject.Properties.Name -contains 'refresh_token')" -ForegroundColor White
    }
    
} catch {
    Write-Host "❌ Login Failed: $($_.Exception.Message)" -ForegroundColor Red
    
    if ($_.Exception.Response) {
        $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
        $errorBody = $reader.ReadToEnd()
        Write-Host "Error Response: $errorBody" -ForegroundColor Red
    }
}
