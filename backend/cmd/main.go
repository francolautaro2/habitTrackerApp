package main

import (
	"fmt"
	"habitTrackerApi/routes"
	"habitTrackerApi/services/auth"
	"habitTrackerApi/services/database"
	"habitTrackerApi/services/habits"
	"habitTrackerApi/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		fmt.Println("error to connect database, ", err)
	}

	// Migrar modelos en la base de datos
	if err := database.Migrate(db); err != nil {
		fmt.Println("error migrating models: ", err)
		return
	}

	// Create instance of User Repository
	userRepository := database.NewDatabaseUserRepository(db)

	// Create instance of User Controllers
	userController := &users.UserController{
		UserRepository: userRepository,
	}

	// Create instance of Auth Controller
	authController := auth.NewAuthController(userRepository)

	// Create habit Repository instance for database
	habitRepository := database.NewDatabaseHabitRepository(db)

	// Create habit controller instance
	habitController := &habits.HabitController{
		HabitRepository: habitRepository,
	}

	// Set router engine
	router := gin.Default()

	// Set headers CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Set routes of api
	routes.RunRoutes(router, userController, authController, habitController)

	// Run the server
	router.Run(":8080")
}
