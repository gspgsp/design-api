package service

import (
	"design-api/validator/auth"
	"design-api/common/env"
	"design-api/common"
	"design-api/model"
	"design-api/util"
	uuid2 "github.com/satori/go.uuid"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) (int, int64) {
	c.Request.ParseForm()
	values := c.Request.Form

	registerParam := &auth.RegisterParam{}
	registerParam.ParseParam(values)

	if code := registerParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
		smsCode, ok := common.Cache.Get(registerParam.CodeKey)
		if !ok {
			return env.SMS_CODE_EXPIRE_ERROR, 0
		}

		if smsCode != registerParam.Code {
			return env.SMS_CODE_VERIFY_ERROR, 0
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
