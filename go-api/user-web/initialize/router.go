package initialize

import (
	"github.com/gin-gonic/gin"
	"go-api/user-web/middlewares"
	router2 "go-api/user-web/router"
	"net/http"
)

func Routers() *gin.Engine {
	router := gin.Default()

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 配置跨域
	router.Use(middlewares.Cors())

	// 路由初始化
	ApiGroup := router.Group("/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)

	return router
}
