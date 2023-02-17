package middlewares

import (
	"dependencies/cookies"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := cookies.ReadCookies(c)
		fmt.Println(data, err)
		c.Next()
	}
}
