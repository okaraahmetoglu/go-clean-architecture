package repository

import (
	"errors"
	"sync"
)

// InMemoryRepository bellek içi bir generic repository uygulaması
type InMemoryRepository[T any, ID Incrementable] struct {
	data   map[ID]T
	mu     sync.RWMutex // Concurrent erişim için
	nextID ID           // Sonraki ID için bir sayaç
}

// NewInMemoryRepository yeni bir InMemoryRepository oluşturur
func NewInMemoryRepository[T any, ID Incrementable]() *InMemoryRepository[T, ID] {
	var initialID ID // ID türü için uygun sıfır değeri (örneğin int için 0)
	return &InMemoryRepository[T, ID]{
		data:   make(map[ID]T),
		nextID: initialID, // ID sayacını 1'den başlatıyoruz
	}
}

// GetAll tüm varlıkları döndürür
func (r *InMemoryRepository[T, ID]) GetAll() ([]T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var entities []T
	for _, entity := range r.data {
		entities = append(entities, entity)
	}
	return entities, nil
}

// GetByID belirli bir ID'ye sahip varlığı döndürür
func (r *InMemoryRepository[T, ID]) GetByID(id ID) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entity, exists := r.data[id]
	if !exists {
		return entity, errors.New("entity not found")
	}
	return entity, nil
}

// Create yeni bir varlık ekler
// Create, yeni bir varlık ekler ve oluşturulan varlığın ID'sini döner
func (r *InMemoryRepository[T, ID]) Create(entity T) (ID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Yeni bir ID oluşturuyoruz
	id := r.nextID
	r.nextID = incrementID(id) // ID'yi artırıyoruz

	// Veriyi ekliyoruz
	r.data[id] = entity

	return id, nil
}

// ID'yi artırmak için yardımcı fonksiyon
func incrementID[ID Incrementable](id ID) ID {
	// ID'nin int veya int64 olduğunda artırma işlemi yapıyoruz
	var newID ID
	switch v := any(id).(type) {
	case int:
		newID = ID(v + 1)
	case int64:
		newID = ID(v + 1)
	default:
		// Diğer türler için uygun işlemleri ekleyebilirsiniz.
		newID = id
	}
	return newID
}

// Delete belirli bir ID'ye sahip varlığı siler
func (r *InMemoryRepository[T, ID]) Delete(id ID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return errors.New("entity not found")
	}
	delete(r.data, id)
	return nil
}

func (repo *InMemoryRepository[T, ID]) Update(id ID, entity T) error {
	if _, exists := repo.data[id]; !exists {
		return errors.New("entity not found")
	}
	repo.data[id] = entity
	return nil
}
