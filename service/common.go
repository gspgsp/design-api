package service

import (
	"design-api/model"
	"design-api/common"
	"github.com/pkg/errors"
	"design-api/common/env"
)

/**
uuid-->id的映射
 */
func getUserId(uuid string) (int64, error) {

	//TODO:: uuid--->id id--->uuid的映射，可以在后台添加这个功能，在添加designer的时候

	designer := models.Designer{}
	if err := common.Db.Table("q_designers").Where("uuid = " + `'` + uuid + `'`).Select("id").Find(&designer).Error; err != nil {
		return 0, errors.New(env.MsgFlags[env.ERROR])
	}

	if designer.Id == 0 {
		return 0, errors.New(env.MsgFlags[env.INVALID_PARAMS])
	}

	return designer.Id, nil
}
