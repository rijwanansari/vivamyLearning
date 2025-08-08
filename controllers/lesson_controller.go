package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/dto"
	"github.com/rijwanansari/vivaLearning/services"
)

type LessonController struct {
	LessonService services.LessonService
	Validator     *validator.Validate
}

func NewLessonController(lessonService services.LessonService) *LessonController {
	return &LessonController{
		LessonService: lessonService,
		Validator:     validator.New(),
	}
}

// CreateLesson creates a new lesson for a course
// POST /api/courses/:courseId/lessons
func (lc *LessonController) CreateLesson(c echo.Context) error {
	courseIDParam := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	var req dto.CreateLessonRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	if err := lc.Validator.Struct(req); err != nil {
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

	lesson, err := lc.LessonService.CreateLesson(uint(courseID), req, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Lesson created successfully",
		Data:    lesson,
	})
}

// GetLesson gets a lesson by ID
// GET /api/lessons/:id
func (lc *LessonController) GetLesson(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid lesson ID",
		})
	}

	var userID *uint
	if uid := getUserIDFromContext(c); uid != 0 {
		userID = &uid
	}

	lesson, err := lc.LessonService.GetLessonByID(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.APIResponse{
			Success: false,
			Error:   "Lesson not found",
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    lesson,
	})
}

// UpdateLesson updates a lesson
// PUT /api/lessons/:id
func (lc *LessonController) UpdateLesson(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid lesson ID",
		})
	}

	var req dto.UpdateLessonRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	if err := lc.Validator.Struct(req); err != nil {
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

	lesson, err := lc.LessonService.UpdateLesson(uint(id), req, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Lesson updated successfully",
		Data:    lesson,
	})
}

// DeleteLesson deletes a lesson
// DELETE /api/lessons/:id
func (lc *LessonController) DeleteLesson(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid lesson ID",
		})
	}

	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	err = lc.LessonService.DeleteLesson(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Lesson deleted successfully",
	})
}

// GetCourseLessons gets all lessons for a course
// GET /api/courses/:courseId/lessons
func (lc *LessonController) GetCourseLessons(c echo.Context) error {
	courseIDParam := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDParam, 10, 32)
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

	lessons, err := lc.LessonService.GetLessonsByCourse(uint(courseID), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    lessons,
	})
}

// GetFreeCourseLessons gets free lessons for a course (preview)
// GET /api/courses/:courseId/lessons/free
func (lc *LessonController) GetFreeCourseLessons(c echo.Context) error {
	courseIDParam := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	lessons, err := lc.LessonService.GetFreeLessonsByCourse(uint(courseID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    lessons,
	})
}

// ReorderLessons reorders lessons in a course
// PUT /api/courses/:courseId/lessons/reorder
func (lc *LessonController) ReorderLessons(c echo.Context) error {
	courseIDParam := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid course ID",
		})
	}

	var req []struct {
		LessonID uint `json:"lesson_id" validate:"required"`
		Sequence int  `json:"sequence" validate:"required,min=1"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	for _, item := range req {
		if err := lc.Validator.Struct(item); err != nil {
			return c.JSON(http.StatusBadRequest, dto.APIResponse{
				Success: false,
				Error:   err.Error(),
			})
		}
	}

	userID := getUserIDFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
	}

	// Convert to the expected type
	var lessonSequences []struct {
		LessonID uint `json:"lesson_id"`
		Sequence int  `json:"sequence"`
	}

	for _, item := range req {
		lessonSequences = append(lessonSequences, struct {
			LessonID uint `json:"lesson_id"`
			Sequence int  `json:"sequence"`
		}{
			LessonID: item.LessonID,
			Sequence: item.Sequence,
		})
	}

	err = lc.LessonService.ReorderLessons(uint(courseID), lessonSequences, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Lessons reordered successfully",
	})
}

// UpdateLessonProgress updates user's progress for a lesson
// POST /api/lessons/progress
func (lc *LessonController) UpdateLessonProgress(c echo.Context) error {
	var req dto.UpdateProgressRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	if err := lc.Validator.Struct(req); err != nil {
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

	result, err := lc.LessonService.UpdateLessonProgress(userID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, *result)
	}

	return c.JSON(http.StatusOK, *result)
}

// MarkLessonCompleted marks a lesson as completed
// POST /api/lessons/:id/complete
func (lc *LessonController) MarkLessonCompleted(c echo.Context) error {
	idParam := c.Param("id")
	lessonID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid lesson ID",
		})
	}

	var req struct {
		WatchTime int `json:"watch_time" validate:"min=0"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	if err := lc.Validator.Struct(req); err != nil {
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

	result, err := lc.LessonService.MarkLessonCompleted(userID, uint(lessonID), req.WatchTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, *result)
	}

	return c.JSON(http.StatusOK, *result)
}

// GetUserLessonProgress gets user's progress for all lessons in a course
// GET /api/courses/:courseId/lessons/progress
func (lc *LessonController) GetUserLessonProgress(c echo.Context) error {
	courseIDParam := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDParam, 10, 32)
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

	lessons, err := lc.LessonService.GetUserLessonProgress(userID, uint(courseID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    lessons,
	})
}
