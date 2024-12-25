package service

// GenericUseCase defines the operations for any entity, including the DTO to entity conversion
type GenericService[T any, D any, ID comparable] interface {
	Create(dto D) (ID, error)
	GetByID(id ID) (D, error)
	GetAll() ([]D, error)
	Update(id ID, dto D) error
	Delete(id ID) error
}
