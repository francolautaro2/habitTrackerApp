package users

import (
	"habitTrackerApi/services/domains"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/*
	ALL LOGIC HTTP USERS CONTROLLERS
*/

// Define a structure for user controllers
type UserController struct {
	UserRepository domains.UserRepository
}

// Controller for create user
func (controller *UserController) CreateUser(c *gin.Context) {
	var user domains.UserClient
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// call User repository to save user
	err = controller.UserRepository.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	// Return ID in response
	c.JSON(201, gin.H{"id": user.ID})
}

// Controller for getting a user by ID
func (controller *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := controller.UserRepository.GetUser(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

// Controller for deleting a user by ID
func (controller *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := controller.UserRepository.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

// Controller for updating a user
func (controller *UserController) UpdateUser(c *gin.Context) {
	var user domains.UserClient
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err := controller.UserRepository.UpdateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(200, gin.H{"message": "User updated successfully"})
}

// Controller for getting all users
func (controller *UserController) GetAllUsers(c *gin.Context) {
	users, err := controller.UserRepository.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(200, users)
}
