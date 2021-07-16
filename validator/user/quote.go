package user

import (
	"design-api/common/env"
	"design-api/model"
	"github.com/gin-gonic/gin"
	//_ "design-api/validator"
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
		//return env.PARAM_REQUIRED, validator.Translate(err)
		//TODO:: 先把这个包（validator.Translate）的内容隐藏起来，将init 方法改为 initRename，因为 验证package来自不同的包，会有问题:
		//报： panic: interface conversion: interface {} is *validator.Validate, not *validator.Validate (types from different packages)
		return env.PARAM_REQUIRED, "参数错误"
	}

	return env.RESPONSE_SUCCESS, &quote
}
