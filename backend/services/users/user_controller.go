package users

import (
	"github.com/gin-gonic/gin"
)

// Define a structure for user controllers
type UserController struct {
	UserRepository UserRepository
}

// Controller for creating a user
func (controller *UserController) CreateUser(c *gin.Context) {
	var user UserClient
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	id, err := controller.UserRepository.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, gin.H{"id": id})
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
	var user UserClient
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
