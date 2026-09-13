package configs

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

// DB is a shared GORM handle (same pattern as szpt_new).
// Prefer injecting *gorm.DB into repositories via constructors.
var DB *gorm.DB

// Config holds runtime settings loaded from environment variables.
// Replace defaults with your project values — keep names stable.
type Config struct {
	// Database
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	// HTTP
	HTTPPort string
	HTTPMode string // gin debug | release | test

	// External auth / IdP (optional in template; used by middleware)
	AuthURL      string
	ClientID     string
	ClientSecret string

	// Example external integrations — rename/remove per project
	ExternalAPIURL string
}

func LoadConfig() *Config {
	return &Config{
		DBUser:     getEnv("DB_USERNAME", "app"),
		DBPassword: getEnv("DB_PASSWORD", "app"),
		DBHost:     getEnv("DB_URL", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "app"),

		HTTPPort: getEnv("HTTP_PORT", "8080"),
		HTTPMode: getEnv("HTTP_MODE", "debug"),

		AuthURL:      getEnv("AUTH_URL", "https://idp.example.com/auth/realms/app/protocol/openid-connect"),
		ClientID:     getEnv("CLIENT_ID", "app_client"),
		ClientSecret: getEnv("CLIENT_SECRET", "change-me"),

		ExternalAPIURL: getEnv("EXTERNAL_API_URL", "https://api.example.com"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}
