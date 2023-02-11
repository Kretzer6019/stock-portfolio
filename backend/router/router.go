package router

import (
	"dependencies/db"
	"dependencies/middlewares"

	"github.com/gin-gonic/gin"
)

func Generate() *gin.Engine {
	r := gin.Default()
	database := db.Init()
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.DatabaseMiddleware(database))
	return configRouter(r)
}

func configRouter(r *gin.Engine) *gin.Engine {
	routes := userRoutes

	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	return r
}
