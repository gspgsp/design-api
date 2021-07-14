package designer

import (
	"github.com/gin-gonic/gin"
	"design-api/common/env"
	"design-api/model"
)

//详情验证参数
type DetailParam struct {
	C *gin.Context
}

/**
验证uuid
 */
func (d *DetailParam) ValidateParam() (int, interface{}) {
	var designer models.Designer

	if err := d.C.ShouldBindUri(&designer); err != nil {
		return env.PARAM_REQUIRED, designer
	}

	return env.RESPONSE_SUCCESS, designer
}
