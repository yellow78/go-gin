package auth

// LoginRequest defines the structure for login requests.
// It requires either an email or a username, and a password.
type LoginRequest struct {
	EmailOrUsername string `json:"emailOrUsername" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

// LoginResponse defines the structure for successful login responses.
// It includes the JWT token, UserID, and Username.
type LoginResponse struct {
	Token    string `json:"token"`
	UserID   string `json:"userID"`
	Username string `json:"username"`
}
