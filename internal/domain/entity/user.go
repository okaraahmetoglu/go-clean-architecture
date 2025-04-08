package entity

type User struct {
	ID       int    `gorm:"primaryKey"`
	UserName string `gorm:"size:30;not null"`
	Name     string `gorm:"size:30;not null"`
	SurName  string `gorm:"size:30;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string
}

// GetID, User entity'sinin ID'sini döner
func (u User) GetID() int {
	return u.ID
}
