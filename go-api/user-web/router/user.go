package router

import (
	"github.com/gin-gonic/gin"
	"go-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	Router.Group("user")
	{
		Router.GET("list", api.GetUserList)
	}
}
