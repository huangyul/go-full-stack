package main

import (
	"context"
	"fmt"
	"go_srvs/user_srv/proto"
	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {

	res, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range res.Data {
		fmt.Printf("%v \r\n", user)
		fmt.Printf("用户名：%s, 手机号：%s \r\n", user.NickName, user.Mobile)
	}
}

func main() {
	Init()
	TestGetUserList()
	conn.Close()
}
