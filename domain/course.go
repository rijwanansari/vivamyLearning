package domain

import "time"

type Course struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	CreatedBy   uint // Admin ID
	CreatedAt   time.Time
	Lessons     []Lesson `gorm:"foreignKey:CourseID"`
}
