package auth

// LoginRequest 定義用戶登入時的請求資料結構。
// 必須提供 Email 或使用者名稱（兩者擇一），以及密碼。
type LoginRequest struct {
	EmailOrUsername string `json:"emailOrUsername" binding:"required"` // 電子郵件或使用者名稱（必填）
	Password        string `json:"password" binding:"required"`        // 密碼（必填）
}

// LoginResponse 定義登入成功時的回應資料結構。
// 回傳 JWT 金鑰、使用者 ID 及使用者名稱。
type LoginResponse struct {
	Token    string `json:"token"`    // JWT 驗證用 Token
	UserID   string `json:"userID"`   // 使用者的唯一識別 ID
	Username string `json:"username"` // 使用者名稱
}
