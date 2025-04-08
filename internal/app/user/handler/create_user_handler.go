package handler

import (
	"fmt"

	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
)

type CreateUserRequest struct {
	User dto.UserDTO
}

// Dependency Injection ile handler'ı oluşturuyoruz
func NewCreateUserHandler(userService *service.UserService, logger *logger.Logger) *CreateUserHandler {
	return &CreateUserHandler{userService: userService, logger: logger}
}

type CreateUserHandler struct {
	userService *service.UserService
	logger      *logger.Logger
}

func (h *CreateUserHandler) Handle(request CreateUserRequest) (interface{}, error) {

	if h.logger == nil {

		h.logger, _ = logger.NewLogger()
	}

	//h.logger.Println("CreateUserHandler called:")
	h.logger.Println("CreateUserHandler called:")

	//h.logger.Printf("createUserRequest : %s", request.User.Name)*/
	h.logger.Printf("createUserRequest : %s", request.User.Name)

	if h == nil {
		//h.logger.Printf("h is nil")
		h.logger.Printf("h is nil")
	}
	if h.userService == nil {
		h.logger.Printf("h.userService is nil")
	}
	userId, err := h.userService.Create(request.User)
	if err == nil {
		h.logger.Printf("User %s created!", request.User.Name)
		return fmt.Sprintf("User id: %d  name: %s created!", userId, request.User.Name), nil
	} else {
		h.logger.Fatalf("User %d %s not created! Hata: %v", userId, request.User.Name, err)
		return nil, fmt.Errorf("User %d %s not created! Hata: %v", userId, request.User.Name, err)
	}

}
