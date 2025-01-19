package initialize

import (
	"boke/go/api-web/global"
	"boke/go/api-web/middlewares"
	"boke/go/api-web/proto"
	"boke/go/api-web/router"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	global.UsersrvClient = proto.NewUserClient(conn)
	global.BlogSrvClient = proto.NewBlogClient(conn)
}
func Routers() *gin.Engine {
	Init()
	Router := gin.Default()
	//跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/bk")
	router.InitBokRouter(ApiGroup)
	router.InitCateRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return Router
}
