package controller

import (
	"go.uber.org/dig"

	"github.com/gin-gonic/gin"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/handler"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/mediator"
)

// UserController
type UserController struct {
	container *dig.Container
}

// NewUserController creates a new UserController
func NewUserController(container *dig.Container) *UserController {
	return &UserController{container: container}
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Handles the creation of a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.UserDTO  true  "User Data"
// @Success 200 {object} dto.Response "Success"
// @Failure 400 {object} dto.Response "Bad Request"
// @Failure 500 {object} dto.Response "Internal Server Error"
// @Router       /users/create [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var userDto dto.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var request handler.CreateUserRequest
	request.User = userDto
	var err error
	var response interface{}

	uc.container.Invoke(func(h handler.CreateUserHandler) {
		response, err = h.Handle(request)
	})

	if err != nil {
		c.JSON(500, dto.Response{
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(200, dto.Response{
		Message: "Success",
		Data:    response,
	})
}

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Updates the information of an existing user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.UserDTO  true  "Updated User Data"
// @Success 200 {object} dto.Response "Success"
// @Failure 400 {object} dto.Response "Bad Request"
// @Failure 500 {object} dto.Response "Internal Server Error"
// @Router       /users/update [post]
func (uc *UserController) UpdateUser(c *gin.Context) {
	var userDto dto.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, dto.Response{
			Message: "nvalid request",
			Data:    err.Error(),
		})
		return
	}

	var request handler.UpdateUserRequest
	request.Id = userDto.ID
	request.User = userDto
	var err error
	var response interface{}

	uc.container.Invoke(func(h handler.UpdateUserHandler) {
		response, err = h.Handle(request)
	})

	if err != nil {
		c.JSON(500, dto.Response{
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(200, dto.Response{
		Message: "Success",
		Data:    response,
	})
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieves a list of all users
// @Tags         Users
// @Produce      json
// @Success 200 {object} dto.ResponseList  "Success"
// @Failure 500 {object} dto.Response "Internal Server Error"
// @Router       /users/getall [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	var request handler.GetAllUsersRequest
	var err error
	var response mediator.Response[[]dto.UserDTO]

	uc.container.Invoke(func(h handler.GetAllUsersHandler) {
		response, err = h.Handle(request)

	})
	if err != nil {
		c.JSON(500, dto.Response{
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(200, handler.ResponseList{Data: response.Data})
}
