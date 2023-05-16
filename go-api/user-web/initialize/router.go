package initialize

import (
	"github.com/gin-gonic/gin"
	router2 "go-api/user-web/router"
)

func Routers() *gin.Engine {
	router := gin.Default()

	// 路由初始化
	ApiGroup := router.Group("/v1")
	router2.InitUserRouter(ApiGroup)

	return router
}
