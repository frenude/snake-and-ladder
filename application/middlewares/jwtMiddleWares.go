package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snake-and-ladder/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
				"data":   nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
