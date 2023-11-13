package users

import "habitTrackerApi/services/habits"

type UserRepository interface {
	CreateUser() error
	GetUser(id string) (User, error)
	DeleteUser(id string) error
	UpdateUser(u User) error
	GetAllUsers() []User
}

type User struct {
	Id       string
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Habits   []habits.Habit
}

func CreateUser() {

}

func DeleteUser() {

}

func UpdateUser() {

}

func GetUser() {

}

func GetAllUsers() {

}
