package service

import (
	"context"
	mongo "design-api/common/log"
	"design-api/config"
	"design-api/util"
	_ "github.com/aliyun/alibaba-cloud-sdk-go"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	_ "gopkg.in/ini.v1"
	"time"
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
	err := aliSmsSend(phone, code)
	if err != nil {
		return "", err
	}

	codeKey += new(util.RandStr).Generate(10)
	//取消使用go 自带的cache
	//common.Cache.Set(codeKey, code, 5*time.Minute)

	mgo := mongo.NewMgo("sms_code")

	var d mongo.SmsMongoInfo
	mgo.GetOne(bson.M{"mobile": phone}, &d)
	expireAt := time.Now().Unix() + 120

	if len(d.Mobile) > 0 {
		mongo.NewMgo("sms_code").GetCollection().UpdateOne(context.Background(), bson.M{"mobile": phone}, bson.M{"$set": bson.M{"code": code, "codeKey": codeKey, "expireAt": expireAt}})
	} else {
		mgo.InsertOne(bson.D{{"mobile", phone}, {"code", code}, {"codeKey", codeKey}, {"expireAt", expireAt}})
	}

	return codeKey, nil
}

/**
阿里短信
*/
func aliSmsSend(phone, code string) error {
	client, err := dysmsapi.NewClientWithAccessKey(config.Config.Sms.RegionId, config.Config.Sms.AccessKeyId, config.Config.Sms.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = config.Config.Sms.SignName
	request.TemplateCode = config.Config.Sms.TemplateCode
	request.TemplateParam = `{"code":"` + code + `","product":"` + config.Config.Sms.SignName + `"}`

	response, err := client.SendSms(request)
	if err != nil {
		mgo := mongo.NewMgo("sms_log")
		mgo.InsertOne(bson.D{{"mobile", phone}, {"code", code}, {"sendErr", err.Error()}})
		return err
	}

	if response.Code == "OK" {
		return nil
	} else {
		return errors.New(response.Message)
	}
}
