package repository

import (
	"github.com/rijwanansari/vivaLearning/domain"
	"gorm.io/gorm"
)

type LessonRepository interface {
	// Basic CRUD operations
	Create(lesson *domain.Lesson) error
	GetByID(id uint) (*domain.Lesson, error)
	Update(lesson *domain.Lesson) error
	Delete(id uint) error

	// Course-specific operations
	GetLessonsByCourse(courseID uint) ([]domain.Lesson, error)
	GetPublishedLessonsByCourse(courseID uint) ([]domain.Lesson, error)
	GetFreeLessonsByCourse(courseID uint) ([]domain.Lesson, error)

	// Sequence management
	GetNextSequence(courseID uint) (int, error)
	ReorderLessons(courseID uint, lessonSequences []struct {
		LessonID uint `json:"lesson_id"`
		Sequence int  `json:"sequence"`
	}) error

	// Progress tracking
	GetUserLessonProgress(userID, courseID uint) ([]domain.UserLesson, error)
	UpdateUserLessonProgress(userLesson *domain.UserLesson) error
	MarkLessonCompleted(userID, lessonID, courseID uint, watchTime int) error
}

type LessonRepositoryImp struct {
	DB *gorm.DB
}

func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &LessonRepositoryImp{DB: db}
}

func (r *LessonRepositoryImp) Create(lesson *domain.Lesson) error {
	return r.DB.Create(lesson).Error
}

func (r *LessonRepositoryImp) GetByID(id uint) (*domain.Lesson, error) {
	var lesson domain.Lesson
	err := r.DB.Preload("Course").First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (r *LessonRepositoryImp) Update(lesson *domain.Lesson) error {
	return r.DB.Save(lesson).Error
}

func (r *LessonRepositoryImp) Delete(id uint) error {
	return r.DB.Delete(&domain.Lesson{}, id).Error
}

func (r *LessonRepositoryImp) GetLessonsByCourse(courseID uint) ([]domain.Lesson, error) {
	var lessons []domain.Lesson
	err := r.DB.Where("course_id = ?", courseID).Order("sequence ASC").Find(&lessons).Error
	return lessons, err
}

func (r *LessonRepositoryImp) GetPublishedLessonsByCourse(courseID uint) ([]domain.Lesson, error) {
	var lessons []domain.Lesson
	err := r.DB.Where("course_id = ? AND is_published = ?", courseID, true).
		Order("sequence ASC").Find(&lessons).Error
	return lessons, err
}

func (r *LessonRepositoryImp) GetFreeLessonsByCourse(courseID uint) ([]domain.Lesson, error) {
	var lessons []domain.Lesson
	err := r.DB.Where("course_id = ? AND is_free = ? AND is_published = ?",
		courseID, true, true).Order("sequence ASC").Find(&lessons).Error
	return lessons, err
}

func (r *LessonRepositoryImp) GetNextSequence(courseID uint) (int, error) {
	var maxSequence struct {
		MaxSeq int
	}

	err := r.DB.Model(&domain.Lesson{}).
		Select("COALESCE(MAX(sequence), 0) as max_seq").
		Where("course_id = ?", courseID).
		Scan(&maxSequence).Error

	if err != nil {
		return 0, err
	}

	return maxSequence.MaxSeq + 1, nil
}

func (r *LessonRepositoryImp) ReorderLessons(courseID uint, lessonSequences []struct {
	LessonID uint `json:"lesson_id"`
	Sequence int  `json:"sequence"`
}) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for _, ls := range lessonSequences {
			err := tx.Model(&domain.Lesson{}).
				Where("id = ? AND course_id = ?", ls.LessonID, courseID).
				Update("sequence", ls.Sequence).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *LessonRepositoryImp) GetUserLessonProgress(userID, courseID uint) ([]domain.UserLesson, error) {
	var userLessons []domain.UserLesson
	err := r.DB.Where("user_id = ? AND course_id = ?", userID, courseID).
		Preload("Lesson").Find(&userLessons).Error
	return userLessons, err
}

func (r *LessonRepositoryImp) UpdateUserLessonProgress(userLesson *domain.UserLesson) error {
	return r.DB.Save(userLesson).Error
}

func (r *LessonRepositoryImp) MarkLessonCompleted(userID, lessonID, courseID uint, watchTime int) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// Create or update user lesson progress
		var userLesson domain.UserLesson
		err := tx.Where("user_id = ? AND lesson_id = ?", userID, lessonID).
			First(&userLesson).Error

		if err == gorm.ErrRecordNotFound {
			// Create new record
			userLesson = domain.UserLesson{
				UserID:      userID,
				LessonID:    lessonID,
				CourseID:    courseID,
				IsCompleted: true,
				WatchTime:   watchTime,
			}
			err = tx.Create(&userLesson).Error
		} else if err == nil {
			// Update existing record
			userLesson.IsCompleted = true
			userLesson.WatchTime = watchTime
			err = tx.Save(&userLesson).Error
		}

		if err != nil {
			return err
		}

		// Update course progress
		return r.updateCourseProgress(tx, userID, courseID)
	})
}

func (r *LessonRepositoryImp) updateCourseProgress(tx *gorm.DB, userID, courseID uint) error {
	// Get total lessons count for the course
	var totalLessons int64
	err := tx.Model(&domain.Lesson{}).Where("course_id = ?", courseID).Count(&totalLessons).Error
	if err != nil {
		return err
	}

	// Get completed lessons count for the user
	var completedLessons int64
	err = tx.Model(&domain.UserLesson{}).
		Where("user_id = ? AND course_id = ? AND is_completed = ?", userID, courseID, true).
		Count(&completedLessons).Error
	if err != nil {
		return err
	}

	// Calculate progress percentage
	progress := float64(0)
	if totalLessons > 0 {
		progress = (float64(completedLessons) / float64(totalLessons)) * 100
	}

	// Update user course progress
	isCompleted := progress >= 100
	err = tx.Model(&domain.UserCourse{}).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Updates(map[string]interface{}{
			"progress":     progress,
			"is_completed": isCompleted,
		}).Error

	return err
}
