package domain

import "time"

type Lesson struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	VideoURL  string // YouTube link
	Script    string // Full script/text
	CourseID  uint
	Sequence  int // Order of appearance
	CreatedAt time.Time
}
