package handler

import (
	"fmt"

	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
)

type CreateUserRequest struct {
	User dto.UserDTO
}

// Dependency Injection ile handler'ı oluşturuyoruz
func NewCreateUserHandler(userService service.UserService) *CreateUserHandler {
	return &CreateUserHandler{userService: userService}
}

type CreateUserHandler struct {
	userService service.UserService
}

func (h *CreateUserHandler) Handle(request CreateUserRequest) (string, error) {
	h.userService.Create(request.User)
	return fmt.Sprintf("User %s created!", request.User.Name), nil
}
