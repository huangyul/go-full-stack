package main

import (
	"go_srvs/user_srv/proto"
	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

//func TestGetUserList(t *testing.T) {
//
//	res, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
//		Pn:    1,
//		PSize: 3,
//	})
//	if err != nil {
//		panic(err)
//	}
//	for _, user := range res.Data {
//		fmt.Println(user.Mobile, user.NickName)
//
//	}
//}

func main() {
	Init()
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
}
