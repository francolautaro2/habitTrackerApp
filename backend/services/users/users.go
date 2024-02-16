package users

import (
	"habitTrackerApi/services/habits"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u UserClient) error
	GetUser(id string) (UserClient, error)
	DeleteUser(id string) error
	UpdateUser(u UserClient) error
	GetAllUsers() ([]UserClient, error)
	GetUserByUsernameOrEmail(usernameOrEmail string) (UserClient, error)
}

type UserClient struct {
	gorm.Model
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Habits   []habits.Habit `json:"habits" gorm:"foreignKey:UserID"` // Especifica la clave foránea aquí
}
