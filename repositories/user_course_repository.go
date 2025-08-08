package repository

import (
	"time"

	"github.com/rijwanansari/vivaLearning/domain"
	"gorm.io/gorm"
)

type UserCourseRepository interface {
	// Enrollment operations
	EnrollUser(userID, courseID uint) (*domain.UserCourse, error)
	UnenrollUser(userID, courseID uint) error
	IsUserEnrolled(userID, courseID uint) (bool, error)
	GetUserCourseProgress(userID, courseID uint) (*domain.UserCourse, error)

	// Progress management
	UpdateProgress(userID, courseID uint, progress float64, lastLessonID uint) error
	MarkCourseCompleted(userID, courseID uint) error

	// User's learning analytics
	GetUserEnrollments(userID uint) ([]domain.UserCourse, error)
	GetUserCompletedCourses(userID uint) ([]domain.UserCourse, error)
	GetUserInProgressCourses(userID uint) ([]domain.UserCourse, error)

	// Course analytics
	GetCourseEnrollments(courseID uint) ([]domain.UserCourse, error)
	GetCourseCompletionStats(courseID uint) (totalEnrolled int, totalCompleted int, avgProgress float64, err error)
}

type UserCourseRepositoryImp struct {
	DB *gorm.DB
}

func NewUserCourseRepository(db *gorm.DB) UserCourseRepository {
	return &UserCourseRepositoryImp{DB: db}
}

func (r *UserCourseRepositoryImp) EnrollUser(userID, courseID uint) (*domain.UserCourse, error) {
	// Check if already enrolled
	var existingEnrollment domain.UserCourse
	err := r.DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&existingEnrollment).Error

	if err == nil {
		// Already enrolled, return existing enrollment
		return &existingEnrollment, nil
	}

	if err != gorm.ErrRecordNotFound {
		// Database error
		return nil, err
	}

	// Create new enrollment
	enrollment := &domain.UserCourse{
		UserID:     userID,
		CourseID:   courseID,
		Progress:   0,
		EnrolledAt: time.Now(),
	}

	err = r.DB.Create(enrollment).Error
	if err != nil {
		return nil, err
	}

	return enrollment, nil
}

func (r *UserCourseRepositoryImp) UnenrollUser(userID, courseID uint) error {
	return r.DB.Where("user_id = ? AND course_id = ?", userID, courseID).
		Delete(&domain.UserCourse{}).Error
}

func (r *UserCourseRepositoryImp) IsUserEnrolled(userID, courseID uint) (bool, error) {
	var count int64
	err := r.DB.Model(&domain.UserCourse{}).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Count(&count).Error

	return count > 0, err
}

func (r *UserCourseRepositoryImp) GetUserCourseProgress(userID, courseID uint) (*domain.UserCourse, error) {
	var userCourse domain.UserCourse
	err := r.DB.Where("user_id = ? AND course_id = ?", userID, courseID).
		Preload("Course").First(&userCourse).Error

	if err != nil {
		return nil, err
	}

	return &userCourse, nil
}

func (r *UserCourseRepositoryImp) UpdateProgress(userID, courseID uint, progress float64, lastLessonID uint) error {
	updates := map[string]interface{}{
		"progress":       progress,
		"last_lesson_id": lastLessonID,
		"updated_at":     time.Now(),
	}

	// If progress is 100%, mark as completed
	if progress >= 100 {
		completedAt := time.Now()
		updates["is_completed"] = true
		updates["completed_at"] = &completedAt
	}

	return r.DB.Model(&domain.UserCourse{}).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Updates(updates).Error
}

func (r *UserCourseRepositoryImp) MarkCourseCompleted(userID, courseID uint) error {
	completedAt := time.Now()
	return r.DB.Model(&domain.UserCourse{}).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Updates(map[string]interface{}{
			"progress":     100,
			"is_completed": true,
			"completed_at": &completedAt,
			"updated_at":   time.Now(),
		}).Error
}

func (r *UserCourseRepositoryImp) GetUserEnrollments(userID uint) ([]domain.UserCourse, error) {
	var enrollments []domain.UserCourse
	err := r.DB.Where("user_id = ?", userID).
		Preload("Course").
		Order("enrolled_at DESC").
		Find(&enrollments).Error

	return enrollments, err
}

func (r *UserCourseRepositoryImp) GetUserCompletedCourses(userID uint) ([]domain.UserCourse, error) {
	var enrollments []domain.UserCourse
	err := r.DB.Where("user_id = ? AND is_completed = ?", userID, true).
		Preload("Course").
		Order("completed_at DESC").
		Find(&enrollments).Error

	return enrollments, err
}

func (r *UserCourseRepositoryImp) GetUserInProgressCourses(userID uint) ([]domain.UserCourse, error) {
	var enrollments []domain.UserCourse
	err := r.DB.Where("user_id = ? AND is_completed = ? AND progress > ?", userID, false, 0).
		Preload("Course").
		Order("updated_at DESC").
		Find(&enrollments).Error

	return enrollments, err
}

func (r *UserCourseRepositoryImp) GetCourseEnrollments(courseID uint) ([]domain.UserCourse, error) {
	var enrollments []domain.UserCourse
	err := r.DB.Where("course_id = ?", courseID).
		Preload("User").
		Order("enrolled_at DESC").
		Find(&enrollments).Error

	return enrollments, err
}

func (r *UserCourseRepositoryImp) GetCourseCompletionStats(courseID uint) (totalEnrolled int, totalCompleted int, avgProgress float64, err error) {
	// Get total enrolled count
	var totalEnrolledInt64 int64
	err = r.DB.Model(&domain.UserCourse{}).
		Where("course_id = ?", courseID).
		Count(&totalEnrolledInt64).Error
	if err != nil {
		return 0, 0, 0, err
	}
	totalEnrolled = int(totalEnrolledInt64)

	// Get total completed count
	var totalCompletedInt64 int64
	err = r.DB.Model(&domain.UserCourse{}).
		Where("course_id = ? AND is_completed = ?", courseID, true).
		Count(&totalCompletedInt64).Error
	if err != nil {
		return 0, 0, 0, err
	}
	totalCompleted = int(totalCompletedInt64)

	// Get average progress
	var result struct {
		AvgProgress float64
	}
	err = r.DB.Model(&domain.UserCourse{}).
		Select("COALESCE(AVG(progress), 0) as avg_progress").
		Where("course_id = ?", courseID).
		Scan(&result).Error

	avgProgress = result.AvgProgress
	return totalEnrolled, totalCompleted, avgProgress, err
}
