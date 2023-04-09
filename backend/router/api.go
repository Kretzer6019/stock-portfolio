package router

import (
	"dependencies/controllers/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

var apiRoutes = gin.RoutesInfo{
	{
		Method:      http.MethodGet,
		Path:        "/",
		HandlerFunc: users.GetUser,
	},
	{
		Method:      http.MethodGet,
		Path:        "/logout",
		HandlerFunc: users.Logout,
	},
}
