package services

import (
	"errors"
	"time"

	"github.com/rijwanansari/vivaLearning/domain"
	"github.com/rijwanansari/vivaLearning/dto"
	repository "github.com/rijwanansari/vivaLearning/repositories"
)

type LessonService interface {
	// Admin operations
	CreateLesson(courseID uint, req dto.CreateLessonRequest, userID uint) (*dto.LessonResponse, error)
	UpdateLesson(id uint, req dto.UpdateLessonRequest, userID uint) (*dto.LessonResponse, error)
	DeleteLesson(id uint, userID uint) error
	ReorderLessons(courseID uint, lessonSequences []struct {
		LessonID uint `json:"lesson_id"`
		Sequence int  `json:"sequence"`
	}, userID uint) error

	// Public operations
	GetLessonByID(id uint, userID *uint) (*dto.LessonResponse, error)
	GetLessonsByCourse(courseID uint, userID *uint) ([]dto.LessonResponse, error)
	GetFreeLessonsByCourse(courseID uint) ([]dto.LessonResponse, error)

	// User progress operations
	UpdateLessonProgress(userID uint, req dto.UpdateProgressRequest) (*dto.APIResponse, error)
	GetUserLessonProgress(userID, courseID uint) ([]dto.LessonResponse, error)
	MarkLessonCompleted(userID, lessonID uint, watchTime int) (*dto.APIResponse, error)
}

type LessonServiceImp struct {
	LessonRepo     repository.LessonRepository
	CourseRepo     repository.CourseRepository
	UserCourseRepo repository.UserCourseRepository
}

func NewLessonService(lessonRepo repository.LessonRepository, courseRepo repository.CourseRepository, userCourseRepo repository.UserCourseRepository) LessonService {
	return &LessonServiceImp{
		LessonRepo:     lessonRepo,
		CourseRepo:     courseRepo,
		UserCourseRepo: userCourseRepo,
	}
}

func (s *LessonServiceImp) CreateLesson(courseID uint, req dto.CreateLessonRequest, userID uint) (*dto.LessonResponse, error) {
	// Check if course exists and user is the creator
	course, err := s.CourseRepo.GetByID(courseID)
	if err != nil {
		return nil, err
	}

	if course.CreatedBy != userID {
		return nil, errors.New("unauthorized to create lesson for this course")
	}

	// Get next sequence number if not provided
	sequence := req.Sequence
	if sequence <= 0 {
		sequence, err = s.LessonRepo.GetNextSequence(courseID)
		if err != nil {
			return nil, err
		}
	}

	lesson := &domain.Lesson{
		Title:       req.Title,
		Description: req.Description,
		VideoURL:    req.VideoURL,
		VideoID:     req.VideoID,
		Script:      req.Script,
		Duration:    req.Duration,
		CourseID:    courseID,
		Sequence:    sequence,
		IsPublished: req.IsPublished,
		IsFree:      req.IsFree,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = s.LessonRepo.Create(lesson)
	if err != nil {
		return nil, err
	}

	return s.mapLessonToResponse(lesson, false), nil
}

func (s *LessonServiceImp) UpdateLesson(id uint, req dto.UpdateLessonRequest, userID uint) (*dto.LessonResponse, error) {
	lesson, err := s.LessonRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if user is the course creator
	course, err := s.CourseRepo.GetByID(lesson.CourseID)
	if err != nil {
		return nil, err
	}

	if course.CreatedBy != userID {
		return nil, errors.New("unauthorized to update this lesson")
	}

	// Update fields if provided
	if req.Title != nil {
		lesson.Title = *req.Title
	}
	if req.Description != nil {
		lesson.Description = *req.Description
	}
	if req.VideoURL != nil {
		lesson.VideoURL = *req.VideoURL
	}
	if req.VideoID != nil {
		lesson.VideoID = *req.VideoID
	}
	if req.Script != nil {
		lesson.Script = *req.Script
	}
	if req.Duration != nil {
		lesson.Duration = *req.Duration
	}
	if req.Sequence != nil {
		lesson.Sequence = *req.Sequence
	}
	if req.IsPublished != nil {
		lesson.IsPublished = *req.IsPublished
	}
	if req.IsFree != nil {
		lesson.IsFree = *req.IsFree
	}

	lesson.UpdatedAt = time.Now()

	err = s.LessonRepo.Update(lesson)
	if err != nil {
		return nil, err
	}

	return s.mapLessonToResponse(lesson, false), nil
}

func (s *LessonServiceImp) DeleteLesson(id uint, userID uint) error {
	lesson, err := s.LessonRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Check if user is the course creator
	course, err := s.CourseRepo.GetByID(lesson.CourseID)
	if err != nil {
		return err
	}

	if course.CreatedBy != userID {
		return errors.New("unauthorized to delete this lesson")
	}

	return s.LessonRepo.Delete(id)
}

func (s *LessonServiceImp) ReorderLessons(courseID uint, lessonSequences []struct {
	LessonID uint `json:"lesson_id"`
	Sequence int  `json:"sequence"`
}, userID uint) error {
	// Check if user is the course creator
	course, err := s.CourseRepo.GetByID(courseID)
	if err != nil {
		return err
	}

	if course.CreatedBy != userID {
		return errors.New("unauthorized to reorder lessons for this course")
	}

	return s.LessonRepo.ReorderLessons(courseID, lessonSequences)
}

func (s *LessonServiceImp) GetLessonByID(id uint, userID *uint) (*dto.LessonResponse, error) {
	lesson, err := s.LessonRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if user has access to the lesson
	hasAccess := false
	isCompleted := false

	if lesson.IsFree || lesson.IsPublished {
		hasAccess = true
	}

	if userID != nil {
		// Check if user is enrolled in the course
		if enrolled, _ := s.UserCourseRepo.IsUserEnrolled(*userID, lesson.CourseID); enrolled {
			hasAccess = true

			// Check if lesson is completed by user
			userLessons, err := s.LessonRepo.GetUserLessonProgress(*userID, lesson.CourseID)
			if err == nil {
				for _, ul := range userLessons {
					if ul.LessonID == lesson.ID && ul.IsCompleted {
						isCompleted = true
						break
					}
				}
			}
		}
	}

	response := s.mapLessonToResponse(lesson, isCompleted)

	// Hide script if user doesn't have access
	if !hasAccess {
		response.Script = ""
	}

	return response, nil
}

func (s *LessonServiceImp) GetLessonsByCourse(courseID uint, userID *uint) ([]dto.LessonResponse, error) {
	var lessons []domain.Lesson
	var err error

	// Check if user is enrolled
	isEnrolled := false
	if userID != nil {
		isEnrolled, _ = s.UserCourseRepo.IsUserEnrolled(*userID, courseID)
	}

	if isEnrolled {
		// Get all published lessons for enrolled users
		lessons, err = s.LessonRepo.GetPublishedLessonsByCourse(courseID)
	} else {
		// Get only free lessons for non-enrolled users
		lessons, err = s.LessonRepo.GetFreeLessonsByCourse(courseID)
	}

	if err != nil {
		return nil, err
	}

	// Get user progress if enrolled
	var userLessons []domain.UserLesson
	if isEnrolled {
		userLessons, _ = s.LessonRepo.GetUserLessonProgress(*userID, courseID)
	}

	var responses []dto.LessonResponse
	for _, lesson := range lessons {
		isCompleted := false

		// Check if lesson is completed by user
		for _, ul := range userLessons {
			if ul.LessonID == lesson.ID && ul.IsCompleted {
				isCompleted = true
				break
			}
		}

		response := s.mapLessonToResponse(&lesson, isCompleted)

		// Hide script for non-enrolled users unless it's a free lesson
		if !isEnrolled && !lesson.IsFree {
			response.Script = ""
		}

		responses = append(responses, *response)
	}

	return responses, nil
}

func (s *LessonServiceImp) GetFreeLessonsByCourse(courseID uint) ([]dto.LessonResponse, error) {
	lessons, err := s.LessonRepo.GetFreeLessonsByCourse(courseID)
	if err != nil {
		return nil, err
	}

	var responses []dto.LessonResponse
	for _, lesson := range lessons {
		responses = append(responses, *s.mapLessonToResponse(&lesson, false))
	}

	return responses, nil
}

func (s *LessonServiceImp) UpdateLessonProgress(userID uint, req dto.UpdateProgressRequest) (*dto.APIResponse, error) {
	// Get lesson details
	lesson, err := s.LessonRepo.GetByID(req.LessonID)
	if err != nil {
		return &dto.APIResponse{
			Success: false,
			Error:   "Lesson not found",
		}, err
	}

	// Check if user is enrolled in the course
	enrolled, err := s.UserCourseRepo.IsUserEnrolled(userID, lesson.CourseID)
	if err != nil || !enrolled {
		return &dto.APIResponse{
			Success: false,
			Error:   "User not enrolled in this course",
		}, errors.New("user not enrolled")
	}

	// Update progress
	if req.IsCompleted {
		err = s.LessonRepo.MarkLessonCompleted(userID, req.LessonID, lesson.CourseID, req.WatchTime)
	} else {
		// Create or update progress record
		userLesson := &domain.UserLesson{
			UserID:      userID,
			LessonID:    req.LessonID,
			CourseID:    lesson.CourseID,
			IsCompleted: false,
			WatchTime:   req.WatchTime,
		}
		err = s.LessonRepo.UpdateUserLessonProgress(userLesson)
	}

	if err != nil {
		return &dto.APIResponse{
			Success: false,
			Error:   "Failed to update lesson progress",
		}, err
	}

	return &dto.APIResponse{
		Success: true,
		Message: "Lesson progress updated successfully",
	}, nil
}

func (s *LessonServiceImp) GetUserLessonProgress(userID, courseID uint) ([]dto.LessonResponse, error) {
	// Check if user is enrolled
	enrolled, err := s.UserCourseRepo.IsUserEnrolled(userID, courseID)
	if err != nil || !enrolled {
		return nil, errors.New("user not enrolled in this course")
	}

	lessons, err := s.LessonRepo.GetPublishedLessonsByCourse(courseID)
	if err != nil {
		return nil, err
	}

	userLessons, err := s.LessonRepo.GetUserLessonProgress(userID, courseID)
	if err != nil {
		return nil, err
	}

	var responses []dto.LessonResponse
	for _, lesson := range lessons {
		isCompleted := false

		// Check if lesson is completed by user
		for _, ul := range userLessons {
			if ul.LessonID == lesson.ID && ul.IsCompleted {
				isCompleted = true
				break
			}
		}

		response := s.mapLessonToResponse(&lesson, isCompleted)
		responses = append(responses, *response)
	}

	return responses, nil
}

func (s *LessonServiceImp) MarkLessonCompleted(userID, lessonID uint, watchTime int) (*dto.APIResponse, error) {
	// Get lesson details
	lesson, err := s.LessonRepo.GetByID(lessonID)
	if err != nil {
		return &dto.APIResponse{
			Success: false,
			Error:   "Lesson not found",
		}, err
	}

	// Check if user is enrolled in the course
	enrolled, err := s.UserCourseRepo.IsUserEnrolled(userID, lesson.CourseID)
	if err != nil || !enrolled {
		return &dto.APIResponse{
			Success: false,
			Error:   "User not enrolled in this course",
		}, errors.New("user not enrolled")
	}

	err = s.LessonRepo.MarkLessonCompleted(userID, lessonID, lesson.CourseID, watchTime)
	if err != nil {
		return &dto.APIResponse{
			Success: false,
			Error:   "Failed to mark lesson as completed",
		}, err
	}

	return &dto.APIResponse{
		Success: true,
		Message: "Lesson marked as completed",
	}, nil
}

// Helper methods
func (s *LessonServiceImp) mapLessonToResponse(lesson *domain.Lesson, isCompleted bool) *dto.LessonResponse {
	return &dto.LessonResponse{
		ID:          lesson.ID,
		Title:       lesson.Title,
		Description: lesson.Description,
		VideoURL:    lesson.VideoURL,
		VideoID:     lesson.VideoID,
		Script:      lesson.Script,
		Duration:    lesson.Duration,
		CourseID:    lesson.CourseID,
		Sequence:    lesson.Sequence,
		IsPublished: lesson.IsPublished,
		IsFree:      lesson.IsFree,
		CreatedAt:   lesson.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   lesson.UpdatedAt.Format(time.RFC3339),
		IsCompleted: isCompleted,
	}
}
