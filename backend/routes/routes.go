package routes

import (
	"habitTrackerApi/services/users"

	"github.com/gin-gonic/gin"
)

func RunRoutes(router *gin.Engine, userController *users.UserController) {
	// Registers Users Routers

	// Login Users Routers

	// Users Routers
	router.POST("/api/users", userController.CreateUser)
	router.GET("/api/users/:id", userController.GetUser)
	router.DELETE("/api/users/:id", userController.DeleteUser)
	router.PUT("/api/users/:id", userController.UpdateUser)
	router.GET("/api/users", userController.GetAllUsers)
}
