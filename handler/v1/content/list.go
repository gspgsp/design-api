package content

import (
	"github.com/gin-gonic/gin"
	"design-api/service"
	"design-api/common/env"
	"design-api/common"
)

func List(c *gin.Context)  {

	filterParam := &service.FilterParam{"","","0|60","nw"}

	code, contents := filterParam.ContentLIst()

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(contents).JsonResponse()
}
