package users

import (
	"habitTrackerApi/services/habits"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u UserClient) (string, error)
	GetUser(id string) (UserClient, error)
	DeleteUser(id string) error
	UpdateUser(u UserClient) error
	GetAllUsers() ([]UserClient, error)
}

type UserClient struct {
	gorm.Model
	Id       string
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Habits   []habits.Habit `gorm:"foreignKey:UserID"`
}

// All User controllers here
