package service

import (
	"design-api/util"
	"log"
)

type SmsService struct {
	Len int
}

/**
发送短信服务
 */
func (sms *SmsService) SendSmsCode(phone string) bool {
	randInt := new(util.RandInt)

	code := randInt.Generate(sms.Len)
	if len(code) > 0 {
		log.Println("code is:" + code)
	}

	return true
}
