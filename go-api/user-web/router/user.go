package router

import (
	"github.com/gin-gonic/gin"
	"go-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("user")
	{
		router.GET("list", api.GetUserList)
	}
}
