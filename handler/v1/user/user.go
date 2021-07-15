package user

import (
	"github.com/gin-gonic/gin"
	"design-api/service"
	"design-api/common/env"
	"design-api/common"
	"design-api/util"
)

/**
用户信息
 */
func UserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")

	user_id := util.Strval(userId)
	userService := service.UserService{UserId: user_id}
	code, user := userService.UserInfo()
	//
	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(user).JsonResponse()
}
