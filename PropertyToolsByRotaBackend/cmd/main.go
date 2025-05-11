package main

import (
	"PropertyToolsByRotaBackend/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set up routes
	router.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
