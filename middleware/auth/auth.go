package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
登录验证中间件
 */
func Auth() gin.HandlerFunc {

	a := func(c *gin.Context) {

		// 验证不通过直接返回
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()

		return

		//c.Next()
	}

	return a
}
