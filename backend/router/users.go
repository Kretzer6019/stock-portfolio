package router

import (
	"dependencies/controllers/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userRoutes = []gin.RouteInfo{
	{
		Method:      http.MethodGet,
		Path:        "/",
		HandlerFunc: users.GetUser,
	},
	{
		Method:      http.MethodPost,
		Path:        "/criar-conta",
		HandlerFunc: users.CreateAccount,
	},
	{
		Method:      http.MethodPost,
		Path:        "/login",
		HandlerFunc: users.Login,
	},
}
