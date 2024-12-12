package repository

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/entity"
)

// UserRepository, User entity'si için repository implementasyonudur
type UserRepository struct {
	*InMemoryRepository[entity.User, int]
}

// NewUserRepository, UserRepository için yeni bir instance oluşturur
func NewUserRepository() *UserRepository {
	return &UserRepository{
		InMemoryRepository: NewInMemoryRepository[entity.User, int](),
	}
}
