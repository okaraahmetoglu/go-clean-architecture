package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig, yapılandırmayı yükler ve döner
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading defaults from environment variables")
	}
}

// GetDatabaseURL, DATABASE_URL ortam değişkenini döner
func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
