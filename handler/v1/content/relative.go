package content

import (
	"github.com/gin-gonic/gin"
	"design-api/validator/content"
	"design-api/common/env"
	"design-api/common"
	"design-api/service"
	"design-api/model"
)

/**
相关列表
 */
func Relative(c *gin.Context) {
	detailParam := content.DetailParam{c}
	if code, content := detailParam.ValidateParam(); code == env.RESPONSE_SUCCESS {

		code, contents := service.RelativeList(content.(models.Content).Uuid)
		if code != env.RESPONSE_SUCCESS {
			common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

			c.Abort()
			return
		}

		common.Format(c).SetData(contents).JsonResponse()
	} else {
		common.Format(c).SetStatus(env.ERROR).SetCode(env.ERROR).SetMessage(env.MsgFlags[env.ERROR]).JsonResponse()
	}
}
