package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set cors policy
	r.Use(corsMiddleware())

	// Add your routes here
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, nil)
	})

	r.Run()
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	return cors.New(config)
}
