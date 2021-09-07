package common

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

//验证码长度
var codeLength = 6

/**
发送短信验证码
 */
func SendSms(c *gin.Context) {

	b, _ := c.GetRawData()
	var m map[string]interface{}
	_ = json.Unmarshal(b,&m)
	log.Printf("m is:%v", m)
	return


	//mobile, _ := c.GetPostForm("mobile")
	//smsParam := &auth.SmsParam{Mobile: mobile}
	//if code, _ := smsParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
	//	sms := &service.SmsService{Len: codeLength}
	//	codeKey, err := sms.SendSmsCode(mobile)
	//	if err != nil {
	//		common.Format(c).SetStatus(env.ERROR).SetCode(env.SMS_CODE_SEND_ERROR).SetMessage(env.MsgFlags[env.SMS_CODE_SEND_ERROR]).JsonResponse()
	//		c.Abort()
	//		return
	//	}
	//
	//	common.Format(c).SetData(map[string]interface{}{"code_key": codeKey}).JsonResponse()
	//} else {
	//	common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
	//}
}
