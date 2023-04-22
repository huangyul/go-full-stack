package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/user-web/initialize"
)

func main() {
	// 初始化日志
	initialize.InitLogger()

	// 2. 初始化路由
	router := initialize.Routers()
	port := 8021

	zap.S().Infof("启动服务器，端口：%d", port)

	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
