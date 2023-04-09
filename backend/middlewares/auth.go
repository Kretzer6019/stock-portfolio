package middlewares

import (
	"dependencies/auth"
	"dependencies/models/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := auth.ExtractUserID(c)
		if err != nil {
			log.Println("Error:", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		// Connect to db
		db := c.MustGet("db").(*gorm.DB)
		// Check if user exists
		_, err = users.SelectUser(int(userID), db)
		if err != nil {
			log.Println("Error:", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		c.Next()
	}
}
