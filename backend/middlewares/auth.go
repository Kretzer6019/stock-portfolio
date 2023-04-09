package middlewares

import (
	"dependencies/auth"
	"dependencies/models/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := auth.ExtractUserID(c)
		if err != nil {
			log.Println("Error:", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		// Check if user exists
		_, err = users.SelectUser(int(userID))
		if err != nil {
			log.Println("Error:", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		c.Next()
	}
}
