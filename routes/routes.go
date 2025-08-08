package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/controllers"
	"github.com/rijwanansari/vivaLearning/middlewares"
)

type Routes struct {
	echo   *echo.Echo
	auth   *controllers.AuthController
	course *controllers.CourseController
	lesson *controllers.LessonController
}

func New(e *echo.Echo, auth *controllers.AuthController, course *controllers.CourseController, lesson *controllers.LessonController) *Routes {
	return &Routes{
		echo:   e,
		auth:   auth,
		course: course,
		lesson: lesson,
	}
}

func (r *Routes) Init() {
	e := r.echo

	// Health check
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "VivaLearning API is running!")
	})

	// API v1 routes
	api := e.Group("/api/v1")

	// Authentication routes (public)
	auth := api.Group("/auth")
	auth.POST("/register", r.auth.RegisterUser)
	auth.POST("/login", r.auth.LoginUser)
	// auth.POST("/refresh", r.auth.RefreshToken)
	// auth.POST("/logout", r.auth.LogoutUser)

	// Public course routes (no authentication required)
	publicCourses := api.Group("/courses")
	publicCourses.GET("", r.course.GetPublishedCourses)                         // GET /api/v1/courses
	publicCourses.GET("/search", r.course.SearchCourses)                        // GET /api/v1/courses/search
	publicCourses.GET("/:id", r.course.GetCourse)                               // GET /api/v1/courses/:id
	publicCourses.GET("/:courseId/lessons/free", r.lesson.GetFreeCourseLessons) // GET /api/v1/courses/:courseId/lessons/free

	// Protected routes (require authentication)
	protected := api.Group("")
	protected.Use(middlewares.JWTMiddleware)

	// User profile routes (TODO: implement user controller)
	// profile := protected.Group("/profile")
	// profile.GET("", r.user.GetProfile)
	// profile.PUT("", r.user.UpdateProfile)

	// User's enrolled courses
	myCourses := protected.Group("/my")
	myCourses.GET("/courses", r.course.GetMyCourses)                  // GET /api/v1/my/courses (created courses)
	myCourses.GET("/enrolled-courses", r.course.GetMyEnrolledCourses) // GET /api/v1/my/enrolled-courses

	// Course management (for creators)
	courseAdmin := protected.Group("/courses")
	courseAdmin.POST("", r.course.CreateCourse)                    // POST /api/v1/courses
	courseAdmin.PUT("/:id", r.course.UpdateCourse)                 // PUT /api/v1/courses/:id
	courseAdmin.DELETE("/:id", r.course.DeleteCourse)              // DELETE /api/v1/courses/:id
	courseAdmin.GET("/:id/analytics", r.course.GetCourseAnalytics) // GET /api/v1/courses/:id/analytics

	// Course enrollment
	enrollment := protected.Group("/courses")
	enrollment.POST("/:id/enroll", r.course.EnrollInCourse)       // POST /api/v1/courses/:id/enroll
	enrollment.DELETE("/:id/enroll", r.course.UnenrollFromCourse) // DELETE /api/v1/courses/:id/enroll
	enrollment.GET("/:id/progress", r.course.GetCourseProgress)   // GET /api/v1/courses/:id/progress

	// Lesson management (for creators)
	lessonAdmin := protected.Group("/courses/:courseId/lessons")
	lessonAdmin.POST("", r.lesson.CreateLesson)          // POST /api/v1/courses/:courseId/lessons
	lessonAdmin.PUT("/reorder", r.lesson.ReorderLessons) // PUT /api/v1/courses/:courseId/lessons/reorder

	// Lesson access (for enrolled users)
	lessons := protected.Group("")
	lessons.GET("/courses/:courseId/lessons", r.lesson.GetCourseLessons)               // GET /api/v1/courses/:courseId/lessons
	lessons.GET("/courses/:courseId/lessons/progress", r.lesson.GetUserLessonProgress) // GET /api/v1/courses/:courseId/lessons/progress
	lessons.GET("/lessons/:id", r.lesson.GetLesson)                                    // GET /api/v1/lessons/:id
	lessons.PUT("/lessons/:id", r.lesson.UpdateLesson)                                 // PUT /api/v1/lessons/:id
	lessons.DELETE("/lessons/:id", r.lesson.DeleteLesson)                              // DELETE /api/v1/lessons/:id

	// Lesson progress tracking
	progress := protected.Group("/lessons")
	progress.POST("/progress", r.lesson.UpdateLessonProgress)    // POST /api/v1/lessons/progress
	progress.POST("/:id/complete", r.lesson.MarkLessonCompleted) // POST /api/v1/lessons/:id/complete

	// Admin routes (require admin role - can be added later)
	admin := protected.Group("/admin")
	admin.GET("/courses", r.course.GetAllCourses) // GET /api/v1/admin/courses
	// admin.GET("/users", r.user.GetAllUsers)
	// admin.GET("/analytics", r.admin.GetPlatformAnalytics)
}
