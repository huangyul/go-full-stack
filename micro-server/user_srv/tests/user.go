package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"micro-server/user_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

// Init 初始化一个服务
func Init() {
	var err error
	conn, err = grpc.Dial("47.106.214.127:3306", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 2,
	})
	if err != nil {
		fmt.Println("创建客户端失败")
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)

	}
}

func main() {
	Init()

	TestGetUserList()

	conn.Close()
}
