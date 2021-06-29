package auth

import (
	"github.com/gin-gonic/gin"
	"design-api/util"
	"log"
	"design-api/common/env"
	"design-api/service"
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

	sms := &service.SmsService{Len: 6}
	sms.SendSmsCode("15122801645")

	c.JSON(env.SUCCESS, gin.H{
		"code":         code,
		"message":      env.MsgFlags[code],
		"token_type":   "bearer",
		"access_token": token,
	})
}
