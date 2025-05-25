package config

import (
	"fmt"
	"os"
	"strconv"
)

// AppConfig holds the application configuration
type AppConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	// JWTSecretKey is the secret key used for signing JWT tokens.
	// IMPORTANT: The default value is for development convenience ONLY.
	// ALWAYS override this with a strong, unique secret in production environments
	// by setting the JWT_SECRET_KEY environment variable.
	JWTSecretKey string
}

// LoadConfig loads configuration from environment variables or uses default values.
func LoadConfig() (*AppConfig, error) {
	cfg := &AppConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "digimon"),
		DBPassword: getEnv("DB_PASSWORD", "digimon123"),
		DBName:     getEnv("DB_NAME", "digimon_game"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		// Load JWT_SECRET_KEY or use a default.
		// IMPORTANT: The default key below is for development convenience ONLY.
		// It MUST be changed for production environments via the JWT_SECRET_KEY env variable.
		JWTSecretKey: getEnv("JWT_SECRET_KEY", "your-default-super-secret-key-please-change-in-prod-!@#$%"),
	}

	// Example of checking for an essential config (optional, based on requirements)
	if cfg.DBHost == "" {
		return nil, fmt.Errorf("DB_HOST is not set and no default value provided")
	}

	// Ensure JWTSecretKey is set (it should be, due to the default)
	if cfg.JWTSecretKey == "" {
		// This case should ideally not be reached if a default is always provided.
		return nil, fmt.Errorf("JWT_SECRET_KEY is not set and no default value was provided")
	}
	// Add more checks if other variables are strictly required and have no defaults

	// Validate ServerPort (optional, but good practice)
	if _, err := strconv.Atoi(cfg.ServerPort); err != nil {
		return nil, fmt.Errorf("invalid SERVER_PORT: %s, must be a number", cfg.ServerPort)
	}

	return cfg, nil
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
