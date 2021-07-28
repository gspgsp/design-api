package service

import (
	"bytes"
	"design-api/common"
	"design-api/common/env"
	"design-api/model"
	"design-api/util"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
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

/**
解析json参数
*/
func parseJsonParams(body io.ReadCloser) (map[string]interface{}, error) {
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	values, err := util.JsonToMap(bytes.NewBuffer(result).String())
	if err != nil {
		return nil, err
	}

	return values, nil
}
