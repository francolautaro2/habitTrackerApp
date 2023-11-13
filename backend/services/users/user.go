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

func (u *UserClient) CreateUser() (string, error) {

}

func (u *UserClient) DeleteUser(id string) {

}

func (u *UserClient) UpdateUser(user UserClient) {

}

func (u *UserClient) GetUser() (UserClient, error) {

}

func (u *UserClient) GetAllUsers() []UserClient {

}
