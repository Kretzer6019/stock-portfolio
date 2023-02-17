package cookies

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func SaveCookie(c *gin.Context, ID, token string) error {
	s = securecookie.New([]byte(os.Getenv("HASH_KEY")), []byte(os.Getenv("BLOCK_KEY")))

	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodedData, err := s.Encode("auth", data)
	if err != nil {
		return err
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "auth",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func ReadCookies(c *gin.Context) (map[string]string, error) {
	cookie, err := c.Cookie("auth")
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	if err = s.Decode("auth", cookie, &data); err != nil {
		return nil, err
	}

	return data, nil
}
