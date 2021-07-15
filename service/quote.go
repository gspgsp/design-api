package service

import (
	"design-api/model"
	"design-api/common"
	"design-api/common/env"
)

/**
保存报价
 */
func StoreQuote(q interface{}) (int, interface{}) {

	quote := q.(*models.Quote)
	if err := common.Db.Create(quote).Error; err != nil {
		return env.ERROR, nil
	}

	return env.RESPONSE_SUCCESS, nil
}
