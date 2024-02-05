package users

import "habitTrackerApi/services/habits"

type UserRepository interface {
	CreateUser() (string, error)
	GetUser(id string) (UserClient, error)
	DeleteUser(id string) error
	UpdateUser(u UserClient) error
	GetAllUsers() []UserClient
}

type UserClient struct {
	Id       string
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Habits   []habits.Habit
}
