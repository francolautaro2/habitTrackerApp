package database

import (
	"habitTrackerApi/services/users"

	"gorm.io/gorm"
)

type DatabaseUserRepository struct {
	Db *gorm.DB
}

func (r *DatabaseUserRepository) CreateUser(u users.UserClient) (string, error) {

	return u.Id, nil
}

func (r *DatabaseUserRepository) DeleteUser(id string) {

}

func (r *DatabaseUserRepository) UpdateUser(user users.UserClient) {

}

func (r *DatabaseUserRepository) GetUser() (users.UserClient, error) {

}

func (r *DatabaseUserRepository) GetAllUsers() []users.UserClient {

}
