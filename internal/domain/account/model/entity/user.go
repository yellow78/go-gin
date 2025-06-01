package entity

import "time"

type User struct {
	ID           string    `gorm:"primaryKey;type:char(36)" json:"id"`
	Username     string    `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Email        string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"` // Hide from JSON responses
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"` // Automatically updated by GORM

	// Game-specific fields
	InGameName       string    `gorm:"uniqueIndex;size:50;not null" json:"in_game_name"`
	Level            int       `gorm:"default:1" json:"level"`
	ExperiencePoints int       `gorm:"default:0" json:"experience_points"`
	LastLoginAt      time.Time `gorm:"null" json:"last_login_at,omitempty"`
}
