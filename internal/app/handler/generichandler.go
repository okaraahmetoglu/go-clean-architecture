package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Incrementable interface {
	int | int64
}

type ID int

type GenericHandler[D any, ID Incrementable] struct {
	useCase interface { // Interface, sadece ihtiyacı olan metotları kapsar
		Create(dto D) (ID, error)
		GetByID(id ID) (D, error)
		GetAll() ([]D, error)
		Update(id ID, dto D) error
		Delete(id ID) error
	}
}

func NewGenericHandler[D any, ID Incrementable](
	useCase interface {
		Create(dto D) (ID, error)
		GetByID(id ID) (D, error)
		GetAll() ([]D, error)
		Update(id ID, dto D) error
		Delete(id ID) error
	},
) *GenericHandler[D, ID] {
	return &GenericHandler[D, ID]{useCase: useCase}
}

// Create creates a new entity from the DTO
func (h *GenericHandler[D, ID]) Create(c *gin.Context) {
	var dto D
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.useCase.Create(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetByID retrieves an entity by its ID
func (h *GenericHandler[D, ID]) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var idAsID ID = ID(id)
	dtoItem, err := h.useCase.GetByID(idAsID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtoItem)
}

// GetAll retrieves all entities
func (h *GenericHandler[D, ID]) GetAll(c *gin.Context) {
	dtos, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos)
}

// Update updates an entity
func (h *GenericHandler[D, ID]) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var dto D
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.useCase.Update(ID(id), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

// Delete deletes an entity
func (h *GenericHandler[D, ID]) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.useCase.Delete(ID(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
