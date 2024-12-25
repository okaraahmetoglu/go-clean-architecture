package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitGorm, Gorm ile PostgreSQL bağlantısını başlatır
func InitGorm() *gorm.DB {
	// DATABASE_URL'i çevreden alıyoruz
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Gorm'u başlat
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log seviyesini ayarla
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Veritabanı bağlantısını test et
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database handle: %v", err)
	}

	// Veritabanı bağlantısı ayarları
	sqlDB.SetMaxOpenConns(10)                  // Maksimum açık bağlantı sayısı
	sqlDB.SetMaxIdleConns(5)                   // Maksimum boşta bekleyen bağlantı sayısı
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Bağlantı maksimum ömrü

	fmt.Println("Connected to the database successfully!")
	return db
}
