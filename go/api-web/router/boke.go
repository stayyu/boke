package router

import (
	"boke/go/api-web/api/blog"
	"boke/go/api-web/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitBokRouter(router *gin.RouterGroup) {
	BokeRouter := router.Group("/bok")
	fmt.Println("配置博客相关url")
	{
		BokeRouter.GET("", blog.List)
		BokeRouter.POST("/create", middlewares.JWTAuth(), middlewares.IsAdminAuth(), blog.New)
		BokeRouter.GET("/detail", blog.Detail)
		BokeRouter.DELETE("/del/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), blog.Delete)
		BokeRouter.PUT("/updata", middlewares.JWTAuth(), middlewares.IsAdminAuth(), blog.Update)
		BokeRouter.POST("/upload", handleUpload)
		BokeRouter.GET("/files/:filename", handleFile)
		BokeRouter.Static("img", "./test/img")
	}
}
