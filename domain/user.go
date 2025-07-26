package domain

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Role      string // "admin" or "user"
	CreatedAt time.Time
}
