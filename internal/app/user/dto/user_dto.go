package dto

type UserDTO struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
	Email    string `json:"email"`
}

// Response represents a generic response structure
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Optional, for additional data
}
