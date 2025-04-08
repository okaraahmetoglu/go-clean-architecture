package repository

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/entity"
	"gorm.io/gorm"
)

// UserRepositoryInterface defines the common methods for user repositories
type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetByID(id int) (entity.User, error)
	Create(entity entity.User) (int, error)
	Delete(id int) error
	Update(id int, entity entity.User) error
	GetUserByUsername(userName string) (entity.User, error)
}

// DbUserRepository implements UserRepositoryInterface using a database
type DbUserRepository struct {
	DbRepository[entity.User, int] // Inherit DbRepository functionality
}

// NewDbUserRepository creates a new DbUserRepository
func NewDbUserRepository(db *gorm.DB) *DbUserRepository {
	return &DbUserRepository{
		DbRepository: *NewDbRepository[entity.User, int](db),
	}
}

type InMemoryUserRepository struct {
	InMemoryRepository[entity.User, int] // Inherit InMemoryRepository functionality
}

// NewInMemoryUserRepository creates a new InMemoryUserRepository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		InMemoryRepository: *NewInMemoryRepository[entity.User, int](),
	}
}

func NewUserRepository(useInMemoryRepository bool, db *gorm.DB) UserRepository {
	var repo UserRepository

	if useInMemoryRepository {
		repo = &InMemoryUserRepository{
			InMemoryRepository: *NewInMemoryRepository[entity.User, int](),
		}
	} else {
		repo = &DbUserRepository{
			DbRepository: *NewDbRepository[entity.User, int](db),
		}
	}

	return repo

}

// GetByID retrieves a user by ID
func (repo *DbUserRepository) GetUserByUsername(userName string) (entity.User, error) {
	var user entity.User
	result := repo.DB.Where("username = ?", userName).First(&user) // Fetch by username
	return user, result.Error
}

func (repo *InMemoryUserRepository) GetUserByUsername(userName string) (entity.User, error) {
	var user entity.User
	//Iterate through the in-memory data to find the user by username
	return user, nil
}
