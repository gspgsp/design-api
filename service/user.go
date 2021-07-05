package service

import (
	"design-api/common/env"
	"design-api/common"
	"design-api/model"
)

type UserService struct {
	UserId string
}

/**
用户信息
 */
func (u *UserService) UserInfo() (int, interface{}) {
	user := models.User{}
	err := common.Db.Where("id = " + u.UserId).Select("id", "name", "nickname", "avatar", "mobile", "email", "password").Find(&user).Error
	if err != nil {
		return env.ACCOUNT_ERROR, user
	}

	return env.RESPONSE_SUCCESS, user
}
