package auth

import (
	"github.com/gin-gonic/gin"
	"design-api/common/env"
	"design-api/util"
	"strings"
)

/**
登录验证中间件
 */
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(env.NO_AUTHED, gin.H{
				"code":    env.NO_AUTHED,
				"message": env.MsgFlags[env.NO_AUTHED],
				"data":    nil,
			})

			c.Abort()
			return
		} else {
			tokenSlice := strings.Split(authorization, " ")
			claim, code := util.ParseToken(tokenSlice[1])
			if code != env.SUCCESS {
				c.JSON(env.ERROR, gin.H{
					"code":    code,
					"message": env.MsgFlags[code],
				})

				c.Abort()
				return
			}

			//将解析后的有效载荷claim重新写入gin.Context引用对象中
			c.Set("userId", claim.UserId)
			c.Next()
		}
	}
}
