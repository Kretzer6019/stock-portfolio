package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	// Add your routes here
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, nil)
	})

	r.Run()
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the appropriate headers to allow cross-origin requests
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Continue with the request if it's not a preflight request
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(204)
		}
	}
}
