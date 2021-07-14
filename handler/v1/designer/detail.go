package designer

import (
	"github.com/gin-gonic/gin"
	"design-api/service"
	"design-api/common/env"
	"design-api/validator/designer"
	"design-api/model"
	"design-api/common"
)

/**
设计师详情
 */
func Detail(c *gin.Context) {

	detailParam := designer.DetailParam{c}
	if code, designer := detailParam.ValidateParam(); code == env.RESPONSE_SUCCESS {

		code, designer = service.GetDesignerDetail(designer.(models.Designer).Uuid)
		if code != env.RESPONSE_SUCCESS {
			common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

			c.Abort()
			return
		}

		common.Format(c).SetData(designer).JsonResponse()
	} else {
		common.Format(c).SetStatus(env.ERROR).SetCode(env.ERROR).SetMessage(env.MsgFlags[env.ERROR]).JsonResponse()
	}
}
