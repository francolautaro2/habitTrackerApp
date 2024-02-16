package routes

import (
	"habitTrackerApi/services/auth"
	"habitTrackerApi/services/habits"
	"habitTrackerApi/services/users"

	"github.com/gin-gonic/gin"
)

func RunRoutes(router *gin.Engine, userController *users.UserController, authController *auth.AuthController, habitController *habits.HabitController) {
	// Registers Users Routers

	// Login Users Routers
	router.POST("/api/login", authController.Login)

	// Users Routers
	router.POST("/api/users", userController.CreateUser)
	router.GET("/api/users/:id", userController.GetUser)
	router.DELETE("/api/users/:id", userController.DeleteUser)
	router.PUT("/api/users/:id", userController.UpdateUser)
	router.GET("/api/users", userController.GetAllUsers)

	// Habit routes
	router.POST("/api/habits", auth.AuthMiddleware(), habitController.CreateHabit)
	router.GET("/api/habits/:id", auth.AuthMiddleware(), habitController.GetHabit)
	router.DELETE("/api/habits/:id", auth.AuthMiddleware(), habitController.DeleteHabit)
	router.PUT("/api/habits/:id", auth.AuthMiddleware(), habitController.UpdateHabit)
	router.GET("/api/habits", auth.AuthMiddleware(), habitController.GetAllHabits)
}
