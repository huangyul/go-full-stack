package main

import (
	"fmt"
	"go-api/user-web/initialize"
	"go.uber.org/zap"
)

func main() {
	// 定义一个全局logger
	initialize.InitLogger()

	// 初始化router
	router := initialize.Routers()
	
	port := 8021

	zap.S().Infof("启动服务器，端口：%d", port)

	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
