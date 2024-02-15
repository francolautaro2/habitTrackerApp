package main

import (
	"fmt"
	"habitTrackerApi/routes"
	"habitTrackerApi/services/database"
	"habitTrackerApi/services/users"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		fmt.Println("error to connect database, ", err)
	}
	// Automigrate the users models
	db.AutoMigrate(&users.UserClient{})

	// Create instance of User Repository
	userRepository := database.NewDatabaseUserRepository(db)

	// Create instance of User Controllers
	userController := &users.UserController{
		UserRepository: userRepository,
	}

	// Set router engine
	router := gin.Default()

	// Set routes of api
	routes.RunRoutes(router, userController)

	// Run the server
	router.Run(":8080")

}
