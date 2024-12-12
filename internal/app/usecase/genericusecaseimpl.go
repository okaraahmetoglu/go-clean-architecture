package usecase

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/repository"
)

// GenericUseCaseImpl is a concrete implementation of the GenericUseCase for any entity
type GenericUseCaseImpl[T any, D any, ID comparable] struct {
	// Repository to handle entity persistence (database, in-memory, etc.)
	// repository Repository[T, ID]
	repository repository.GenericRepository[T, ID]
}

// NewGenericUseCase creates a new instance of GenericUseCaseImpl with a repository
func NewGenericUseCase[T any, D any, ID comparable](repo repository.GenericRepository[T, ID]) *GenericUseCaseImpl[T, D, ID] {
	return &GenericUseCaseImpl[T, D, ID]{repository: repo}
}

// Create creates a new entity from DTO
func (uc *GenericUseCaseImpl[T, D, ID]) Create(dtoItem D) (ID, error) {
	// Convert DTO to entity
	var entity T
	err := dto.DTOToEntity(dtoItem, &entity)
	if err != nil {
		return *new(ID), err
	}

	// Save the entity to the repository
	return uc.repository.Create(entity)
}

// GetByID retrieves an entity by ID
func (uc *GenericUseCaseImpl[T, D, ID]) GetByID(id ID) (D, error) {
	var entityDto D
	entity, err := uc.repository.GetByID(id)
	if err == nil {
		err = dto.EntityToDTO(entity, entityDto)
	}
	return entityDto, err
}

// GetAll retrieves all entities
func (uc *GenericUseCaseImpl[T, D, ID]) GetAll() ([]D, error) {
	var entityDtoList []D
	entityList, err := uc.repository.GetAll()
	if err == nil {
		err = dto.EntityToDTO(entityList, entityDtoList)
	}
	return entityDtoList, err
}

// Update updates an entity based on DTO
func (uc *GenericUseCaseImpl[T, D, ID]) Update(id ID, dtoItem D) error {
	// Convert DTO to entity
	var entity T
	err := dto.DTOToEntity(dtoItem, &entity)
	if err != nil {
		return err
	}

	// Update the entity in the repository
	return uc.repository.Update(id, entity)
}

// Delete deletes an entity by ID
func (uc *GenericUseCaseImpl[T, D, ID]) Delete(id ID) error {
	return uc.repository.Delete(id)
}
