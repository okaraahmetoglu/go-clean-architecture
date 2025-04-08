package handler

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/mediator"
)

type GetUserByIdRequest struct {
	Id int
}

type GetUserByIdHandler struct {
	userService *service.UserService
	logger      *logger.Logger
}

func NewGetUserByIdUHandler(userService *service.UserService, logger *logger.Logger) *GetUserByIdHandler {
	return &GetUserByIdHandler{userService: userService, logger: logger}
}

func (h *GetUserByIdHandler) Handle(request GetUserByIdRequest) (interface{}, error) {
	userDto, err := h.userService.GetByID(request.Id)
	return mediator.Response[dto.UserDTO]{Data: userDto}, err
}
