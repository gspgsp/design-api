package service

import (
	"design-api/common"
	"design-api/common/env"
	mongo "design-api/common/log"
	"design-api/model"
	"design-api/util"
	"design-api/validator/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Register /**注册服务
func Register(c *gin.Context) (int, int64) {

	m := util.JsonParamParse(c)
	registerParam := &auth.RegisterParam{}
	registerParam.ParseParam(m)

	if code, _ := registerParam.ValidateParam(); code == env.RESPONSE_SUCCESS {

		var d mongo.SmsMongoInfo
		mongo.NewMgo("sms_code").GetOne(bson.M{"codeKey": registerParam.CodeKey}, &d)

		//取消使用go 自带的cache
		//smsCode, ok := common.Cache.Get(registerParam.CodeKey)
		if (d == mongo.SmsMongoInfo{}) {
			return env.SMS_CODE_KEY_INVALID, 0
		}

		if d.ExpireAt < time.Now().Unix() {
			return env.SMS_CODE_EXPIRE_ERROR, 0
		}

		if d.Code != registerParam.Code {
			return env.SMS_CODE_VERIFY_ERROR, 0
		}

		if d.Mobile != registerParam.Mobile {
			return env.SMS_CODE_INVALID_MOBILE, 0
		}

		//TODO::注册业务 要不要 DAO 层?
		password, _ := util.PasswordHash(registerParam.Password)
		uuid := fmt.Sprintf("%s", uuid2.NewV4())
		randStr := util.RandStr{}
		nickname := randStr.Generate(10)

		user := models.User{Mobile: registerParam.Mobile, Password: password, Uuid: uuid, Nickname: nickname, RegisterAt: time.Now().Unix(), RegisterIp: util.ClientIp(c), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
		common.Db.Create(&user)

		return env.RESPONSE_SUCCESS, *user.ID
	} else {
		return env.PARAM_REQUIRED, 0
	}
}
