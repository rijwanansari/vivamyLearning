package services

import (
	"errors"
	"math"
	"strings"
	"time"

	"github.com/rijwanansari/vivaLearning/domain"
	"github.com/rijwanansari/vivaLearning/dto"
	repository "github.com/rijwanansari/vivaLearning/repositories"
)

type CourseService interface {
	// Admin operations
	CreateCourse(req dto.CreateCourseRequest, creatorID uint) (*dto.CourseResponse, error)
	UpdateCourse(id uint, req dto.UpdateCourseRequest, userID uint) (*dto.CourseResponse, error)
	DeleteCourse(id uint, userID uint) error
	GetCourseByID(id uint, userID *uint) (*dto.CourseResponse, error)
	GetAllCourses() ([]dto.CourseListResponse, error)
	GetCoursesByCreator(creatorID uint) ([]dto.CourseListResponse, error)

	// Public operations
	GetPublishedCourses(userID *uint) ([]dto.CourseListResponse, error)
	SearchCourses(filter dto.CourseFilterRequest, userID *uint) (*dto.PaginatedResponse, error)

	// User operations
	EnrollInCourse(courseID uint, userID uint) (*dto.APIResponse, error)
	UnenrollFromCourse(courseID uint, userID uint) (*dto.APIResponse, error)
	GetUserEnrolledCourses(userID uint) ([]dto.CourseListResponse, error)
	GetUserCourseProgress(courseID uint, userID uint) (*dto.UserProgressResponse, error)

	// Statistics
	GetCourseAnalytics(courseID uint) (map[string]interface{}, error)
}

type CourseServiceImp struct {
	CourseRepo     repository.CourseRepository
	UserCourseRepo repository.UserCourseRepository
	LessonRepo     repository.LessonRepository
}

func NewCourseService(courseRepo repository.CourseRepository, userCourseRepo repository.UserCourseRepository, lessonRepo repository.LessonRepository) CourseService {
	return &CourseServiceImp{
		CourseRepo:     courseRepo,
		UserCourseRepo: userCourseRepo,
		LessonRepo:     lessonRepo,
	}
}

func (s *CourseServiceImp) CreateCourse(req dto.CreateCourseRequest, creatorID uint) (*dto.CourseResponse, error) {
	course := &domain.Course{
		Title:            req.Title,
		Description:      req.Description,
		ShortDescription: req.ShortDescription,
		Thumbnail:        req.Thumbnail,
		Level:            req.Level,
		Category:         req.Category,
		Tags:             req.Tags,
		Price:            req.Price,
		IsPublished:      req.IsPublished,
		CreatedBy:        creatorID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := s.CourseRepo.Create(course)
	if err != nil {
		return nil, err
	}

	return s.mapCourseToResponse(course, nil), nil
}

func (s *CourseServiceImp) UpdateCourse(id uint, req dto.UpdateCourseRequest, userID uint) (*dto.CourseResponse, error) {
	course, err := s.CourseRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if user is the creator (authorization should be handled in controller/middleware)
	if course.CreatedBy != userID {
		return nil, errors.New("unauthorized to update this course")
	}

	// Update fields if provided
	if req.Title != nil {
		course.Title = *req.Title
	}
	if req.Description != nil {
		course.Description = *req.Description
	}
	if req.ShortDescription != nil {
		course.ShortDescription = *req.ShortDescription
	}
	if req.Thumbnail != nil {
		course.Thumbnail = *req.Thumbnail
	}
	if req.Level != nil {
		course.Level = *req.Level
	}
	if req.Category != nil {
		course.Category = *req.Category
	}
	if req.Tags != nil {
		course.Tags = *req.Tags
	}
	if req.Price != nil {
		course.Price = *req.Price
	}
	if req.IsPublished != nil {
		course.IsPublished = *req.IsPublished
	}

	course.UpdatedAt = time.Now()

	err = s.CourseRepo.Update(course)
	if err != nil {
		return nil, err
	}

	return s.mapCourseToResponse(course, nil), nil
}

func (s *CourseServiceImp) DeleteCourse(id uint, userID uint) error {
	course, err := s.CourseRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Check if user is the creator
	if course.CreatedBy != userID {
		return errors.New("unauthorized to delete this course")
	}

	return s.CourseRepo.Delete(id)
}

func (s *CourseServiceImp) GetCourseByID(id uint, userID *uint) (*dto.CourseResponse, error) {
	course, err := s.CourseRepo.GetByIDWithLessons(id)
	if err != nil {
		return nil, err
	}

	var userProgress *dto.UserProgressResponse
	if userID != nil {
		if enrolled, _ := s.UserCourseRepo.IsUserEnrolled(*userID, id); enrolled {
			if progress, err := s.UserCourseRepo.GetUserCourseProgress(*userID, id); err == nil {
				userProgress = &dto.UserProgressResponse{
					Progress:     progress.Progress,
					LastLessonID: progress.LastLessonID,
					IsCompleted:  progress.IsCompleted,
					EnrolledAt:   progress.EnrolledAt.Format(time.RFC3339),
				}
				if progress.CompletedAt != nil {
					completedAtStr := progress.CompletedAt.Format(time.RFC3339)
					userProgress.CompletedAt = &completedAtStr
				}
			}
		}
	}

	return s.mapCourseToResponse(course, userProgress), nil
}

func (s *CourseServiceImp) GetAllCourses() ([]dto.CourseListResponse, error) {
	courses, err := s.CourseRepo.List()
	if err != nil {
		return nil, err
	}

	return s.mapCoursesToListResponse(courses, nil), nil
}

func (s *CourseServiceImp) GetCoursesByCreator(creatorID uint) ([]dto.CourseListResponse, error) {
	courses, err := s.CourseRepo.GetCoursesByCreator(creatorID)
	if err != nil {
		return nil, err
	}

	return s.mapCoursesToListResponse(courses, nil), nil
}

func (s *CourseServiceImp) GetPublishedCourses(userID *uint) ([]dto.CourseListResponse, error) {
	courses, err := s.CourseRepo.GetPublishedCourses()
	if err != nil {
		return nil, err
	}

	return s.mapCoursesToListResponse(courses, userID), nil
}

func (s *CourseServiceImp) SearchCourses(filter dto.CourseFilterRequest, userID *uint) (*dto.PaginatedResponse, error) {
	// Set default pagination values
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	courses, total, err := s.CourseRepo.SearchCourses(filter)
	if err != nil {
		return nil, err
	}

	courseList := s.mapCoursesToListResponse(courses, userID)
	totalPages := int(math.Ceil(float64(total) / float64(filter.Limit)))

	return &dto.PaginatedResponse{
		Data:       courseList,
		Total:      total,
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalPages: totalPages,
	}, nil
}

func (s *CourseServiceImp) EnrollInCourse(courseID uint, userID uint) (*dto.APIResponse, error) {
	// Check if course exists and is published
	course, err := s.CourseRepo.GetByID(courseID)
	if err != nil {
		return &dto.APIResponse{
			Success: false,
			Error:   "Course not found",
		}, err
	}

	if !course.IsPublished {
		return &dto.APIResponse{
			Success: false,
			Error:   "Course is not published",
		}, errors.New("course not published")
	}

	// Enroll user
	enrollment, err := s.UserCourseRepo.EnrollUser(userID, courseID)
	if err != nil {
		return &dto.APIResponse{
			Success: false,
			Error:   "Failed to enroll in course",
		}, err
	}

	return &dto.APIResponse{
		Success: true,
		Message: "Successfully enrolled in course",
		Data:    enrollment,
	}, nil
}

func (s *CourseServiceImp) UnenrollFromCourse(courseID uint, userID uint) (*dto.APIResponse, error) {
	err := s.UserCourseRepo.UnenrollUser(userID, courseID)
	if err != nil {
		return &dto.APIResponse{
			Success: false,
			Error:   "Failed to unenroll from course",
		}, err
	}

	return &dto.APIResponse{
		Success: true,
		Message: "Successfully unenrolled from course",
	}, nil
}

func (s *CourseServiceImp) GetUserEnrolledCourses(userID uint) ([]dto.CourseListResponse, error) {
	enrollments, err := s.UserCourseRepo.GetUserEnrollments(userID)
	if err != nil {
		return nil, err
	}

	var courses []domain.Course
	for _, enrollment := range enrollments {
		courses = append(courses, enrollment.Course)
	}

	return s.mapCoursesToListResponse(courses, &userID), nil
}

func (s *CourseServiceImp) GetUserCourseProgress(courseID uint, userID uint) (*dto.UserProgressResponse, error) {
	progress, err := s.UserCourseRepo.GetUserCourseProgress(userID, courseID)
	if err != nil {
		return nil, err
	}

	response := &dto.UserProgressResponse{
		Progress:     progress.Progress,
		LastLessonID: progress.LastLessonID,
		IsCompleted:  progress.IsCompleted,
		EnrolledAt:   progress.EnrolledAt.Format(time.RFC3339),
	}

	if progress.CompletedAt != nil {
		completedAtStr := progress.CompletedAt.Format(time.RFC3339)
		response.CompletedAt = &completedAtStr
	}

	return response, nil
}

func (s *CourseServiceImp) GetCourseAnalytics(courseID uint) (map[string]interface{}, error) {
	lessonCount, enrolledCount, avgProgress, err := s.CourseRepo.GetCourseStats(courseID)
	if err != nil {
		return nil, err
	}

	totalEnrolled, totalCompleted, detailedAvgProgress, err := s.UserCourseRepo.GetCourseCompletionStats(courseID)
	if err != nil {
		return nil, err
	}

	completionRate := float64(0)
	if totalEnrolled > 0 {
		completionRate = (float64(totalCompleted) / float64(totalEnrolled)) * 100
	}

	return map[string]interface{}{
		"lesson_count":          lessonCount,
		"enrolled_count":        enrolledCount,
		"total_enrolled":        totalEnrolled,
		"total_completed":       totalCompleted,
		"completion_rate":       completionRate,
		"average_progress":      avgProgress,
		"detailed_avg_progress": detailedAvgProgress,
	}, nil
}

// Helper methods
func (s *CourseServiceImp) mapCourseToResponse(course *domain.Course, userProgress *dto.UserProgressResponse) *dto.CourseResponse {
	tags := []string{}
	if course.Tags != "" {
		for _, tag := range strings.Split(course.Tags, ",") {
			tags = append(tags, strings.TrimSpace(tag))
		}
	}

	lessonCount, enrolledCount, completionRate, _ := s.CourseRepo.GetCourseStats(course.ID)

	response := &dto.CourseResponse{
		ID:               course.ID,
		Title:            course.Title,
		Description:      course.Description,
		ShortDescription: course.ShortDescription,
		Thumbnail:        course.Thumbnail,
		Level:            course.Level,
		Category:         course.Category,
		Tags:             tags,
		Duration:         course.Duration,
		Price:            course.Price,
		IsPublished:      course.IsPublished,
		CreatedBy:        course.CreatedBy,
		CreatedAt:        course.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        course.UpdatedAt.Format(time.RFC3339),
		LessonCount:      lessonCount,
		EnrolledCount:    enrolledCount,
		CompletionRate:   completionRate,
		UserProgress:     userProgress,
	}

	// Add lessons if loaded
	if len(course.Lessons) > 0 {
		for _, lesson := range course.Lessons {
			response.Lessons = append(response.Lessons, dto.LessonResponse{
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
			})
		}
	}

	return response
}

func (s *CourseServiceImp) mapCoursesToListResponse(courses []domain.Course, userID *uint) []dto.CourseListResponse {
	var responses []dto.CourseListResponse

	for _, course := range courses {
		tags := []string{}
		if course.Tags != "" {
			for _, tag := range strings.Split(course.Tags, ",") {
				tags = append(tags, strings.TrimSpace(tag))
			}
		}

		lessonCount, enrolledCount, completionRate, _ := s.CourseRepo.GetCourseStats(course.ID)

		response := dto.CourseListResponse{
			ID:               course.ID,
			Title:            course.Title,
			ShortDescription: course.ShortDescription,
			Thumbnail:        course.Thumbnail,
			Level:            course.Level,
			Category:         course.Category,
			Tags:             tags,
			Duration:         course.Duration,
			Price:            course.Price,
			LessonCount:      lessonCount,
			EnrolledCount:    enrolledCount,
			CompletionRate:   completionRate,
		}

		// Check if user is enrolled
		if userID != nil {
			if enrolled, _ := s.UserCourseRepo.IsUserEnrolled(*userID, course.ID); enrolled {
				response.IsEnrolled = true
			}
		}

		responses = append(responses, response)
	}

	return responses
}
