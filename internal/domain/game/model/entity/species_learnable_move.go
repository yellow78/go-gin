package entity

import "time" // Added for potential CreatedAt/UpdatedAt if desired, though not strictly required by logic yet

// SpeciesLearnableMove defines a move that a Digimon species can learn,
// potentially with conditions like level.
type SpeciesLearnableMove struct {
	ID            string `gorm:"primaryKey;type:char(36)" json:"id"` // Dedicated ID for the join record
	SpeciesID     string `gorm:"type:char(36);index:idx_species_learn_move_keys,unique;not null" json:"species_id"` // FK to digimon_species.id
	MoveID        string `gorm:"type:char(36);index:idx_species_learn_move_keys,unique;not null" json:"move_id"`       // FK to digimon_moves.id
	LearnMethod   string `gorm:"size:50;not null" json:"learn_method"` // e.g., "LevelUp", "Tutor", "EggMove"
	RequiredLevel int    `gorm:"null" json:"required_level,omitempty"` // Nullable, e.g., for non-LevelUp methods
	Notes         string `gorm:"type:text" json:"notes,omitempty"`     // Any additional notes or conditions

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"` // Optional: when this learnable move entry was defined
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"` // Optional: if this definition changes

	// Optional GORM relations (can be helpful for some types of queries)
	// Species DigimonSpecies `gorm:"foreignKey:SpeciesID"`
	// Move    DigimonMove    `gorm:"foreignKey:MoveID"`
}
