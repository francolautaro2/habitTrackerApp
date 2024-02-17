package database

import (
	"habitTrackerApi/services/domains"

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
func (r *DatabaseUserRepository) CreateUser(user domains.UserClient) error {
	result := r.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get a user
func (r *DatabaseUserRepository) GetUser(id string) (domains.UserClient, error) {
	var user domains.UserClient
	result := r.Db.First(&user, "id = ?", id)
	if result.Error != nil {
		return domains.UserClient{}, result.Error
	}
	return user, nil
}

// Update user
func (r *DatabaseUserRepository) UpdateUser(user domains.UserClient) error {
	result := r.Db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete user
func (r *DatabaseUserRepository) DeleteUser(id string) error {
	result := r.Db.Delete(&domains.UserClient{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get all users from database
func (r *DatabaseUserRepository) GetAllUsers() ([]domains.UserClient, error) {
	var users []domains.UserClient
	result := r.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Get User by email or username
func (r *DatabaseUserRepository) GetUserByUsernameOrEmail(usernameOrEmail string) (domains.UserClient, error) {
	var user domains.UserClient
	result := r.Db.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user)
	if result.Error != nil {
		return domains.UserClient{}, result.Error
	}
	return user, nil
}
