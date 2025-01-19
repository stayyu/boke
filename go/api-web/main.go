package main

import (
	"boke/go/api-web/global"
	"boke/go/api-web/initialize"
	"boke/go/api-web/proto"
	"context"
	"fmt"
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

func TestCreatecategory() {
	req, err := global.BlogSrvClient.CreateCategory(context.Background(), &proto.CreateCategoryInfo{
		Name: "科学1",
		Id:   0003,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(req.Susee)
}
func TestCreateUser() {
	rsp, err := global.UsersrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NiceName: "admin",
		PassWord: "admin123",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)

}

func TestCreateblog() {
	rsp, err := global.BlogSrvClient.CreateBlog(context.Background(), &proto.CreateBlogInfo{
		Title:      "标题1",
		CategoryId: 2,
		Context:    "ssssssssssssssssssss",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Susee)
}

func TestFindCategory() {
	detail, err := global.BlogSrvClient.GetCategoryDetail(context.Background(), &proto.CategoryInfoRequest{
		Id: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(detail.Data)
}

func TestFindblog() {
	detail, err := global.BlogSrvClient.GetBlogList(context.Background(), &proto.BlogFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(detail.Data)
}

func main() {
	Router := initialize.Routers()
	if err := Router.Run("127.0.0.1:8081"); err != nil {
		panic(err)
	}
	//Init()
	//TestFindCategory()
	//TestCreateblog()
	//TestFindblog()
}
