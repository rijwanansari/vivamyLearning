# VivaLearning API Test Script
# Make sure the server is running on localhost:8080 before running this script

$baseUrl = "http://localhost:8080"
$apiUrl = "$baseUrl/api/v1"

Write-Host "🚀 Starting VivaLearning API Tests..." -ForegroundColor Green
Write-Host "Base URL: $baseUrl" -ForegroundColor Yellow
Write-Host ""

# Global variables for storing tokens and IDs
$global:accessToken = ""
$global:userId = 0
$global:courseId = 0
$global:lessonId = 0

function Test-Endpoint {
    param(
        [string]$Method,
        [string]$Url,
        [string]$Description,
        [string]$Body = "",
        [hashtable]$Headers = @{}
    )
    
    Write-Host "🔍 Testing: $Description" -ForegroundColor Cyan
    Write-Host "   $Method $Url" -ForegroundColor Gray
    
    try {
        $response = if ($Body -eq "") {
            Invoke-RestMethod -Uri $Url -Method $Method -Headers $Headers
        } else {
            Invoke-RestMethod -Uri $Url -Method $Method -Headers $Headers -Body $Body -ContentType "application/json"
        }
        
        Write-Host "   ✅ Success" -ForegroundColor Green
        Write-Host "   Response: $($response | ConvertTo-Json -Depth 2)" -ForegroundColor White
        Write-Host ""
        return $response
    }
    catch {
        Write-Host "   ❌ Failed: $($_.Exception.Message)" -ForegroundColor Red
        Write-Host ""
        return $null
    }
}

# 1. Health Check
Write-Host "📊 1. HEALTH CHECK" -ForegroundColor Magenta
$healthResponse = Test-Endpoint -Method "GET" -Url "$baseUrl/ping" -Description "Health Check"

# 2. Authentication Tests
Write-Host "🔐 2. AUTHENTICATION TESTS" -ForegroundColor Magenta

# Register a test user
$registerBody = @{
    name = "Test User"
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

$registerResponse = Test-Endpoint -Method "POST" -Url "$apiUrl/auth/register" -Description "Register User" -Body $registerBody

# Login with the test user
$loginBody = @{
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

$loginResponse = Test-Endpoint -Method "POST" -Url "$apiUrl/auth/login" -Description "Login User" -Body $loginBody

if ($loginResponse -and $loginResponse.data.access_token) {
    $global:accessToken = $loginResponse.data.access_token
    Write-Host "🎫 Access token obtained!" -ForegroundColor Green
} else {
    Write-Host "❌ Failed to get access token. Some tests may fail." -ForegroundColor Red
}

# Headers with authentication
$authHeaders = @{
    "Authorization" = "Bearer $global:accessToken"
}

# 3. Public Course Tests (No Authentication)
Write-Host "📚 3. PUBLIC COURSE TESTS" -ForegroundColor Magenta

Test-Endpoint -Method "GET" -Url "$apiUrl/courses" -Description "Get Published Courses"
Test-Endpoint -Method "GET" -Url "$apiUrl/courses/search?category=Programming&level=beginner&page=1&limit=5" -Description "Search Courses"

# 4. Course Management Tests (Authentication Required)
Write-Host "🎓 4. COURSE MANAGEMENT TESTS" -ForegroundColor Magenta

# Create a test course
$courseBody = @{
    title = "Introduction to Go Programming"
    description = "Learn Go programming from scratch with hands-on examples"
    short_description = "Go programming basics for beginners"
    thumbnail = "https://example.com/go-course.jpg"
    level = "beginner"
    category = "Programming"
    tags = "go,programming,backend,api"
    price = 99.99
    is_published = $true
} | ConvertTo-Json

$createCourseResponse = Test-Endpoint -Method "POST" -Url "$apiUrl/courses" -Description "Create Course" -Body $courseBody -Headers $authHeaders

if ($createCourseResponse -and $createCourseResponse.data.id) {
    $global:courseId = $createCourseResponse.data.id
    Write-Host "📝 Course created with ID: $global:courseId" -ForegroundColor Green
}

# Get course details
if ($global:courseId -gt 0) {
    Test-Endpoint -Method "GET" -Url "$apiUrl/courses/$global:courseId" -Description "Get Course Details"
}

# Update course
$updateCourseBody = @{
    description = "Updated: Learn Go programming from scratch with hands-on examples and real-world projects"
    price = 79.99
} | ConvertTo-Json

if ($global:courseId -gt 0) {
    Test-Endpoint -Method "PUT" -Url "$apiUrl/courses/$global:courseId" -Description "Update Course" -Body $updateCourseBody -Headers $authHeaders
}

# Get my courses
Test-Endpoint -Method "GET" -Url "$apiUrl/my/courses" -Description "Get My Created Courses" -Headers $authHeaders

# 5. Course Enrollment Tests
Write-Host "📋 5. COURSE ENROLLMENT TESTS" -ForegroundColor Magenta

if ($global:courseId -gt 0) {
    Test-Endpoint -Method "POST" -Url "$apiUrl/courses/$global:courseId/enroll" -Description "Enroll in Course" -Headers $authHeaders
    Test-Endpoint -Method "GET" -Url "$apiUrl/my/enrolled-courses" -Description "Get My Enrolled Courses" -Headers $authHeaders
    Test-Endpoint -Method "GET" -Url "$apiUrl/courses/$global:courseId/progress" -Description "Get Course Progress" -Headers $authHeaders
}

# 6. Lesson Management Tests
Write-Host "📖 6. LESSON MANAGEMENT TESTS" -ForegroundColor Magenta

if ($global:courseId -gt 0) {
    # Create lessons
    $lesson1Body = @{
        title = "Introduction to Go Syntax"
        description = "Learn the basic syntax of Go programming language"
        video_url = "https://youtube.com/watch?v=example1"
        video_id = "example1"
        script = "Welcome to our Go programming course. In this lesson, we'll cover the basic syntax..."
        duration = 600
        sequence = 1
        is_published = $true
        is_free = $true
    } | ConvertTo-Json

    $createLessonResponse = Test-Endpoint -Method "POST" -Url "$apiUrl/courses/$global:courseId/lessons" -Description "Create Lesson 1" -Body $lesson1Body -Headers $authHeaders
    
    if ($createLessonResponse -and $createLessonResponse.data.id) {
        $global:lessonId = $createLessonResponse.data.id
        Write-Host "📚 Lesson created with ID: $global:lessonId" -ForegroundColor Green
    }

    # Create second lesson
    $lesson2Body = @{
        title = "Variables and Data Types"
        description = "Understanding Go variables and data types"
        video_url = "https://youtube.com/watch?v=example2"
        video_id = "example2"
        script = "In this lesson, we'll explore Go variables and data types..."
        duration = 720
        sequence = 2
        is_published = $true
        is_free = $false
    } | ConvertTo-Json

    Test-Endpoint -Method "POST" -Url "$apiUrl/courses/$global:courseId/lessons" -Description "Create Lesson 2" -Body $lesson2Body -Headers $authHeaders

    # Get course lessons
    Test-Endpoint -Method "GET" -Url "$apiUrl/courses/$global:courseId/lessons" -Description "Get Course Lessons" -Headers $authHeaders
    Test-Endpoint -Method "GET" -Url "$apiUrl/courses/$global:courseId/lessons/free" -Description "Get Free Course Lessons"
}

# Get lesson details
if ($global:lessonId -gt 0) {
    Test-Endpoint -Method "GET" -Url "$apiUrl/lessons/$global:lessonId" -Description "Get Lesson Details" -Headers $authHeaders
}

# 7. Progress Tracking Tests
Write-Host "📈 7. PROGRESS TRACKING TESTS" -ForegroundColor Magenta

if ($global:lessonId -gt 0) {
    # Update lesson progress
    $progressBody = @{
        lesson_id = $global:lessonId
        watch_time = 300
        is_completed = $false
    } | ConvertTo-Json

    Test-Endpoint -Method "POST" -Url "$apiUrl/lessons/progress" -Description "Update Lesson Progress" -Body $progressBody -Headers $authHeaders

    # Mark lesson as completed
    $completionBody = @{
        watch_time = 580
    } | ConvertTo-Json

    Test-Endpoint -Method "POST" -Url "$apiUrl/lessons/$global:lessonId/complete" -Description "Mark Lesson Completed" -Body $completionBody -Headers $authHeaders

    # Get lesson progress
    Test-Endpoint -Method "GET" -Url "$apiUrl/courses/$global:courseId/lessons/progress" -Description "Get User Lesson Progress" -Headers $authHeaders
}

# 8. Analytics Tests
Write-Host "📊 8. ANALYTICS TESTS" -ForegroundColor Magenta

if ($global:courseId -gt 0) {
    Test-Endpoint -Method "GET" -Url "$apiUrl/courses/$global:courseId/analytics" -Description "Get Course Analytics" -Headers $authHeaders
}

# 9. Admin Tests
Write-Host "👨‍💼 9. ADMIN TESTS" -ForegroundColor Magenta

Test-Endpoint -Method "GET" -Url "$apiUrl/admin/courses" -Description "Get All Courses (Admin)" -Headers $authHeaders

# 10. Lesson Reordering Test
Write-Host "🔄 10. LESSON REORDERING TEST" -ForegroundColor Magenta

if ($global:courseId -gt 0) {
    $reorderBody = @(
        @{
            lesson_id = $global:lessonId
            sequence = 2
        }
    ) | ConvertTo-Json

    Test-Endpoint -Method "PUT" -Url "$apiUrl/courses/$global:courseId/lessons/reorder" -Description "Reorder Lessons" -Body $reorderBody -Headers $authHeaders
}

# 11. Unenrollment Test
Write-Host "🚪 11. UNENROLLMENT TEST" -ForegroundColor Magenta

if ($global:courseId -gt 0) {
    Test-Endpoint -Method "DELETE" -Url "$apiUrl/courses/$global:courseId/enroll" -Description "Unenroll from Course" -Headers $authHeaders
}

Write-Host ""
Write-Host "🎉 API Testing Complete!" -ForegroundColor Green
Write-Host "📝 Summary:" -ForegroundColor Yellow
Write-Host "   - Health Check: Tested" -ForegroundColor White
Write-Host "   - Authentication: Register & Login" -ForegroundColor White
Write-Host "   - Course Management: Create, Read, Update" -ForegroundColor White
Write-Host "   - Course Enrollment: Enroll, Progress, Unenroll" -ForegroundColor White
Write-Host "   - Lesson Management: Create, Read, Reorder" -ForegroundColor White
Write-Host "   - Progress Tracking: Update & Complete" -ForegroundColor White
Write-Host "   - Analytics: Course Statistics" -ForegroundColor White
Write-Host "   - Admin Functions: All Courses Access" -ForegroundColor White
Write-Host ""
Write-Host "🔧 Course ID used: $global:courseId" -ForegroundColor Cyan
Write-Host "📚 Lesson ID used: $global:lessonId" -ForegroundColor Cyan
Write-Host ""
