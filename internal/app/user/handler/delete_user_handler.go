package handler

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
)

type DeleteUserRequest struct {
	Id int
}

type DeleteUserHandler struct {
	userService service.UserService
}

// Dependency Injection ile handler'ı oluşturuyoruz
func NewDeleteUserHandler(userService service.UserService) *DeleteUserHandler {
	return &DeleteUserHandler{userService: userService}
}

func (h *DeleteUserHandler) Handle(request DeleteUserRequest) (bool, error) {
	err := h.userService.Delete(request.Id)
	if err == nil {
		return true, nil
	}
	return false, err
}
