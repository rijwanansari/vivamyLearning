package domain

import "time"

type Course struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Title            string    `gorm:"not null" json:"title" validate:"required"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	Thumbnail        string    `json:"thumbnail"`
	Level            string    `gorm:"default:'beginner'" json:"level"` // beginner, intermediate, advanced
	Category         string    `json:"category"`
	Tags             string    `json:"tags"`     // comma-separated tags
	Duration         int       `json:"duration"` // total duration in minutes
	Price            float64   `gorm:"default:0" json:"price"`
	IsPublished      bool      `gorm:"default:false" json:"is_published"`
	CreatedBy        uint      `json:"created_by"` // Admin ID
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	// Relationships
	Lessons     []Lesson     `gorm:"foreignKey:CourseID" json:"lessons,omitempty"`
	UserCourses []UserCourse `gorm:"foreignKey:CourseID" json:"user_courses,omitempty"`

	// Computed fields (not stored in DB)
	LessonCount    int     `gorm:"-" json:"lesson_count,omitempty"`
	EnrolledCount  int     `gorm:"-" json:"enrolled_count,omitempty"`
	CompletionRate float64 `gorm:"-" json:"completion_rate,omitempty"`
}
