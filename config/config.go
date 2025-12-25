package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
	APIPrefix  string
	Env        string
}

func Load() *Config {
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("ERROR loading .env file: %v", err)
	}

	cfg := &Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "rootpassword"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3307"),
		DBName:     getEnv("DB_NAME", "ecom"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		APIPrefix:  getEnv("API_PREFIX", "/api/v1"),
		Env:        getEnv("ENV", "development"),
	}
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
