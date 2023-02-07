package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Users struct {
	ID       int    `gorm:"id"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users Users
	err := db.First(&users, "email = ?", "vitorrkretzer@hotmail.com").Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, users)
}
