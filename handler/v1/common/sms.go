package common

import (
	//_ "github.com/patrickmn/go-cache"
	"design-api/service"
	"github.com/gin-gonic/gin"
	"design-api/common/env"
	"design-api/validator/auth"
	"design-api/common"
)

//验证码长度
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
			common.Format(c).SetStatus(env.ERROR).SetCode(env.RESPONSE_FAIL).SetMessage(env.MsgFlags[env.RESPONSE_FAIL]).JsonResponse()

			c.Abort()
			return
		}

		common.Format(c).JsonResponse()
	} else {
		common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
	}
}
