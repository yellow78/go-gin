package entity

import "time"

// Assuming accountEntity could be:
// import accountEntity "go-gin/internal/domain/account/model/entity"

type Digimon struct {
	ID               string    `gorm:"primaryKey;type:char(36)" json:"id"`
	PlayerID         string    `gorm:"type:char(36);index;not null" json:"player_id"`  // Foreign key to User.ID
	SpeciesID        string    `gorm:"type:char(36);index;not null" json:"species_id"` // Foreign key to DigimonSpecies.ID
	Nickname         string    `gorm:"size:100" json:"nickname,omitempty"`
	CurrentLevel     int       `gorm:"not null;default:1" json:"current_level"`
	CurrentAttack    int       `gorm:"not null" json:"current_attack"`
	CurrentDefense   int       `gorm:"not null" json:"current_defense"`
	CurrentSpeed     int       `gorm:"not null" json:"current_speed"`
	ExperiencePoints int64     `gorm:"not null;default:0" json:"experience_points"` // Changed to int64
	AcquiredAt       time.Time `gorm:"autoCreateTime" json:"acquired_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Optional: To load related data if needed, though often handled by usecases/DTOs
	// User    *accountEntity.User `gorm:"foreignKey:PlayerID" json:"-"`
	// Species *DigimonSpecies     `gorm:"foreignKey:SpeciesID" json:"-"`
}
