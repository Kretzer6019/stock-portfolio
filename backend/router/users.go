package router

import (
	"dependencies/controllers/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userRoutes = gin.RoutesInfo{
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
