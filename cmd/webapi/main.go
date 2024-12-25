package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/config"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/container"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/mediator"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/server"
)

func main() {

	// Config yükleme
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config yüklenirken hata oluştu: %v", err)
	}

	container, err := container.BuildContainer()

	// Tüm bağımlılıkları kaydetme
	if err != nil {
		log.Fatalf("Bağımlılıklar kaydedilirken hata oluştu: %v", err)
	}

	// Mediator oluştur
	m := mediator.NewMediator()

	// Handler'ları otomatik register et
	if err := m.AutoRegisterHandlers("./internal/app/"); err != nil {
		log.Fatalf("Failed to auto-register handlers: %v", err)
	}

	// Router ve Logger ile sunucu başlatma
	err = container.Invoke(func(router *gin.Engine) {
		// Sunucu oluşturma
		srv := server.NewHTTPServer(cfg.Server.Port, router)

		log.Printf("Uygulama başlatılıyor: %d", cfg.Server.Port)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Sunucu başlatılamadı: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Container invoke sırasında hata oluştu: %v", err)
	}
}
