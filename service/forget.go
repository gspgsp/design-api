package service

import (
	"design-api/common"
	"design-api/common/env"
	mongo "design-api/common/log"
	models "design-api/model"
	"design-api/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

/**
忘记密码
*/
func Forget(c *gin.Context) int {
	m := util.JsonParamParse(c)

	password, ok := m["password"]
	if !ok {
		return env.PARAM_REQUIRED
	}

	codeKey, ok := m["code_key"]
	if !ok {
		return env.PARAM_REQUIRED
	}

	var d mongo.SmsMongoInfo
	mongo.NewMgo("sms_code").GetOne(bson.M{"codeKey": codeKey}, &d)

	var mobile string
	if len(d.Mobile) > 0 {
		log.Println("mobile is:" + mobile)
		password, _ := util.PasswordHash(password)
		err := common.Db.Model(&models.User{}).Where("mobile =?", mobile).Update("password", password).Error

		if err != nil {
			return env.RESPONSE_FAIL
		}

		return env.RESPONSE_SUCCESS
	}

	return env.PARAM_REQUIRED
}
