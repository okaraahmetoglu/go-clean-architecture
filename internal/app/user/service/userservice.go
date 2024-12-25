package service

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/entity"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/repository"
)

type UserService struct {
	service.GenericService[entity.User, dto.UserDTO, int]
}

// NewUserUseCase creates a new UserUseCase
func NewUserService() *UserService {

	var repo = repository.NewUserRepository()

	// Yeni GenericUseCase oluşturuluyor
	genericService := service.NewGenericService[entity.User, dto.UserDTO, int](repo)

	// UserUseCase ile GenericUseCase kullanılarak döndürülen instance
	return &UserService{genericService}
}
