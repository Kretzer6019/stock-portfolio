package middlewares

import (
	"dependencies/cookies"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := cookies.ReadCookies(c)
		if err != nil {
			log.Println("Error:", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		c.Next()
	}
}
