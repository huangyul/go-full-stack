package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go-api/user-web/global"
	"go-api/user-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	consuleInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consuleInfo.Host, consuleInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		fmt.Println(err)

		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	// 调用接口
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitSrvConn2() {
	// 从注册中心获取用户信息
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}

	consulData, err := consulClient.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}

	for _, v := range consulData {
		userSrvHost = v.Address
		userSrvPort = v.Port
		break
	}

	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接失败", "msg", err.Error())
	}
	// 调用接口
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
