package main

import (
	"PropertyToolsByRotaBackend/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.Default()

	// Set up routes
	router.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
