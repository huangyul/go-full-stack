package router

import (
	"github.com/gin-gonic/gin"
	"go-api/user-web/api"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("user")
	zap.S().Info("配置用户相关url")
	{
		router.GET("list", api.GetUserList)
		router.POST("pwd_login", api.PasswordLogin)
	}
}
