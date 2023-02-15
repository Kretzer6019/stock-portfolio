package users

import (
	"dependencies/models/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

type CreteAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateAccount(c *gin.Context) {
	var request CreteAccountRequest
	c.ShouldBindJSON(&request)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	id, err := users.InsertUser(request.Email, string(hashedPassword), db)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, id)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var request LoginRequest
	c.ShouldBindJSON(&request)

	db := c.MustGet("db").(*gorm.DB)

	user, err := users.SelectUserByEmail(request.Email, db)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, "Logado com Sucesso!")
}
