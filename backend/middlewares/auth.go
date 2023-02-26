package middlewares

import (
	"dependencies/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateToken(c)
		if err != nil {
			log.Println("Error:", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		c.Next()
	}
}
