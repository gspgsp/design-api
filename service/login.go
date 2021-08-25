package service

import (
	"design-api/common"
	"design-api/common/env"
	"design-api/model"
	"design-api/util"
	"design-api/validator/auth"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) (int, models.User) {
	//postman可以传form，也可以传json，对于php来讲 框架都可以解析，但是gin的话，必须区别对待.直接form是获取不到参数的
	//c.Request.ParseForm()
	//values := c.Request.PostForm
	var user models.User
	//
	method := c.Request.Method
	if method == "POST" {
		values, err := parseJsonParams(c.Request.Body)
		if err != nil {
			return env.PARAM_REQUIRED, user
		}

		loginParam := &auth.LoginParam{}
		loginParam.ParseParam(values)

		if code, _ := loginParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
			err := common.Db.Where("mobile = "+loginParam.Mobile).Select("id", "name", "nickname", "avatar", "mobile", "email", "password").Find(&user).Error
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

	return env.INVALID_METHOD, user
}
