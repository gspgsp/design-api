package content

import (
	"github.com/gin-gonic/gin"
	"design-api/model"
	"design-api/common/env"
)

//详情验证参数
type DetailParam struct {
	C *gin.Context
}

/**
验证uuid
 */
func (d *DetailParam) ValidateParam() (int, interface{}) {
	var content models.Content

	if err := d.C.ShouldBindUri(&content); err != nil {
		return env.PARAM_REQUIRED, content
	}

	return env.RESPONSE_SUCCESS, content
}
