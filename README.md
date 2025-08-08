# VivaLearning - Learning Management System API

A comprehensive learning management system built with Go, Echo framework, GORM, and PostgreSQL. This API provides full course and lesson management capabilities with user enrollment, progress tracking, and analytics.

## 🚀 Features

- **User Authentication & Authorization** with JWT tokens
- **Course Management** - Create, update, delete, and publish courses
- **Lesson Management** - Organize lessons with sequencing and progress tracking
- **User Enrollment** - Course enrollment and unenrollment
- **Progress Tracking** - Track user progress through courses and lessons
- **Search & Filtering** - Advanced course search with multiple filters
- **Analytics** - Course completion rates and user progress analytics
- **Content Access Control** - Free preview lessons and enrollment-based access

## 🛠️ Tech Stack

- **Backend:** Go 1.24+
- **Web Framework:** Echo v4
- **Database:** PostgreSQL with GORM ORM
- **Authentication:** JWT (JSON Web Tokens)
- **Configuration:** Environment variables with Viper
- **Validation:** go-playground/validator

## 📋 Prerequisites

Before running the application, make sure you have:

- Go 1.24 or higher installed
- PostgreSQL database running
- Git for version control

## ⚙️ Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/rijwanansari/vivamyLearning.git
cd vivamyLearning
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Environment Configuration

Create a `.env` file in the root directory:

```env
# Application Configuration
APP_NAME=vivaLearning
APP_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=admin
DB_SCHEMA=Mysample

# Database Connection Pool Settings
DB_MAX_IDLE_CONN=10
DB_MAX_OPEN_CONN=100
DB_MAX_CONN_LIFETIME=3600

# Database Debug Mode
DB_DEBUG=true

# Logger Configuration
LOG_LEVEL=info
LOG_FILE_PATH=logs/app.log

# JWT Configuration
JWT_ACCESS_TOKEN_SECRET=your-super-secret-access-token-key-change-in-production
JWT_REFRESH_TOKEN_SECRET=your-super-secret-refresh-token-key-change-in-production
JWT_ACCESS_TOKEN_EXPIRY=900
JWT_REFRESH_TOKEN_EXPIRY=604800
```

### 4. Database Setup

Make sure PostgreSQL is running and create a database:

```sql
CREATE DATABASE Mysample;
```

The application will automatically create the required tables using GORM auto-migration.

### 5. Build and Run

```bash
# Build the application
go build

# Run the application
./vivaLearning serve

# Or run directly with go
go run main.go serve
```

The API will be available at `http://localhost:8080`

## 📚 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Health Check
- `GET /ping` - API health check

### 🔐 Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/auth/register` | Register new user | No |
| POST | `/auth/login` | User login | No |

### 🎓 Course Management Endpoints

#### Public Endpoints (No Authentication)
| Method | Endpoint | Description | Parameters |
|--------|----------|-------------|------------|
| GET | `/courses` | Get all published courses | - |
| GET | `/courses/search` | Search courses with filters | `category`, `level`, `min_price`, `max_price`, `tags`, `search`, `page`, `limit`, `sort_by`, `sort_order` |
| GET | `/courses/{id}` | Get course details | - |
| GET | `/courses/{courseId}/lessons/free` | Get free preview lessons | - |

#### Protected Endpoints (Authentication Required)

**Course Creation & Management:**
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/courses` | Create new course | Yes |
| PUT | `/courses/{id}` | Update course | Yes (Creator only) |
| DELETE | `/courses/{id}` | Delete course | Yes (Creator only) |
| GET | `/courses/{id}/analytics` | Get course analytics | Yes (Creator only) |
| GET | `/my/courses` | Get courses created by user | Yes |

**Course Enrollment:**
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/courses/{id}/enroll` | Enroll in course | Yes |
| DELETE | `/courses/{id}/enroll` | Unenroll from course | Yes |
| GET | `/my/enrolled-courses` | Get enrolled courses | Yes |
| GET | `/courses/{id}/progress` | Get course progress | Yes |

### 📚 Lesson Management Endpoints

**Lesson Creation & Management:**
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/courses/{courseId}/lessons` | Create lesson | Yes (Creator only) |
| PUT | `/lessons/{id}` | Update lesson | Yes (Creator only) |
| DELETE | `/lessons/{id}` | Delete lesson | Yes (Creator only) |
| PUT | `/courses/{courseId}/lessons/reorder` | Reorder lessons | Yes (Creator only) |

**Lesson Access:**
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/courses/{courseId}/lessons` | Get course lessons | Yes (Enrolled users) |
| GET | `/lessons/{id}` | Get lesson details | Yes |
| GET | `/courses/{courseId}/lessons/progress` | Get lesson progress | Yes |

**Progress Tracking:**
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/lessons/progress` | Update lesson progress | Yes |
| POST | `/lessons/{id}/complete` | Mark lesson as completed | Yes |

### 👨‍💼 Admin Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/admin/courses` | Get all courses (including unpublished) | Yes (Admin) |

## 📝 Request/Response Examples

### User Registration
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### User Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create Course
```bash
curl -X POST http://localhost:8080/api/v1/courses \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Introduction to Go Programming",
    "description": "Learn Go programming from scratch",
    "short_description": "Go programming basics",
    "level": "beginner",
    "category": "Programming",
    "tags": "go,programming,backend",
    "price": 99.99,
    "is_published": true
  }'
```

### Search Courses
```bash
curl "http://localhost:8080/api/v1/courses/search?category=Programming&level=beginner&page=1&limit=10"
```

### Enroll in Course
```bash
curl -X POST http://localhost:8080/api/v1/courses/1/enroll \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Create Lesson
```bash
curl -X POST http://localhost:8080/api/v1/courses/1/lessons \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Introduction to Variables",
    "description": "Learn about Go variables",
    "video_url": "https://youtube.com/watch?v=example",
    "duration": 600,
    "sequence": 1,
    "is_published": true,
    "is_free": true
  }'
```

### Mark Lesson as Completed
```bash
curl -X POST http://localhost:8080/api/v1/lessons/1/complete \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "watch_time": 580
  }'
```

## 🗃️ Database Schema

### Core Models

**User**
- ID, Name, Email, Password (hashed)
- CreatedAt, UpdatedAt

**Course**
- ID, Title, Description, ShortDescription
- Thumbnail, Level, Category, Tags
- Duration, Price, IsPublished
- CreatedBy, CreatedAt, UpdatedAt

**Lesson**
- ID, Title, Description
- VideoURL, VideoID, Script
- Duration, CourseID, Sequence
- IsPublished, IsFree
- CreatedAt, UpdatedAt

**UserCourse** (Enrollment tracking)
- ID, UserID, CourseID
- LastLessonID, Progress, IsCompleted
- EnrolledAt, CompletedAt, UpdatedAt

**UserLesson** (Progress tracking)
- ID, UserID, LessonID, CourseID
- IsCompleted, WatchTime
- CompletedAt, CreatedAt, UpdatedAt

## 🔧 Configuration

The application uses environment variables for configuration. Key settings include:

- **Database:** Connection details and pool settings
- **JWT:** Token secrets and expiry times
- **Server:** Port and application name
- **Logging:** Level and file path

## 🏗️ Project Structure

```
├── cmd/                    # Command line interface
│   ├── root.go            # Root command configuration
│   └── serve.go           # Server start command
├── config/                # Configuration management
│   └── config.go          # Environment configuration
├── conn/                  # Database connection
│   └── db.go             # Database setup and migration
├── controllers/           # HTTP request handlers
│   ├── auth_controller.go
│   ├── course_controller.go
│   └── lesson_controller.go
├── domain/               # Domain models
│   ├── user.go
│   ├── course.go
│   ├── lesson.go
│   ├── usercourse.go
│   └── user_lesson.go
├── dto/                  # Data transfer objects
│   ├── course_dto.go
│   └── lesson_dto.go
├── middlewares/          # HTTP middlewares
├── repositories/         # Data access layer
│   ├── user_repository.go
│   ├── course_repository.go
│   ├── lesson_repository.go
│   └── user_course_repository.go
├── routes/               # Route definitions
│   └── routes.go
├── services/             # Business logic layer
│   ├── auth_service.go
│   ├── user_service.go
│   ├── course_service.go
│   ├── lesson_service.go
│   └── token_service.go
├── types/                # Type definitions
├── utils/                # Utility functions
├── .env                  # Environment variables
├── go.mod               # Go modules
├── go.sum               # Go dependencies
└── main.go              # Application entry point
```

## 🧪 Testing

You can test the API endpoints using tools like:

- **Postman** - Import the API collection
- **curl** - Command line testing (examples above)
- **httpie** - User-friendly HTTP client
- **Insomnia** - API testing tool

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Check PostgreSQL is running
   - Verify database credentials in `.env`
   - Ensure database exists

2. **JWT Token Invalid**
   - Check token format in Authorization header
   - Verify JWT secrets in configuration
   - Ensure token hasn't expired

3. **Permission Denied**
   - Check user authentication
   - Verify user owns the resource (for creator-only endpoints)

### Debug Mode

Enable debug mode by setting `DB_DEBUG=true` in your `.env` file to see SQL queries.

## 🔒 Security Considerations

- Change JWT secrets in production
- Use HTTPS in production
- Implement rate limiting
- Add input sanitization
- Use secure password hashing (already implemented)

## 🚀 Deployment

### Production Setup

1. Set production environment variables
2. Use a production-grade PostgreSQL instance
3. Configure reverse proxy (nginx)
4. Set up SSL certificates
5. Implement monitoring and logging

### Docker Deployment (Optional)

```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o vivalearning

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/vivalearning .
CMD ["./vivalearning", "serve"]
```

## 📄 License

This project is licensed under the MIT License.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📞 Support

For questions or issues, please create an issue in the GitHub repository.

---

**Happy Learning! 🎓**
