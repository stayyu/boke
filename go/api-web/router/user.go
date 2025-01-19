package router

import (
	"boke/go/api-web/api"
	"boke/go/api-web/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{

		UserRouter.POST("/pwd_login", api.PassWordLogin)
		UserRouter.POST("/register", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.Register)

	}
}
