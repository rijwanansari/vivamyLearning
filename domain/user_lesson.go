package domain

import "time"

// UserLesson tracks individual lesson completion by users
type UserLesson struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	UserID      uint       `gorm:"not null" json:"user_id"`
	LessonID    uint       `gorm:"not null" json:"lesson_id"`
	CourseID    uint       `gorm:"not null" json:"course_id"`
	IsCompleted bool       `gorm:"default:false" json:"is_completed"`
	WatchTime   int        `gorm:"default:0" json:"watch_time"` // Time watched in seconds
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Relationships
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Lesson Lesson `gorm:"foreignKey:LessonID" json:"lesson,omitempty"`
	Course Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`
} // Add unique constraint for UserID, LessonID
