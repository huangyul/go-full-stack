package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/v1")

	// 初始化user的路由
	router.InitUserRouter(ApiGroup)

	return Router
}
