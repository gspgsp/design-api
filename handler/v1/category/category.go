package category

import (
	"github.com/gin-gonic/gin"
	"design-api/common/env"
	"design-api/service"
	"design-api/common"
)

/**
分类列表
 */
func Category(c *gin.Context) {
	code, categories := service.CategoryList()

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(categories).JsonResponse()
}
