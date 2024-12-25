package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/handler"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/mediator"
)

// UserController
type UserController struct {
	Mediator *mediator.Mediator
}

// NewUserController creates a new UserController
func NewUserController(mediator *mediator.Mediator) *UserController {
	return &UserController{Mediator: mediator}
}

// CreateUser method to handle user creation
func (uc *UserController) CreateUser(c *gin.Context) {
	var userDto dto.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var request handler.CreateUserRequest
	request.User = userDto
	// Mediator'a gönderiyoruz
	response, err := uc.Mediator.Send(request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Handler'dan gelen yanıtı dönüyoruz
	c.JSON(200, gin.H{"message": response})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var userDto dto.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var request handler.UpdateUserRequest
	request.Id = userDto.ID
	request.User = userDto
	// Mediator'a gönderiyoruz
	response, err := uc.Mediator.Send(request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Handler'dan gelen yanıtı dönüyoruz
	c.JSON(200, gin.H{"message": response})
}
