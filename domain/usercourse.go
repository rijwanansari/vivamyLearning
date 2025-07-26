package domain

import "time"

type UserCourse struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	CourseID   uint
	LastLesson uint    // Track last viewed lesson
	Progress   float64 // % completed
	UpdatedAt  time.Time
}
