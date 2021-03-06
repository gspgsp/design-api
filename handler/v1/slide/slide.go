package slide

import (
	"github.com/gin-gonic/gin"
	"design-api/service"
	"design-api/common/env"
	"design-api/common"
)

/**
幻灯片
 */
func Slide(c *gin.Context) {
	code, slides := service.SlideList()

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(slides).JsonResponse()
}
