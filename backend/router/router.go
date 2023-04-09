package router

import (
	"dependencies/db"
	"dependencies/middlewares"

	"github.com/gin-gonic/gin"
)

func Generate() *gin.Engine {
	r := gin.Default()
	db.Init()
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.DatabaseMiddleware(db.DB))
	return configRouter(r)
}

func configRouter(r *gin.Engine) *gin.Engine {
	routes := userRoutes

	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	apiRoutes := apiRoutes
	for _, route := range apiRoutes {
		r.Handle(route.Method, route.Path, middlewares.Auth(), route.HandlerFunc)
	}

	return r
}
