package category

import (
	"github.com/gin-gonic/gin"
	"design-api/common/env"
	"design-api/service"
	"design-api/common"
)

/**
风格列表
 */
func Style(c *gin.Context) {
	code, styles := service.StyleList()

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(styles).JsonResponse()
}
