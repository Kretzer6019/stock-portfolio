package middlewares

import (
	"dependencies/cookies"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := cookies.ReadCookies(c)
		if err != nil {
			fmt.Println("AQUII")
		}
		c.Next()
	}
}
