package auth

import (
	"github.com/gin-gonic/gin"
	"design-api/util"
	"log"
	"design-api/common/env"
)

func Register(c *gin.Context) {

	token, code := util.GenerateToken("101", "guo")

	claim, _ := c.Get("claim")
	log.Printf("claim is:%v", claim)

	if code != env.SUCCESS {
		c.JSON(env.ERROR, gin.H{
			"code":    code,
			"message": env.MsgFlags[code],
		})

		c.Abort()
		return
	}

	c.JSON(env.SUCCESS, gin.H{
		"code":         env.RESPONSE_SUCCESS,
		"message":      env.MsgFlags[env.RESPONSE_SUCCESS],
		"token_type":   "Bearer",
		"access_token": token,
	})
}
