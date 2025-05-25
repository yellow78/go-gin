package entity

import "time"

// PlayerDigimonKnownMove links a player's specific Digimon instance to a move it currently knows.
type PlayerDigimonKnownMove struct {
	ID              string `gorm:"primaryKey;type:char(36)" json:"id"` // Dedicated ID for the join record
	PlayerDigimonID string `gorm:"type:char(36);index:idx_pdigimon_known_move_keys,unique;not null" json:"player_digimon_id"` // FK to player_digimon.id
	MoveID          string `gorm:"type:char(36);index:idx_pdigimon_known_move_keys,unique;not null" json:"move_id"`            // FK to digimon_moves.id
	SlotPosition    int    `gorm:"null" json:"slot_position,omitempty"` // e.g., 1-4 if Digimon has limited active move slots, nullable if not used
	LearnedAt       time.Time `gorm:"autoCreateTime" json:"learned_at"` // Timestamp when the move was set/learned by this instance

	// Optional GORM relations
	// PlayerDigimon Digimon    `gorm:"foreignKey:PlayerDigimonID"`
	// Move          DigimonMove `gorm:"foreignKey:MoveID"`
}
