package users

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context) {
	c.JSON(200, "Hello World")
}