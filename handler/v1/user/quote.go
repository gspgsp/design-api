package user

import (
	"github.com/gin-gonic/gin"
	"design-api/validator/user"
	"design-api/common/env"
	"design-api/common"
	"design-api/service"
	"design-api/model"
	"design-api/util"
	"strconv"
)

/**
报价
 */
func Quote(c *gin.Context) {

	quoteParam := user.QuoteParam{c}
	if code, quote := quoteParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
		userId, _ := c.Get("userId")
		quote.(*models.Quote).UserId, _ = strconv.Atoi(util.Strval(userId))
		cd, _ := service.StoreQuote(quote)
		if cd == env.RESPONSE_SUCCESS {
			common.Format(c).JsonResponse()
		} else {
			common.Format(c).SetCode(env.ERROR).JsonResponse()
		}
	} else {
		common.Format(c).SetCode(env.INVALID_PARAMS).SetData(quote).JsonResponse()
	}
}
