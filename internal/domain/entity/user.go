package entity

import "gorm.io/gorm"

// User, bir kullanıcıyı temsil eder
type User struct {
	gorm.Model
	ID    int
	Name  string
	Email string
	Age   int
}

// GetID, User entity'sinin ID'sini döner
func (u User) GetID() int {
	return u.ID
}
