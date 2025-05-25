package entity

import "time"

type DigimonSpecies struct {
	ID            string    `gorm:"primaryKey;type:char(36)" json:"id"` // Assuming UUIDs
	Name          string    `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Attribute     string    `gorm:"size:20;not null" json:"attribute"` // e.g., Data, Vaccine, Virus
	Stage         string    `gorm:"size:20;not null" json:"stage"`     // e.g., Rookie, Champion
	BaseAttack    int       `gorm:"not null" json:"base_attack"`
	BaseDefense   int       `gorm:"not null" json:"base_defense"`
	BaseSpeed     int       `gorm:"not null" json:"base_speed"`
	SpriteURL     string    `gorm:"size:255" json:"sprite_url,omitempty"`
	Description   string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// Consider EvolutionInfo or PrerequisiteForEvolution fields/tables later
}
