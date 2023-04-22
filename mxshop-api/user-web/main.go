package main

import "mxshop-api/user-web/initialize"

func main() {

	// 1. 初始化路由
	router := initialize.Routers()

	router.Run(":5555")
}
