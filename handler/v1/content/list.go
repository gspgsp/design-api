package content

import (
	"github.com/gin-gonic/gin"
	"design-api/service"
	"design-api/common/env"
	"design-api/common"
	"strconv"
)

/**
内容列表
 */
func List(c *gin.Context) {
	st := c.DefaultQuery("st", "")
	sp := c.DefaultQuery("sp", "")
	si := c.DefaultQuery("si", "")
	order := c.DefaultQuery("order", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	filterParam := &service.FilterParam{st, sp, si, order, service.Limit{page, 10}}
	code, contents := filterParam.ContentList()

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(contents).JsonResponse()
}
