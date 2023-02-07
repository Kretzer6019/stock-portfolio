package router

import (
	"dependencies/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Generate() *gin.Engine {
	r := gin.Default()
	DB := db.Init()
	r.Use(db.DatabaseMiddleware(DB))
	r.Use(corsMiddleware())
	return configRouter(r)
}

func configRouter(r *gin.Engine) *gin.Engine {
	routes := userRoutes

	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	return r
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	return cors.New(config)
}
