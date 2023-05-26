package initialize

import (
	"github.com/gin-gonic/gin"
	"go-api/user-web/middlewares"
	router2 "go-api/user-web/router"
)

func Routers() *gin.Engine {
	router := gin.Default()

	// 配置跨域
	router.Use(middlewares.Cors())

	// 路由初始化
	ApiGroup := router.Group("/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)

	return router
}
