package common

import (
	"design-api/common"
	"design-api/common/env"
	"design-api/service"
	"design-api/util"
	"design-api/validator/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

//验证码长度
var (
	statusCode int
	codeKey    string
	codeLength = 6
)

// SendSms /** 发送短信验证码
func SendSms(c *gin.Context) {
	m := util.JsonParamParse(c)
	sendType, ok := m["send_type"]
	if !ok {
		common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
		return
	}

	mobile, ok := m["mobile"]
	if !ok {
		common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
		return
	}

	if sendType == "login" {
		//TODO::send
		codeKey, statusCode = sendSms(mobile)
	} else if sendType == "forget" {
		if forgetStep, ok := m["forget_step"]; ok {
			if forgetStep == "one" {
				if captchaKey, ok := m["captcha_key"]; ok {
					s, _ := common.Cache.Get(captchaKey)
					if captchaCode, ok := m["captcha_code"]; ok {
						sCode, _ := s.(string)
						if strings.ToLower(captchaCode) == strings.ToLower(sCode) {
							//TODO::send
							codeKey, statusCode = sendSms(mobile)
						} else {
							common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
							return
						}
					} else {
						common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
						return
					}
				} else {
					common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
					return
				}
			} else if forgetStep == "two" {
				//TODO::send
				codeKey, statusCode = sendSms(mobile)
			} else {
				common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
				return
			}
		} else {
			common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
			return
		}
	}

	if statusCode != env.SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[statusCode]).JsonResponse()
		return
	}

	common.Format(c).SetData(map[string]interface{}{"code_key": codeKey}).JsonResponse()
}

func sendSms(mobile string) (string, int) {
	smsParam := &auth.SmsParam{Mobile: mobile}
	if code, _ := smsParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
		sms := &service.SmsService{Len: codeLength}
		codeKey, err := sms.SendSmsCode(mobile)
		if err != nil {
			return "", env.SMS_CODE_SEND_ERROR
		}

		return codeKey, env.SUCCESS
	} else {
		return "", env.PARAM_REQUIRED
	}
}
