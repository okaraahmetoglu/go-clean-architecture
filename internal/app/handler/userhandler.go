package handler

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/usecase"
)

// UserHandler handles HTTP requests for user-related actions
type UserHandler struct {
	useCase *usecase.UserUseCase
}

// NewUserHandler initializes a UserHandler
func NewUserHandler(useCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

// Create creates a new user
