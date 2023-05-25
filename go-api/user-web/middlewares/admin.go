package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-api/user-web/models"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		// 2的话才是管理员
		if currentUser.AuthorityId == 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "用户没有权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
