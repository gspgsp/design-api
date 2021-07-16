package user

import (
	"design-api/common"
	"design-api/common/env"
	"design-api/service"
	"design-api/util"
	"design-api/validator/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserInfo /**用户信息
func UserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")

	userService := service.UserService{UserId: util.Strval(userId)}
	code, user := userService.UserInfo()
	//
	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(user).JsonResponse()
}

// UpdateUserInfo /**更新用户信息
func UpdateUserInfo(c *gin.Context) {
	updateParam := user.UpdateParam{C: c}

	if code, res := updateParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
		u := res.(*user.UpdateParam)

		userId, _ := c.Get("userId")
		userService := &service.UserService{UserId: util.Strval(userId), Name: u.Name, NickName: u.NickName}
		code = userService.UpdateUserInfo(u.OperateType)

		if code == env.RESPONSE_SUCCESS {
			common.Format(c).JsonResponse()
		} else {
			common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()
		}
	} else {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()
	}
}

// UserQuotes /**用户报价信息
func UserQuotes(c *gin.Context) {
	userId, _ := c.Get("userId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	userService := &service.UserService{UserId: util.Strval(userId), Page: page}
	code, quotes := userService.GetUserQuote()
	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(quotes).JsonResponse()
}

// UserFavors /**用户收藏信息
func UserFavors(c *gin.Context) {
	userId, _ := c.Get("userId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	userService := &service.UserService{UserId: util.Strval(userId), Page: page}
	code, favors := userService.GetUserFavor()
	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(favors).JsonResponse()
}

// UserStars /**用户关注信息
func UserStars(c *gin.Context) {
	userId, _ := c.Get("userId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	userService := &service.UserService{UserId: util.Strval(userId), Page: page}
	code, designers := userService.GetUserStar()
	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(designers).JsonResponse()
}
