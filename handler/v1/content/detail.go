package content

import (
	"github.com/gin-gonic/gin"
	"design-api/common/env"
	"design-api/model"
	"design-api/validator/content"
	"design-api/service"
	"design-api/common"
)

/**
内容详情
 */
func Detail(c *gin.Context) {

	detailParam := content.DetailParam{c}
	if code, content := detailParam.ValidateParam(); code == env.RESPONSE_SUCCESS {

		code, content = service.Detail(content.(models.Content).Uuid)
		if code != env.RESPONSE_SUCCESS {
			common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

			c.Abort()
			return
		}

		common.Format(c).SetData(content).JsonResponse()
	} else {
		common.Format(c).SetStatus(env.ERROR).SetCode(env.ERROR).SetMessage(env.MsgFlags[env.ERROR]).JsonResponse()
	}
}
