package auth

import (
	"dependencies/cookies"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const tokenExpireDays = 7

func CreateToken(userID int) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(tokenExpireDays * time.Hour * 24).Unix()
	permissions["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(os.Getenv("AUTH_JWT_SECRET_KEY")))
}

func ValidateToken(c *gin.Context) error {
	cookie, err := cookies.ReadCookies(c)
	if err != nil {
		return err
	}

	token, err := jwt.Parse(cookie["token"], returnVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("wrong method for verification key")
	}

	return []byte(os.Getenv("AUTH_JWT_SECRET_KEY")), nil
}

func ExtractUserID(c *gin.Context) (int64, error) {
	cookie, err := cookies.ReadCookies(c)
	if err != nil {
		return 0, err
	}

	token, err := jwt.Parse(cookie["token"], returnVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("invalid token")
}
