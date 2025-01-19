package main

import (
	"boke/go/api-server/handler"
	"boke/go/api-server/initialize"
	"boke/go/api-server/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//IP := flag.String("ip", "127.0.0.1", "IP地址")
	//Port := flag.Int("port", 0, "端口号")
	initialize.InitDB()

	server := grpc.NewServer()
	proto.RegisterBlogServer(server, &handler.BlogServer{})
	proto.RegisterUserServer(server, &handler.UserServer{})

	//lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	lis, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		panic("failed to listen" + err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic("failed to server" + err.Error())
	}

}
