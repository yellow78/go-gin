package game

// If needed for any date fields in DTOs, e.g. AcquiredAt

// AcquireDigimonRequest is used when a player wants to acquire a new Digimon.
type AcquireDigimonRequest struct {
	SpeciesID string `json:"species_id" binding:"required,uuid"` // Assuming SpeciesID is a UUID
	Nickname  string `json:"nickname,omitempty"`                 // Optional initial nickname
}

// UpdateDigimonNicknameRequest is used to change a Digimon's nickname.
type UpdateDigimonNicknameRequest struct {
	Nickname string `json:"nickname" binding:"required,max=100"`
}

// DigimonSpeciesResponse represents the data for a Digimon species.
type DigimonSpeciesResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Attribute   string `json:"attribute"`
	Stage       string `json:"stage"`
	BaseAttack  int    `json:"base_attack"`
	BaseDefense int    `json:"base_defense"`
	BaseSpeed   int    `json:"base_speed"`
	SpriteURL   string `json:"sprite_url,omitempty"`
	Description string `json:"description,omitempty"`
}

// DigimonResponse represents the data for a player-owned Digimon.
type DigimonResponse struct {
	ID               string                 `json:"id"`
	PlayerID         string                 `json:"player_id"`
	Nickname         string                 `json:"nickname,omitempty"`
	CurrentLevel     int                    `json:"current_level"`
	CurrentAttack    int                    `json:"current_attack"`
	CurrentDefense   int                    `json:"current_defense"`
	CurrentSpeed     int                    `json:"current_speed"`
	ExperiencePoints int64                  `json:"experience_points"`
	AcquiredAt       string                 `json:"acquired_at"` // RFC3339 formatted
	Species          DigimonSpeciesResponse `json:"species"`     // Embed species details
}
