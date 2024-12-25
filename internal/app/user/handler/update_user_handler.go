package handler

import (
	"fmt"

	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
)

type UpdateUserRequest struct {
	User dto.UserDTO
	Id   int
}

type UpdateUserHandler struct {
	userService service.UserService
}

func NewUpdateUserHandler(userService service.UserService) *UpdateUserHandler {
	return &UpdateUserHandler{userService: userService}
}

func (h *UpdateUserHandler) Handle(request UpdateUserRequest) (string, error) {
	h.userService.Update(request.Id, request.User)
	return fmt.Sprintf("User %s updated!", request.User.Name), nil
}
