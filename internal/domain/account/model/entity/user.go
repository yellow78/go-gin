package entity

import "time"

type User struct {
	ID           string    `gorm:"primaryKey;type:char(36)" json:"id"`
	Username     string    `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Email        string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"password_hash"`
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
