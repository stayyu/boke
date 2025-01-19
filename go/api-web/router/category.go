package router

import (
	"boke/go/api-web/api/cate"
	"boke/go/api-web/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitCateRouter(router *gin.RouterGroup) {
	CateRouter := router.Group("/cate")
	fmt.Println("配置分类相关url")
	{
		CateRouter.GET("", cate.DetailList)
		CateRouter.POST("/create", middlewares.JWTAuth(), middlewares.IsAdminAuth(), cate.New)
		CateRouter.DELETE("/del/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), cate.Delete)
		CateRouter.PUT("/update", middlewares.JWTAuth(), middlewares.IsAdminAuth(), cate.Update)
	}
}
