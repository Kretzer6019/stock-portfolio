package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}

	return cors.New(config)
}
