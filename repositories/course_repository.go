package repository

import (
	"strings"

	"github.com/rijwanansari/vivaLearning/domain"
	"github.com/rijwanansari/vivaLearning/dto"
	"gorm.io/gorm"
)

type CourseRepository interface {
	// Basic CRUD operations
	Create(course *domain.Course) error
	GetByID(id uint) (*domain.Course, error)
	GetByIDWithLessons(id uint) (*domain.Course, error)
	Update(course *domain.Course) error
	Delete(id uint) error
	List() ([]domain.Course, error)

	// Advanced operations
	GetPublishedCourses() ([]domain.Course, error)
	GetCoursesByCreator(creatorID uint) ([]domain.Course, error)
	SearchCourses(filter dto.CourseFilterRequest) ([]domain.Course, int64, error)
	GetUserEnrolledCourses(userID uint) ([]domain.Course, error)

	// Statistics
	GetCourseStats(courseID uint) (lessonCount int, enrolledCount int, avgProgress float64, err error)
}

type CourseRepositoryImp struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImp{DB: db}
}

func (r *CourseRepositoryImp) Create(course *domain.Course) error {
	return r.DB.Create(course).Error
}

func (r *CourseRepositoryImp) GetByID(id uint) (*domain.Course, error) {
	var course domain.Course
	err := r.DB.First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepositoryImp) GetByIDWithLessons(id uint) (*domain.Course, error) {
	var course domain.Course
	err := r.DB.Preload("Lessons", func(db *gorm.DB) *gorm.DB {
		return db.Order("sequence ASC")
	}).First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepositoryImp) Update(course *domain.Course) error {
	return r.DB.Save(course).Error
}

func (r *CourseRepositoryImp) Delete(id uint) error {
	return r.DB.Delete(&domain.Course{}, id).Error
}

func (r *CourseRepositoryImp) List() ([]domain.Course, error) {
	var courses []domain.Course
	err := r.DB.Preload("Lessons").Order("created_at DESC").Find(&courses).Error
	return courses, err
}

func (r *CourseRepositoryImp) GetPublishedCourses() ([]domain.Course, error) {
	var courses []domain.Course
	err := r.DB.Where("is_published = ?", true).Order("created_at DESC").Find(&courses).Error
	return courses, err
}

func (r *CourseRepositoryImp) GetCoursesByCreator(creatorID uint) ([]domain.Course, error) {
	var courses []domain.Course
	err := r.DB.Where("created_by = ?", creatorID).Order("created_at DESC").Find(&courses).Error
	return courses, err
}

func (r *CourseRepositoryImp) SearchCourses(filter dto.CourseFilterRequest) ([]domain.Course, int64, error) {
	var courses []domain.Course
	var total int64

	query := r.DB.Model(&domain.Course{})

	// Apply filters
	if filter.Category != "" {
		query = query.Where("category ILIKE ?", "%"+filter.Category+"%")
	}

	if filter.Level != "" {
		query = query.Where("level = ?", filter.Level)
	}

	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}

	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}

	if filter.IsPublished != nil {
		query = query.Where("is_published = ?", *filter.IsPublished)
	}

	if filter.Search != "" {
		searchTerm := "%" + filter.Search + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ?", searchTerm, searchTerm)
	}

	if filter.Tags != "" {
		tags := strings.Split(filter.Tags, ",")
		for _, tag := range tags {
			query = query.Where("tags ILIKE ?", "%"+strings.TrimSpace(tag)+"%")
		}
	}

	// Count total records
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (filter.Page - 1) * filter.Limit
	query = query.Offset(offset).Limit(filter.Limit)

	// Apply sorting
	sortField := "created_at"
	if filter.SortBy != "" {
		sortField = filter.SortBy
	}

	sortOrder := "DESC"
	if filter.SortOrder == "asc" {
		sortOrder = "ASC"
	}

	query = query.Order(sortField + " " + sortOrder)

	err = query.Find(&courses).Error
	return courses, total, err
}

func (r *CourseRepositoryImp) GetUserEnrolledCourses(userID uint) ([]domain.Course, error) {
	var courses []domain.Course
	err := r.DB.Joins("JOIN user_courses ON user_courses.course_id = courses.id").
		Where("user_courses.user_id = ?", userID).
		Order("user_courses.enrolled_at DESC").
		Find(&courses).Error
	return courses, err
}

func (r *CourseRepositoryImp) GetCourseStats(courseID uint) (lessonCount int, enrolledCount int, avgProgress float64, err error) {
	// Get lesson count
	var lessonCountInt64 int64
	err = r.DB.Model(&domain.Lesson{}).Where("course_id = ?", courseID).Count(&lessonCountInt64).Error
	if err != nil {
		return 0, 0, 0, err
	}
	lessonCount = int(lessonCountInt64)

	// Get enrolled count
	var enrolledCountInt64 int64
	err = r.DB.Model(&domain.UserCourse{}).Where("course_id = ?", courseID).Count(&enrolledCountInt64).Error
	if err != nil {
		return 0, 0, 0, err
	}
	enrolledCount = int(enrolledCountInt64)

	// Get average progress
	var result struct {
		AvgProgress float64
	}
	err = r.DB.Model(&domain.UserCourse{}).
		Select("COALESCE(AVG(progress), 0) as avg_progress").
		Where("course_id = ?", courseID).
		Scan(&result).Error

	avgProgress = result.AvgProgress
	return lessonCount, enrolledCount, avgProgress, err
}
