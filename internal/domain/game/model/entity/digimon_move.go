package entity

import "time" // Not strictly needed for this entity, but good for consistency if other game entities use it

type DigimonMove struct {
	ID             string    `gorm:"primaryKey;type:char(36)" json:"id"` // Assuming UUIDs
	Name           string    `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Description    string    `gorm:"type:text" json:"description,omitempty"`
	Power          int       `gorm:"not null;default:0" json:"power"`
	Accuracy       int       `gorm:"not null;default:100" json:"accuracy"` // Range 0-100
	MoveType       string    `gorm:"size:50;not null" json:"move_type"`    // e.g., "Fire", "Water", "Physical", "Special", "Status"
	MPCost         int       `gorm:"not null;default:0" json:"mp_cost"`
	Target         string    `gorm:"size:50;not null" json:"target"` // e.g., "SingleEnemy", "AllEnemies", "Self", "Ally"
	BaseCritChance int       `gorm:"not null;default:5" json:"base_crit_chance"` // Range 0-100
	Effect         string    `gorm:"size:255" json:"effect,omitempty"`           // e.g., "Burn:10%", "Heal:25%", "StatBoost:Attack:Self:1_stage"
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
