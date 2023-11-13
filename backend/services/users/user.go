package users

import "habitTrackerApi/services/habits"

type UserRepository interface {
	CreateUser() (string, error)
	GetUser(id string) (User, error)
	DeleteUser(id string) error
	UpdateUser(u User) error
	GetAllUsers() []User
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

func (u *UserClient) DeleteUser() {

}

func (u *UserClient) UpdateUser() {

}

func (u *UserClient) GetUser() {

}

func (u *UserClient) GetAllUsers() {

}
