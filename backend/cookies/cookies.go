package cookies

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func SaveCookie(c *gin.Context, ID, token string) error {
	if os.Getenv("FRONTEND_URL") != "" {
		c.SetSameSite(http.SameSiteNoneMode)
	}

	s = securecookie.New([]byte(os.Getenv("HASH_KEY")), []byte(os.Getenv("BLOCK_KEY")))

	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodedData, err := s.Encode("user", data)
	if err != nil {
		return err
	}

	c.SetCookie("user", encodedData, 3600, "/", "", true, true)

	return nil
}

func ReadCookies(c *gin.Context) (map[string]string, error) {
	cookie, err := c.Cookie("user")
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	if err = s.Decode("user", cookie, &data); err != nil {
		return nil, err
	}

	return data, nil
}
