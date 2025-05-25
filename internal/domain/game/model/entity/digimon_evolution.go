package entity

import "time"

type DigimonEvolution struct {
	ID              string    `gorm:"primaryKey;type:char(36)" json:"id"` // Or int auto_increment if preferred for this table
	FromSpeciesID   string    `gorm:"type:char(36);index;not null" json:"from_species_id"`
	ToSpeciesID     string    `gorm:"type:char(36);index;not null" json:"to_species_id"`
	Method          string    `gorm:"size:50;not null" json:"method"` // e.g., "LevelUp", "ItemUse", "DNA"
	RequiredLevel   int       `gorm:"null" json:"required_level,omitempty"` // Nullable if not level based
	RequiredItemID  string    `gorm:"type:char(36);null" json:"required_item_id,omitempty"` // Nullable, FK to items.id
	Notes           string    `gorm:"type:text" json:"notes,omitempty"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// GORM relations (optional here, can be handled by explicit joins or preloads in repositories)
	// FromSpecies *DigimonSpecies `gorm:"foreignKey:FromSpeciesID" json:"-"`
	// ToSpecies   *DigimonSpecies `gorm:"foreignKey:ToSpeciesID" json:"-"`
	// RequiredItem *Item          `gorm:"foreignKey:RequiredItemID" json:"-"`
}
