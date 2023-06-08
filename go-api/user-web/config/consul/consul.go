package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// Register 注册服务
func Register(name string, id string, tags []string, address string, port int) {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", "192.168.121.136", 8500)

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// 配置健康检查
	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.0.112:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "5s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.Port = port
	registration.Address = address
	registration.ID = id
	registration.Tags = tags
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

// GetAllService 发现所有服务
func GetAllService() {
	config := api.DefaultConfig()
	config.Address = "http://192.168.121.136:8500"

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for name, _ := range services {
		fmt.Println(name)
	}
}

// FilterService 获取特定的服务
func FilterService() {
	config := api.DefaultConfig()
	config.Address = "http://192.168.121.136:8500"

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}
	for name, _ := range services {
		fmt.Println(name)
	}
}

func main() {
	//Register("user-web", "user-web", []string{"web"}, "192.168.0.112", 8021)
	GetAllService()
	FilterService()
}
