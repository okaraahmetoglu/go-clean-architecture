package repository

// Artırılabilir türler için bir tür kısıtlaması tanımlıyoruz
type Incrementable interface {
	int | int32 | int64 | uint | uint32 | uint64
}

// GenericRepository generic bir repository arayüzü
type GenericRepository[T any, ID comparable] interface {
	GetAll() ([]T, error)
	GetByID(id ID) (T, error)
	Create(entity T) (ID, error)
	Delete(id ID) error
	Update(id ID, entity T) error
}

// Create implements GenericRepository.
