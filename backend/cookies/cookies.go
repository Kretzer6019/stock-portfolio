package cookies

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func SaveCookie(c *gin.Context, ID string, token string) error {
	s = securecookie.New([]byte(os.Getenv("HASH_KEY")), []byte(os.Getenv("BLOCK_KEY")))

	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodedData, err := s.Encode("Authorization", data)
	if err != nil {
		return err
	}

	c.SetCookie("Authorization", encodedData, 3600, "/", "localhost", true, true)

	return nil
}

func ReadCookies(c *gin.Context) (map[string]string, error) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	err = s.Decode("Authorization", cookie, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
