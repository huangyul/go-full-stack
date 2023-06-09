package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"go_srvs/user_srv/global"
	"go_srvs/user_srv/handler"
	"go_srvs/user_srv/initialize"
	"go_srvs/user_srv/proto"
	"go_srvs/user_srv/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("Port", 50051, "端口号")
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()
	zap.S().Info("ip: ", *IP)
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}

	zap.S().Info("port:", *Port)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic(err)
	}

	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 注册服务
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "192.168.121.136:8500"

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}

	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.0.112:%d", *Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.Address = "192.168.0.112"
	registration.Port = *Port
	registration.ID = global.ServerConfig.Name
	registration.Tags = []string{"srv", "user-srv"}
	registration.Check = check

	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
