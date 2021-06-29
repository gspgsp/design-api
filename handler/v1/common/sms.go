package common

import (
	//_ "github.com/patrickmn/go-cache"
	"design-api/service"
	"github.com/gin-gonic/gin"
	"design-api/common/env"
	"design-api/validator/auth"
)

var codeLength = 6
/**
发送短信验证码
 */
func SendSms(c *gin.Context) {
	mobile, _ := c.GetPostForm("mobile")

	smsParam := &auth.SmsParam{Mobile: mobile}
	if code := smsParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
		sms := &service.SmsService{Len: codeLength}
		if sms.SendSmsCode(mobile) == false {
			c.JSON(env.ERROR, gin.H{
				"code":    env.RESPONSE_FAIL,
				"message": env.MsgFlags[env.RESPONSE_FAIL],
			})

			c.Abort()
			return
		}

		c.JSON(env.SUCCESS, gin.H{
			"code":    env.RESPONSE_SUCCESS,
			"message": env.MsgFlags[env.RESPONSE_SUCCESS],
		})
	} else {
		c.JSON(env.ERROR, gin.H{
			"code":    env.PARAM_REQUIRED,
			"message": env.MsgFlags[env.PARAM_REQUIRED],
		})
	}
}
