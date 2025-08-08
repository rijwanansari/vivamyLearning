package domain

import "time"

type Lesson struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title" validate:"required"`
	Description string    `json:"description"`
	VideoURL    string    `json:"video_url"` // YouTube link or video file URL
	VideoID     string    `json:"video_id"`  // YouTube video ID
	Script      string    `json:"script"`    // Full script/text content
	Duration    int       `json:"duration"`  // Duration in seconds
	CourseID    uint      `gorm:"not null" json:"course_id" validate:"required"`
	Sequence    int       `gorm:"not null" json:"sequence"` // Order of appearance
	IsPublished bool      `gorm:"default:false" json:"is_published"`
	IsFree      bool      `gorm:"default:false" json:"is_free"` // Preview lesson
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Course Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`

	// Computed fields (not stored in DB)
	IsCompleted bool `gorm:"-" json:"is_completed,omitempty"` // For user context
}
