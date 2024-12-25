package handler

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/user/service"
)

type GetAllUsersRequest struct {
	RecordCount int
	PageCount   int
}

type GetAllUsersHandler struct {
	userService service.UserService
}

// Dependency Injection ile handler'ı oluşturuyoruz
func NewGetAllUsersHandler(userService service.UserService) *GetAllUsersHandler {
	return &GetAllUsersHandler{userService: userService}
}

func (h *GetAllUsersHandler) Handle(request GetAllUsersHandler) ([]dto.UserDTO, error) {

	userDtoList, err := h.userService.GetAll()
	return userDtoList, err
}
