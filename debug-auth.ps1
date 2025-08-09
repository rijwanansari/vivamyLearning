# VivaLearning - End-to-End Authentication Debug Test

Write-Host "🔍 VivaLearning Authentication Debug Test" -ForegroundColor Cyan
Write-Host "=========================================" -ForegroundColor Cyan
Write-Host ""

# Test 1: Backend Health Check
Write-Host "1. 🔧 Testing Backend Health..." -ForegroundColor Yellow
try {
    $healthResponse = Invoke-RestMethod -Uri "http://localhost:8080/ping" -Method GET
    Write-Host "✅ Backend Health: SUCCESS" -ForegroundColor Green
    Write-Host "   Response: $($healthResponse | ConvertTo-Json)" -ForegroundColor White
} catch {
    Write-Host "❌ Backend Health: FAILED" -ForegroundColor Red
    Write-Host "   Error: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

# Test 2: Frontend Health Check
Write-Host ""
Write-Host "2. 🎨 Testing Frontend Health..." -ForegroundColor Yellow
try {
    $frontendResponse = Invoke-WebRequest -Uri "http://localhost:3000" -Method GET -UseBasicParsing -TimeoutSec 10
    Write-Host "✅ Frontend Health: SUCCESS" -ForegroundColor Green
    Write-Host "   Status: $($frontendResponse.StatusCode)" -ForegroundColor White
} catch {
    Write-Host "❌ Frontend Health: FAILED" -ForegroundColor Red
    Write-Host "   Error: $($_.Exception.Message)" -ForegroundColor Red
}

# Test 3: User Registration (Backend API)
Write-Host ""
Write-Host "3. 📝 Testing User Registration..." -ForegroundColor Yellow

$testUser = @{
    name = "Debug Test User"
    email = "debug.test@example.com"
    password = "debugpass123"
} | ConvertTo-Json

Write-Host "   Request Data: $testUser" -ForegroundColor Gray

try {
    $registerResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/register" -Method POST -Body $testUser -ContentType "application/json"
    Write-Host "✅ Registration: SUCCESS" -ForegroundColor Green
    Write-Host "   Full Response: $($registerResponse | ConvertTo-Json -Depth 3)" -ForegroundColor White
    
    if ($registerResponse.data -and $registerResponse.data.access_token) {
        $accessToken = $registerResponse.data.access_token
        Write-Host "   Access Token Length: $($accessToken.Length) characters" -ForegroundColor Cyan
        Write-Host "   Token Preview: $($accessToken.Substring(0, [Math]::Min(50, $accessToken.Length)))..." -ForegroundColor Cyan
    }
} catch {
    Write-Host "⚠️  Registration Failed (user might exist): $($_.Exception.Message)" -ForegroundColor Yellow
    
    # Try login instead
    Write-Host "   Trying login instead..." -ForegroundColor Gray
    
    $loginData = @{
        email = "debug.test@example.com"
        password = "debugpass123"
    } | ConvertTo-Json
    
    try {
        $loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/login" -Method POST -Body $loginData -ContentType "application/json"
        Write-Host "✅ Login: SUCCESS" -ForegroundColor Green
        Write-Host "   Full Response: $($loginResponse | ConvertTo-Json -Depth 3)" -ForegroundColor White
        
        if ($loginResponse.data -and $loginResponse.data.access_token) {
            $accessToken = $loginResponse.data.access_token
            Write-Host "   Access Token Length: $($accessToken.Length) characters" -ForegroundColor Cyan
        }
    } catch {
        Write-Host "❌ Login Also Failed: $($_.Exception.Message)" -ForegroundColor Red
        if ($_.Exception.Response) {
            $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
            $responseBody = $reader.ReadToEnd()
            Write-Host "   Response Body: $responseBody" -ForegroundColor Red
        }
    }
}

# Test 4: Frontend to Backend CORS Test
Write-Host ""
Write-Host "4. 🌐 Testing CORS Configuration..." -ForegroundColor Yellow

try {
    $corsHeaders = @{
        "Origin" = "http://localhost:3000"
        "Access-Control-Request-Method" = "POST"
        "Access-Control-Request-Headers" = "Content-Type, Authorization"
    }
    
    $corsResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/auth/login" -Method OPTIONS -Headers $corsHeaders -UseBasicParsing
    Write-Host "✅ CORS Preflight: SUCCESS" -ForegroundColor Green
    Write-Host "   Status: $($corsResponse.StatusCode)" -ForegroundColor White
    
    $corsResponseHeaders = $corsResponse.Headers
    if ($corsResponseHeaders["Access-Control-Allow-Origin"]) {
        Write-Host "   Allow-Origin: $($corsResponseHeaders['Access-Control-Allow-Origin'])" -ForegroundColor Cyan
    }
    if ($corsResponseHeaders["Access-Control-Allow-Methods"]) {
        Write-Host "   Allow-Methods: $($corsResponseHeaders['Access-Control-Allow-Methods'])" -ForegroundColor Cyan
    }
    if ($corsResponseHeaders["Access-Control-Allow-Headers"]) {
        Write-Host "   Allow-Headers: $($corsResponseHeaders['Access-Control-Allow-Headers'])" -ForegroundColor Cyan
    }
} catch {
    Write-Host "❌ CORS Test: FAILED" -ForegroundColor Red
    Write-Host "   Error: $($_.Exception.Message)" -ForegroundColor Red
}

# Test 5: Direct API Login Test (Simulating Frontend)
Write-Host ""
Write-Host "5. 🔐 Testing Direct API Login (Frontend Simulation)..." -ForegroundColor Yellow

$frontendHeaders = @{
    "Content-Type" = "application/json"
    "Origin" = "http://localhost:3000"
    "Referer" = "http://localhost:3000"
}

$loginPayload = @{
    email = "debug.test@example.com"
    password = "debugpass123"
} | ConvertTo-Json

Write-Host "   Headers: $($frontendHeaders | ConvertTo-Json)" -ForegroundColor Gray
Write-Host "   Payload: $loginPayload" -ForegroundColor Gray

try {
    $apiLoginResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/login" -Method POST -Body $loginPayload -Headers $frontendHeaders
    Write-Host "✅ Frontend-style Login: SUCCESS" -ForegroundColor Green
    Write-Host "   Response Structure:" -ForegroundColor Cyan
    Write-Host "   - Success: $($apiLoginResponse.success)" -ForegroundColor White
    Write-Host "   - Message: $($apiLoginResponse.message)" -ForegroundColor White
    
    if ($apiLoginResponse.data) {
        Write-Host "   - Data Object: EXISTS" -ForegroundColor White
        if ($apiLoginResponse.data.user) {
            Write-Host "     - User ID: $($apiLoginResponse.data.user.id)" -ForegroundColor White
            Write-Host "     - User Email: $($apiLoginResponse.data.user.email)" -ForegroundColor White
        }
        if ($apiLoginResponse.data.access_token) {
            Write-Host "     - Access Token: EXISTS ($($apiLoginResponse.data.access_token.Length) chars)" -ForegroundColor White
        }
        if ($apiLoginResponse.data.refresh_token) {
            Write-Host "     - Refresh Token: EXISTS ($($apiLoginResponse.data.refresh_token.Length) chars)" -ForegroundColor White
        }
    }
    
} catch {
    Write-Host "❌ Frontend-style Login: FAILED" -ForegroundColor Red
    Write-Host "   Error: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "   Status Code: $($_.Exception.Response.StatusCode)" -ForegroundColor Red
    
    if ($_.Exception.Response) {
        try {
            $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
            $errorBody = $reader.ReadToEnd()
            Write-Host "   Error Response Body: $errorBody" -ForegroundColor Red
        } catch {
            Write-Host "   Could not read error response body" -ForegroundColor Red
        }
    }
}

# Test 6: Check Frontend API Configuration
Write-Host ""
Write-Host "6. 🔧 Frontend API Configuration Check..." -ForegroundColor Yellow
Write-Host "   Expected API Base URL: http://localhost:8080/api/v1" -ForegroundColor White
Write-Host "   Expected Frontend URL: http://localhost:3000" -ForegroundColor White

# Summary
Write-Host ""
Write-Host "📊 Debug Summary" -ForegroundColor Magenta
Write-Host "================" -ForegroundColor Magenta
Write-Host "If the backend login works but frontend shows 'Login failed':" -ForegroundColor Yellow
Write-Host "1. Check browser Developer Tools > Network tab for failed requests" -ForegroundColor White
Write-Host "2. Look for CORS errors in browser console" -ForegroundColor White
Write-Host "3. Verify the response format matches what authService expects" -ForegroundColor White
Write-Host "4. Check if the response.data structure is correct" -ForegroundColor White
Write-Host ""
Write-Host "🔍 Next Steps:" -ForegroundColor Cyan
Write-Host "1. Open browser at http://localhost:3000" -ForegroundColor White
Write-Host "2. Open Developer Tools (F12)" -ForegroundColor White
Write-Host "3. Try to login with: debug.test@example.com / debugpass123" -ForegroundColor White
Write-Host "4. Check Network tab and Console for errors" -ForegroundColor White
