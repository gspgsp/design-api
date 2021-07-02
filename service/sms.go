package service

import (
	"design-api/util"
	"log"
	"design-api/common"
	"time"
	_ "github.com/aliyun/alibaba-cloud-sdk-go"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	_ "gopkg.in/ini.v1"
	"design-api/config"
	"github.com/pkg/errors"
)

var codeKey = "key-"

type SmsService struct {
	Len int
}

/**
发送短信服务
 */
func (sms *SmsService) SendSmsCode(phone string) (string, error) {
	randInt := new(util.RandInt)

	code := randInt.Generate(sms.Len)
	//阿里云短信
	//err := aliSmsSend(phone, code)
	//if err != nil {
	//	return "", err
	//}

	codeKey += new(util.RandStr).Generate(10)
	common.Cache.Set(codeKey, code, 5*time.Minute)

	log.Println("code is:" + code + " codeKey is:" + codeKey)
	return codeKey, nil
}

/**
阿里短信
 */
func aliSmsSend(phone, code string) (error) {
	client, err := dysmsapi.NewClientWithAccessKey(config.Config.Sms.RegionId, config.Config.Sms.AccessKeyId, config.Config.Sms.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = config.Config.Sms.SignName
	request.TemplateCode = config.Config.Sms.TemplateCode
	request.TemplateParam = `{"code":"` + code + `","product":"` + config.Config.Sms.SignName + `"}`

	response, err := client.SendSms(request)
	if err != nil {
		log.Printf("发送验证码失败:%s", err.Error())
		return err
	}

	if response.Code == "OK" {
		return nil
	} else {
		return errors.New(response.Message)
	}
}
