package handler

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/mediator"
)

type ResponseList struct {
	Data []dto.UserDTO `json:"data"`
}

type GetAllUsersRequest struct {
	RecordCount int
	PageCount   int
}

type GetAllUsersHandler struct {
	userService *service.UserService
	logger      *logger.Logger
}

// Dependency Injection ile handler'ı oluşturuyoruz
func NewGetAllUsersHandler(userService *service.UserService, logger *logger.Logger) *GetAllUsersHandler {

	if userService == nil {
		logger.Fatalf("UserService  is null")
	}

	return &GetAllUsersHandler{userService: userService, logger: logger}
}

func (h *GetAllUsersHandler) Handle(request GetAllUsersRequest) (mediator.Response[[]dto.UserDTO], error) {
	userDtoList, err := h.userService.GetAll()
	return mediator.Response[[]dto.UserDTO]{Data: userDtoList}, err
}
