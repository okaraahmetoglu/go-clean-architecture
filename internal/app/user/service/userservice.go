package service

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/dtomapper"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/entity"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/repository"
)

type UserService struct {
	service.GenericService[entity.User, dto.UserDTO, int]
	userRepo repository.UserRepository
}

// NewUserUseCase creates a new UserUseCase
func NewUserService(repo repository.UserRepository) *UserService {

	//var repo = repository.NewUserRepository()

	// Yeni GenericUseCase oluşturuluyor
	genericService := service.NewGenericService[entity.User, dto.UserDTO, int](repo)
	userRepo := repo

	// UserUseCase ile GenericUseCase kullanılarak döndürülen instance
	return &UserService{genericService, userRepo}
}

// GetByID retrieves an entity by ID
func (us *UserService) GetUserByUsername(userName string) (dto.UserDTO, error) {
	var entityDto dto.UserDTO
	entity, err := us.userRepo.GetUserByUsername(userName)
	if err == nil {
		err = dtomapper.EntityToDTO(entity, entityDto)
	}
	return entityDto, err
}
