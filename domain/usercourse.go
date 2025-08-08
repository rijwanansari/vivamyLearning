package domain

import "time"

type UserCourse struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	UserID       uint       `gorm:"not null" json:"user_id"`
	CourseID     uint       `gorm:"not null" json:"course_id"`
	LastLessonID uint       `json:"last_lesson_id"`            // Track last viewed lesson
	Progress     float64    `gorm:"default:0" json:"progress"` // % completed (0-100)
	IsCompleted  bool       `gorm:"default:false" json:"is_completed"`
	EnrolledAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"enrolled_at"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at"`

	// Relationships
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Course Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`

	// Add unique constraint to prevent duplicate enrollments
} // Add gorm index for UserID, CourseID combination
