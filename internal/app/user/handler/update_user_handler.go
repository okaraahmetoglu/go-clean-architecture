package handler

import (
	"fmt"

	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
)

type UpdateUserRequest struct {
	User dto.UserDTO
	Id   int
}

type UpdateUserHandler struct {
	userService *service.UserService
	logger      *logger.Logger
}

func NewUpdateUserHandler(userService *service.UserService, logger *logger.Logger) *UpdateUserHandler {
	return &UpdateUserHandler{userService: userService, logger: logger}
}

func (h *UpdateUserHandler) Handle(request UpdateUserRequest) (interface{}, error) {
	h.userService.Update(request.Id, request.User)
	return fmt.Sprintf("User %s updated!", request.User.Name), nil
}
