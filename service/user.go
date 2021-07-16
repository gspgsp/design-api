package service

import (
	"design-api/common"
	"design-api/common/env"
	"design-api/model"
	"log"
	"strconv"
)

type UserService struct {
	UserId   string
	Name     string
	NickName string
	Page     int
}

// UserInfo /**用户信息
func (u *UserService) UserInfo() (int, interface{}) {
	user := models.User{}
	err := common.Db.Where("id = "+u.UserId).Select("id", "name", "nickname", "avatar", "mobile", "email").Find(&user).Error
	if err != nil {
		return env.ACCOUNT_ERROR, user
	}

	return env.RESPONSE_SUCCESS, user
}

// UpdateUserInfo /**更新用户信息
func (u *UserService) UpdateUserInfo(operateType int) int {
	param := make(map[string]interface{})
	if operateType == 1 {
		param["name"] = u.Name
	} else if operateType == 2 {
		param["nickname"] = u.NickName
	}

	log.Printf("id is:%v\n", u)

	if err := common.Db.Table("q_users").Where("id = " + u.UserId).Updates(param).Error; err != nil {
		return env.RESPONSE_FAIL
	}

	return env.RESPONSE_SUCCESS
}

// GetUserQuote /**用户报价信息
func (u *UserService) GetUserQuote() (int, interface{}) {
	sql := "select id, name, phone, address, size, status from q_quotes where user_id = " + u.UserId + " limit " + strconv.Itoa((u.Page-1)*10) + ", 10"

	quotes := make([]models.Quote, 0)
	if err := common.Db.Raw(sql).Scan(&quotes).Error; err != nil {
		return env.RESPONSE_FAIL, nil
	}

	return env.RESPONSE_SUCCESS, quotes
}

// GetUserFavor /**用户收藏信息
func (u *UserService) GetUserFavor() (int, interface{}) {
	sql := "select c.id, c.uuid, c.title, c.sub_title, c.size, c.mb_cover_picture from q_contents as c where c.id in(select design_id from q_user_favors where user_id = " + u.UserId + ") limit " + strconv.Itoa((u.Page-1)*10) + ", 10"

	contents := make([]models.Content, 0)
	if err := common.Db.Raw(sql).Scan(&contents).Error; err != nil {
		return env.RESPONSE_FAIL, nil
	}

	return env.RESPONSE_SUCCESS, contents
}

// GetUserStar /**用户关注信息
func (u *UserService) GetUserStar() (int, interface{}) {
	sql := "select d.id, d.uuid, d.nick_name, d.photo from q_designers as d where d.id in(select designer_id from q_fans where user_id = " + u.UserId + ") limit " + strconv.Itoa((u.Page-1)*10) + ", 10"

	designers := make([]models.Designer, 0)
	if err := common.Db.Raw(sql).Scan(&designers).Error; err != nil {
		return env.RESPONSE_FAIL, nil
	}

	return env.RESPONSE_SUCCESS, designers
}
