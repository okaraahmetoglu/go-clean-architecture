package dto

import (
	"github.com/mitchellh/mapstructure"
)

// EntityToDTO - Generic mapping function for Entity to DTO
func EntityToDTO(entity interface{}, dto interface{}) error {
	return mapstructure.Decode(entity, dto)
}

// DTOToEntity - Generic mapping function for DTO to Entity
func DTOToEntity(dto interface{}, entity interface{}) error {
	return mapstructure.Decode(dto, entity)
}
