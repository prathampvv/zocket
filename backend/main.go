package main

import (
	"backend/config"
	"backend/database"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	database.ConnectDB()

	// Create a new Gin router
	r := gin.Default()

	// Setup routes
	routes.AuthRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
