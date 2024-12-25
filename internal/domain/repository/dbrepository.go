package repository

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

// Incrementable arayüzü ID'lerin arttırılabilir olduğunu belirtir

// DbRepository GORM ile veritabanı işlemleri yapan generic repository
type DbRepository[T any, ID Incrementable] struct {
	DB *gorm.DB
}

// NewDbRepository GORM ile DB bağlantısı kurarak yeni bir DbRepository oluşturur
func NewDbRepository[T any, ID Incrementable](db *gorm.DB) *DbRepository[T, ID] {
	return &DbRepository[T, ID]{DB: db}
}

// GetAll tüm varlıkları veritabanından getirir
func (r *DbRepository[T, ID]) GetAll() ([]T, error) {
	var entities []T
	result := r.DB.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

// GetByID belirli bir ID'ye sahip varlığı getirir
func (r *DbRepository[T, ID]) GetByID(id ID) (T, error) {
	var entity T
	result := r.DB.First(&entity, id)
	if result.Error != nil {
		return entity, errors.New("entity not found")
	}
	return entity, nil
}

// Create yeni bir varlık ekler ve oluşturulan varlığın ID'sini döner
func (r *DbRepository[T, ID]) Create(entity T) (ID, error) {
	result := r.DB.Create(&entity)
	if result.Error != nil {
		var zeroID ID
		return zeroID, result.Error
	}

	// Reflect ile ID alanını okumak
	val := reflect.ValueOf(entity)
	idField := val.FieldByName("ID") // ID alanı varsa alınır
	if !idField.IsValid() {
		var zeroID ID
		return zeroID, errors.New("entity does not have an ID field")
	}

	// ID alanını döndür
	id, ok := idField.Interface().(ID)
	if !ok {
		var zeroID ID
		return zeroID, errors.New("ID field is not of the expected type")
	}

	return id, nil
}

// Delete belirli bir ID'ye sahip varlığı siler
func (r *DbRepository[T, ID]) Delete(id ID) error {
	var entity T
	result := r.DB.Delete(&entity, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("entity not found")
	}
	return nil
}

// Update varlığı günceller
func (r *DbRepository[T, ID]) Update(id ID, entity T) error {
	var existingEntity T
	result := r.DB.First(&existingEntity, id)
	if result.Error != nil {
		return errors.New("entity not found")
	}

	// GORM ile varlık güncelleniyor
	result = r.DB.Save(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
