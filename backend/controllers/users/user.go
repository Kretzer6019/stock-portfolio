package users

import (
	"dependencies/models/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	user, err := users.SelectUser(1, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
