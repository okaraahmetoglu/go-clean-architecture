package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/okaraahmetoglu/go-clean-architecture/docs" // Swagger docs package
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/config"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/container"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/server"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/interface/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	logger, err := logger.NewLogger()
	if err != nil {
		panic("Logger oluşturulamadı!")
	}

	defer logger.Close()

	// Config yükleme
	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Config yüklenirken hata oluştu: %v", err)
	}

	logger.Println("BuildContainer başladı...")
	container, err := container.BuildContainer(logger)
	logger.Println("Bitti...")

	// Tüm bağımlılıkları kaydetme
	if err != nil {
		log.Fatalf("Bağımlılıklar kaydedilirken hata oluştu: %v", err)
	}

	// Handler'ları otomatik register et

	// Router ve Logger ile sunucu başlatma
	err = container.Invoke(func(router *gin.Engine) {
		// Sunucu oluşturma

		userController := controller.NewUserController(container)

		router.GET("/users/getall", userController.GetAllUsers)
		router.POST("/users/create", userController.CreateUser)
		router.POST("/users/update", userController.UpdateUser)

		// Swagger route
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		srv := server.NewHTTPServer(cfg.Server.Port, router)

		//log.Printf("Uygulama başlatılıyor:")
		log.Printf("Uygulama başlatılıyor: %d", cfg.Server.Port)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Sunucu başlatılamadı: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Container invoke sırasında hata oluştu: %v", err)
	}
}
