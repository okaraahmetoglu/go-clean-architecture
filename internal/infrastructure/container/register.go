package container

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/handler"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/repository"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/config"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/database"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func RegisterDependencies(container *dig.Container, l *logger.Logger) error {
	l, errLogger := logger.NewLogger()
	if errLogger != nil {
		l.Fatalf(errLogger.Error())
	}
	l.Info("RegisterDependencies başladı")

	err := container.Provide(func() *logger.Logger {
		logger, _ := logger.NewLogger()
		if logger == nil {
			panic("Logger is nil!")
		}
		return logger
	})
	if err != nil {
		l.Fatalf("Logger container'a eklenirken hata oluştu: %v", err)
	}

	err = container.Provide(func() *gin.Engine {
		router := gin.Default() // Gin'in varsayılan router'ı

		return router
	})
	if err != nil {
		return fmt.Errorf("Gin router kaydedilemedi: %v", err)
	}

	l.Info("gin.Engine provided")

	err = container.Provide(func() *config.Config {
		jSonConfig, _ := config.Load()
		return jSonConfig
	})
	if err != nil {
		l.Fatalf("Error providing db: %v", err)
	}

	l.Info("config.Config provided")

	// DB bağlantısını konteynere kaydet
	err = container.Provide(func(config *config.Config) *gorm.DB {
		db := database.InitGorm(config)
		if db == nil { // DB bağlantısı başarısız olursa
			panic("DB connection failed!")
		}
		return db
	})
	if err != nil {
		fmt.Printf("Error providing db: %v", err)
	}
	l.Info("gorm.DB provided")

	// Repository'yi konteynere kaydet
	err = container.Provide(func(config *config.Config, db *gorm.DB) repository.UserRepository {

		userRepository := repository.NewUserRepository(config.UseInMemoryRepository, db)
		if userRepository == nil {
			panic("UserRepository is nil!")
		}
		return userRepository
	})

	if err != nil {
		l.Fatalf("Error providing UserRepository: %v", err)
	}
	l.Info("UserRepository provided")

	// UserService'i konteynere kaydet
	err = container.Provide(func(userRepo repository.UserRepository) *service.UserService {
		if userRepo == nil {
			panic("UserRepository (userRepo) is nil!")
		}
		userService := service.NewUserService(userRepo)
		if userService == nil {
			panic("UserService is nil!")
		}
		return userService
	})

	if err != nil {
		l.Fatalf("Error providing UserService: %v", err)
	}

	container.Provide(func(service *service.UserService, logger *logger.Logger) *handler.CreateUserHandler {
		return handler.NewCreateUserHandler(service, logger)
	})

	container.Provide(func(service *service.UserService, logger *logger.Logger) *handler.DeleteUserHandler {
		return handler.NewDeleteUserHandler(service, logger)
	})

	container.Provide(func(service *service.UserService, logger *logger.Logger) *handler.UpdateUserHandler {
		return handler.NewUpdateUserHandler(service, logger)
	})

	container.Provide(func(service *service.UserService, logger *logger.Logger) *handler.GetAllUsersHandler {
		return handler.NewGetAllUsersHandler(service, logger)
	})

	container.Provide(func(service *service.UserService, logger *logger.Logger) *handler.GetUserByIdHandler {
		return handler.NewGetUserByIdUHandler(service, logger)
	})

	/*
		// Handler'ları kaydet
		container.Provide(func(userService *service.UserService) *handler.CreateUserHandler {
			return handler.NewCreateUserHandler(userService)
		})

		container.Provide(func(userService *service.UserService) *handler.DeleteUserHandler {
			return handler.NewDeleteUserHandler(userService)
		})

		container.Provide(func(userService *service.UserService) *handler.UpdateUserHandler {
			return handler.NewUpdateUserHandler(userService)
		})

		container.Provide(func(userService *service.UserService) *handler.GetAllUsersHandler {
			return handler.NewGetAllUsersHandler(userService)
		})

		container.Provide(func(userService *service.UserService) *handler.GetUserByIdHandler {
			return handler.NewGetUserByIdUHandler(userService)
		})
	*/
	// HandlerRegistry'yi konteynere kaydet
	//container.Provide(func(container *dig.Container) *mediator.HandlerRegistry {
	//	return mediator.NewHandlerRegistry(container)
	//})

	// Mediator'ü konteynere kaydet
	/*
		container.Provide(func() *mediator.Mediator {
			return mediator.NewMediator(container)
		})

		if err := container.Invoke(func(m *mediator.Mediator) {


			m.Registry.Register(handler.CreateUserRequest{}, func(service *service.UserService) *handler.CreateUserHandler {
				return handler.NewCreateUserHandler(service)
			})

			m.Registry.Register(handler.DeleteUserRequest{}, func(service *service.UserService) *handler.DeleteUserHandler {
				return handler.NewDeleteUserHandler(service)
			})

			m.Registry.Register(handler.UpdateUserRequest{}, func(service *service.UserService) *handler.UpdateUserHandler {
				return handler.NewUpdateUserHandler(service)
			})

			m.Registry.Register(handler.GetAllUsersRequest{}, func(service *service.UserService) *handler.GetAllUsersHandler {
				return handler.NewGetAllUsersHandler(service)
			})

			m.Registry.Register(handler.GetUserByIdRequest{}, func(service *service.UserService) *handler.GetUserByIdHandler {
				return handler.NewGetUserByIdUHandler(service)
			})*/

	/*m.Registry.Register(handler.CreateUserRequest{}, &handler.CreateUserHandler{})
		m.Registry.Register(handler.DeleteUserRequest{}, &handler.DeleteUserHandler{})
		m.Registry.Register(handler.UpdateUserRequest{}, &handler.UpdateUserHandler{})
		m.Registry.Register(handler.GetAllUsersRequest{}, &handler.GetAllUsersHandler{})
		m.Registry.Register(handler.GetUserByIdRequest{}, &handler.GetUserByIdHandler{})
	}); err != nil {
		return fmt.Errorf("failed to register mediator: %v", err)
	}
	*/

	l.Info("All provided")
	testRegister(container, l)

	return nil
}

func testRegister(container *dig.Container, appLogger *logger.Logger) error {

	var err error
	var response interface{}

	appLogger.Info("testRegister called")

	container.Invoke(func(h *handler.CreateUserHandler) {
		appLogger, _ = logger.NewLogger()

		appLogger.Printf("CreateUserHandler %v", h)
		response, err = h.Handle(handler.CreateUserRequest{User: dto.UserDTO{Name: "Osman", Email: "ok"}})

		if err != nil {
			appLogger.Fatalf("Hata : %v", err)
		}
	})

	if err != nil {
		appLogger.Fatalf("Failed to invoke CreateUserHandler: %v", err)
	} else {
		//appLogger.Printf("CreateUserRequest Response: ")
		appLogger.Printf("CreateUserRequest Response: %v", response)
	}

	appLogger.Info("DeleteUserHandler calling")
	container.Invoke(func(h *handler.DeleteUserHandler) {
		response, err = h.Handle(handler.DeleteUserRequest{Id: 3})
	})
	appLogger.Info("DeleteUserHandler called but error not checked.")
	if err != nil {
		appLogger.Printf("DeleteUserRequest Error: %v", err)
	} else {
		appLogger.Printf("DeleteUserRequest Response: %v", response)
	}
	appLogger.Info("DeleteUserHandler called")

	container.Invoke(func(h *handler.GetAllUsersHandler) {
		response, err = h.Handle(handler.GetAllUsersRequest{})
	})

	if err != nil {
		appLogger.Printf("GetAllUsersRequest Error: %v", err)
	} else {
		appLogger.Printf("GetAllUsersRequest Response: %v", response)
	}
	return nil
}
