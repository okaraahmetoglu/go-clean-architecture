package container

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/handler"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/database"
	"go.uber.org/dig"
)

func RegisterDependencies(container *dig.Container) error {

	if err := container.Provide(func() *gin.Engine {
		router := gin.Default() // Gin'in varsayılan router'ı
		return router
	}); err != nil {
		return fmt.Errorf("Gin router kaydedilemedi: %v", err)
	}

	dependencies := []interface{}{
		database.InitGorm,
		service.NewUserService,
		handler.NewCreateUserHandler,
		handler.NewDeleteUserHandler,
		handler.NewDeleteUserHandler,
		handler.NewGetAllUsersHandler,
		handler.NewGetUserByIdUHandler,
		handler.NewUpdateUserHandler,
	}

	for _, dep := range dependencies {
		if err := container.Provide(dep); err != nil {
			return fmt.Errorf("failed to register dependency: %v", err)
		}
	}

	return nil
}
