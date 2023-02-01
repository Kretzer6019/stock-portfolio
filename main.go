package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Add your routes here
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Gin",
		})
	})

	router.Run()
}
