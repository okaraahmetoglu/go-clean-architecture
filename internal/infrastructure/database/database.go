package database

import (
	"time"

	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/entity"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/config"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// InitGorm, Gorm ile PostgreSQL bağlantısını başlatır
func InitGorm(config *config.Config) *gorm.DB {
	appGormLogger, err := logger.NewLogger()

	appGormLogger.Println("InitGorm Başlatıldı....")

	appGormLogger.Println("Log check before gorm initialization")

	// DATABASE_URL'i çevreden alıyoruz
	dsn := config.Database.URL //os.Getenv("DATABASE_URL")
	appGormLogger.Printf("Database-Url : %s", dsn)
	if dsn == "" {
		appGormLogger.Fatalf("DATABASE_URL environment variable not set")
	}

	// Gorm'u başlat
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		appGormLogger.Fatalf("Failed to connect to database: %v", err)
	}

	appGormLogger.Println("AutoMigrate Başladı....")
	// models.AllModels içindeki tüm modelleri migrate et
	err = db.AutoMigrate(entity.AllEntities...)
	if err != nil {
		appGormLogger.Fatalf("Migration sırasında hata: %v", err)
	}

	appGormLogger.Println("AutoMigrate Tamamlandı....")

	// Veritabanı bağlantısını test et
	sqlDB, err := db.DB()
	if err != nil {
		appGormLogger.Fatalf("Failed to get database handle: %v", err)
	}

	// Veritabanı bağlantısı ayarları
	sqlDB.SetMaxOpenConns(10)                  // Maksimum açık bağlantı sayısı
	sqlDB.SetMaxIdleConns(5)                   // Maksimum boşta bekleyen bağlantı sayısı
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Bağlantı maksimum ömrü

	appGormLogger.Println("Connected to the database successfully!")
	return db
}
