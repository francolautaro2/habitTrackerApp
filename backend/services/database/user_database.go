package database

import (
	"habitTrackerApi/services/users"

	"gorm.io/gorm"
)

type DatabaseUserRepository struct {
	Db *gorm.DB
}

func NewDatabaseUserRepository(db *gorm.DB) *DatabaseUserRepository {
	return &DatabaseUserRepository{
		Db: db,
	}
}

// Create user in database
func (r *DatabaseUserRepository) CreateUser(user users.UserClient) (string, error) {
	result := r.Db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Id, nil
}

// Get a user
func (r *DatabaseUserRepository) GetUser(id string) (users.UserClient, error) {
	var user users.UserClient
	result := r.Db.First(&user, "id = ?", id)
	if result.Error != nil {
		return users.UserClient{}, result.Error
	}
	return user, nil
}

// Update user
func (r *DatabaseUserRepository) UpdateUser(user users.UserClient) error {
	result := r.Db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete user
func (r *DatabaseUserRepository) DeleteUser(id string) error {
	result := r.Db.Delete(&users.UserClient{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get all users from database
func (r *DatabaseUserRepository) GetAllUsers() ([]users.UserClient, error) {
	var users []users.UserClient
	result := r.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
