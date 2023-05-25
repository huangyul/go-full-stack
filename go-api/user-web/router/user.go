package router

import (
	"github.com/gin-gonic/gin"
	"go-api/user-web/api"
	"go-api/user-web/middlewares"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("user") // .Use(middlewares.JWTAuth())
	zap.S().Info("配置用户相关url")
	{
		router.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		router.POST("pwd_login", api.PasswordLogin)
	}
}
