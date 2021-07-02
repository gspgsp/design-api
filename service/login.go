package service

import (
	"github.com/gin-gonic/gin"
	"design-api/validator/auth"
	"design-api/common/env"
	"design-api/common"
	"design-api/model"
	"design-api/util"
)

func Login(c *gin.Context) (int, models.User) {

	c.Request.ParseForm()
	values := c.Request.Form

	loginParam := &auth.LoginParam{}
	loginParam.ParseParam(values)
	var user models.User

	if code := loginParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
		err := common.Db.Where("mobile = " + loginParam.Mobile).Select("id", "name", "nickname", "avatar", "mobile", "email", "password").Find(&user).Error
		if err != nil {
			return env.ACCOUNT_ERROR, user
		}

		err = util.PasswordCheck([]byte(user.Password), []byte(loginParam.Password))
		if err != nil {
			return env.ACCOUNT_ERROR, user
		}

		return env.RESPONSE_SUCCESS, user
	} else {
		return env.PARAM_REQUIRED, user
	}
}
