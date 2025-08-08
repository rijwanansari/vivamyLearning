package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/dto"
	"github.com/rijwanansari/vivaLearning/services"
)

type CourseController struct {
	CourseService services.CourseService
	LessonService services.LessonService
	Validator     *validator.Validate
}

func NewCourseController(courseService services.CourseService, lessonService services.LessonService) *CourseController {
	return &CourseController{
		CourseService: courseService,
		LessonService: lessonService,
		Validator:     validator.New(),
	}
}

// Course CRUD operations

// CreateCourse creates a new course
// POST /api/courses
func (cc *CourseController) CreateCourse(c echo.Context) error {
	var req dto.CreateCourseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	if err := cc.Validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	// Get user ID from context (should be set by auth middleware)
	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	course, err := cc.CourseService.CreateCourse(req, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Course created successfully",
		Data:    course,
	})
}

// GetCourse gets a course by ID
// GET /api/courses/:id
func (cc *CourseController) GetCourse(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	var userID *uint
	if uid := getUserIDFromContext(c); uid != 0 {
		userID = &uid
	}

	course, err := cc.CourseService.GetCourseByID(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.APIResponse{
			Success: false,
			Error:   "Course not found",
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    course,
	})
}

// UpdateCourse updates a course
// PUT /api/courses/:id
func (cc *CourseController) UpdateCourse(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	var req dto.UpdateCourseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	if err := cc.Validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	course, err := cc.CourseService.UpdateCourse(uint(id), req, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Course updated successfully",
		Data:    course,
	})
}

// DeleteCourse deletes a course
// DELETE /api/courses/:id
func (cc *CourseController) DeleteCourse(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	err = cc.CourseService.DeleteCourse(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Course deleted successfully",
	})
}

// GetAllCourses gets all courses (admin)
// GET /api/admin/courses
func (cc *CourseController) GetAllCourses(c echo.Context) error {
	courses, err := cc.CourseService.GetAllCourses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    courses,
	})
}

// GetPublishedCourses gets all published courses
// GET /api/courses
func (cc *CourseController) GetPublishedCourses(c echo.Context) error {
	var userID *uint
	if uid := getUserIDFromContext(c); uid != 0 {
		userID = &uid
	}

	courses, err := cc.CourseService.GetPublishedCourses(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    courses,
	})
}

// SearchCourses searches courses with filters
// GET /api/courses/search
func (cc *CourseController) SearchCourses(c echo.Context) error {
	var filter dto.CourseFilterRequest
	if err := c.Bind(&filter); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid query parameters",
		})
	}

	if err := cc.Validator.Struct(filter); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	var userID *uint
	if uid := getUserIDFromContext(c); uid != 0 {
		userID = &uid
	}

	result, err := cc.CourseService.SearchCourses(filter, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    result,
	})
}

// GetMyCourses gets courses created by the authenticated user
// GET /api/my/courses
func (cc *CourseController) GetMyCourses(c echo.Context) error {
	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	courses, err := cc.CourseService.GetCoursesByCreator(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    courses,
	})
}

// EnrollInCourse enrolls user in a course
// POST /api/courses/:id/enroll
func (cc *CourseController) EnrollInCourse(c echo.Context) error {
	idParam := c.Param("id")
	courseID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	result, err := cc.CourseService.EnrollInCourse(uint(courseID), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, *result)
	}

	return c.JSON(http.StatusOK, *result)
}

// UnenrollFromCourse unenrolls user from a course
// DELETE /api/courses/:id/enroll
func (cc *CourseController) UnenrollFromCourse(c echo.Context) error {
	idParam := c.Param("id")
	courseID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	result, err := cc.CourseService.UnenrollFromCourse(uint(courseID), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, *result)
	}

	return c.JSON(http.StatusOK, *result)
}

// GetMyEnrolledCourses gets courses user is enrolled in
// GET /api/my/enrolled-courses
func (cc *CourseController) GetMyEnrolledCourses(c echo.Context) error {
	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	courses, err := cc.CourseService.GetUserEnrolledCourses(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    courses,
	})
}

// GetCourseProgress gets user's progress for a specific course
// GET /api/courses/:id/progress
func (cc *CourseController) GetCourseProgress(c echo.Context) error {
	idParam := c.Param("id")
	courseID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	progress, err := cc.CourseService.GetUserCourseProgress(uint(courseID), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.APIResponse{
			Success: false,
			Error:   "Course progress not found",
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    progress,
	})
}

// GetCourseAnalytics gets analytics for a course
// GET /api/courses/:id/analytics
func (cc *CourseController) GetCourseAnalytics(c echo.Context) error {
	idParam := c.Param("id")
	courseID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	analytics, err := cc.CourseService.GetCourseAnalytics(uint(courseID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    analytics,
	})
}

// Helper function to get user ID from context
func getUserIDFromContext(c echo.Context) uint {
	// This should be set by your authentication middleware
	if userID := c.Get("user_id"); userID != nil {
		if id, ok := userID.(uint); ok {
			return id
		}
		if id, ok := userID.(float64); ok {
			return uint(id)
		}
		if idStr, ok := userID.(string); ok {
			if id, err := strconv.ParseUint(idStr, 10, 32); err == nil {
				return uint(id)
			}
		}
	}
	return 0
}
