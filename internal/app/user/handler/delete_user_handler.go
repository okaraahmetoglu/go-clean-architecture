package handler

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
)

type DeleteUserRequest struct {
	Id int
}

type DeleteUserHandler struct {
	userService *service.UserService
	logger      *logger.Logger
}

// Dependency Injection ile handler'ı oluşturuyoruz
func NewDeleteUserHandler(userService *service.UserService, appLogger *logger.Logger) *DeleteUserHandler {
	return &DeleteUserHandler{userService: userService, logger: appLogger}
}

func (h *DeleteUserHandler) Handle(request DeleteUserRequest) (interface{}, error) {
	//h.logger.Println("DeleteUserHandler called:")

	err := h.userService.Delete(request.Id)
	if err == nil {
		return true, nil
	}
	return false, err
}
