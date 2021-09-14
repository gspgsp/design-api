package auth

import (
	"design-api/common"
	"design-api/common/env"
	"design-api/service"
	"github.com/gin-gonic/gin"
)

/**
忘记密码
*/
func Forget(c *gin.Context) {
	code := service.Forget(c)
	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetStatus(env.SUCCESS).SetCode(code).JsonResponse()
}
