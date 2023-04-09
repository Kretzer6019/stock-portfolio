package users

import (
	"dependencies/auth"
	"dependencies/cookies"
	"dependencies/models/users"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	/* db := c.MustGet("db").(*gorm.DB)

	user, err := users.SelectUser(0, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} */

	c.JSON(http.StatusOK, nil)
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

type LoginResponse struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
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

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = cookies.SaveCookie(c, fmt.Sprintf("%d", user.ID), token)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func Logout(c *gin.Context) {
	// Extract user ID from cookie
	userID, err := auth.ExtractUserID(c)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// Connect to db
	db := c.MustGet("db").(*gorm.DB)
	// Check if user exists
	_, err = users.SelectUser(int(userID), db)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// Clean cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "Authorization",
		Value:   "",
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})

	c.JSON(http.StatusUnprocessableEntity, nil)
}
