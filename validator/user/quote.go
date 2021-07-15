package user

import (
	"github.com/gin-gonic/gin"
	"design-api/model"
	"design-api/common/env"
	"design-api/validator"
)

//报价验证参数
type QuoteParam struct {
	C *gin.Context
}

/**
验证提交参数
 */
func (q *QuoteParam) ValidateParam() (int, interface{}) {
	var quote models.Quote
	if err := q.C.ShouldBind(&quote); err != nil {
		return env.PARAM_REQUIRED, validator.Translate(err)
	}

	return env.RESPONSE_SUCCESS, &quote
}
