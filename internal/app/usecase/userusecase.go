package usecase

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/entity"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/repository"
)

type UserUseCase struct {
	GenericUseCase[entity.User, dto.UserDTO, int]
}

// NewUserUseCase creates a new UserUseCase
func NewUserUseCase() *UserUseCase {

	var repo = repository.NewUserRepository()

	// Yeni GenericUseCase oluşturuluyor
	genericUseCase := NewGenericUseCase[entity.User, dto.UserDTO, int](repo)

	// UserUseCase ile GenericUseCase kullanılarak döndürülen instance
	return &UserUseCase{genericUseCase}
}
