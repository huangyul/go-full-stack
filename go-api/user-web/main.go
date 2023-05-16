package main

import (
	"fmt"
	"go-api/user-web/global"
	"go-api/user-web/initialize"
	"go.uber.org/zap"
)

func main() {
	// 定义一个全局logger
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()

	// 初始化router
	router := initialize.Routers()

	zap.S().Infof("启动服务器，端口：%d", global.ServerConfig.Port)

	err := router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
	if err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
