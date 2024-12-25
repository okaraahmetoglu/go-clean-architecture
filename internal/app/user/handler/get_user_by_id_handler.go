package handler

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
)

type GetUserByIdRequest struct {
	Id int
}

type GetUserByIdUHandler struct {
	userService service.UserService
}

func NewGetUserByIdUHandler(userService service.UserService) *GetUserByIdUHandler {
	return &GetUserByIdUHandler{userService: userService}
}

func (h *GetUserByIdUHandler) Handle(request GetUserByIdRequest) (dto.UserDTO, error) {
	userDto, err := h.userService.GetByID(request.Id)
	return userDto, err
}
