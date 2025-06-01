package config

import (
	"fmt"
	"os"
	"strconv"
)

// AppConfig 結構用來儲存應用程式的設定資訊
type AppConfig struct {
	DBHost     string // 資料庫主機位址
	DBPort     string // 資料庫連接埠
	DBUser     string // 資料庫使用者名稱
	DBPassword string // 資料庫密碼
	DBName     string // 資料庫名稱
	ServerPort string // 伺服器監聽的埠號

	// JWTSecretKey 為用於簽署 JWT Token 的密鑰
	// 重要說明：預設值僅供開發環境使用，
	// 正式環境中請務必透過 JWT_SECRET_KEY 環境變數設定為強而唯一的密鑰
	JWTSecretKey string
}

// LoadConfig 從環境變數載入應用程式設定，若未設定則使用預設值
func LoadConfig() (*AppConfig, error) {
	cfg := &AppConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "digimon"),
		DBPassword: getEnv("DB_PASSWORD", "digimon123"),
		DBName:     getEnv("DB_NAME", "digimon_game"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		// 載入 JWT_SECRET_KEY，若未提供則使用預設值
		// 重要說明：預設密鑰僅供開發使用，請務必於正式環境中自行設定
		JWTSecretKey: getEnv("JWT_SECRET_KEY", "your-default-super-secret-key-please-change-in-prod-!@#$%"),
	}

	// 檢查是否有必要的設定（可選）
	if cfg.DBHost == "" {
		return nil, fmt.Errorf("DB_HOST 未設定，且未提供預設值")
	}

	// 確認 JWT 密鑰是否已設定（理論上預設值已涵蓋）
	if cfg.JWTSecretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY 未設定，且未提供預設值")
	}

	// 驗證 ServerPort 是否為合法數字（建議檢查）
	if _, err := strconv.Atoi(cfg.ServerPort); err != nil {
		return nil, fmt.Errorf("無效的 SERVER_PORT：%s，必須是數字", cfg.ServerPort)
	}

	return cfg, nil
}

// getEnv 嘗試取得指定的環境變數，若未設定則回傳預設值
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
