package entity

// User, bir kullanıcıyı temsil eder
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// GetID, User entity'sinin ID'sini döner
func (u User) GetID() int {
	return u.ID
}
