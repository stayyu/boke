package middlewares

import (
	"boke/go/api-web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		customUser := claims.(*models.CustomClaims)
		if customUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
		}
		return
	}
}
