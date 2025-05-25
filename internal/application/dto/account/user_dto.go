package account

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"` // RFC3339

	// Game-specific fields
	InGameName       string `json:"in_game_name"`
	Level            int    `json:"level"`
	ExperiencePoints int    `json:"experience_points"`
	LastLoginAt      string `json:"last_login_at,omitempty"` // RFC3339, omitempty
}
