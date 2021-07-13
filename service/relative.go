package service

import (
	"design-api/common"
	"design-api/model"
	"design-api/common/env"
	"log"
)

const (
	limit = 6
)

/**
相关列表
 */
func RelativeList(uuid interface{}) (int, interface{}) {
	contents := make([]models.Content, 0)

	uid, ok := uuid.(string)
	if ok {
		if err := common.Db.Table("q_contents").Where("status = 1 and uuid <> ? ", uid).Select("id,uuid,mb_cover_picture,title,sub_title,search_count,description").Limit(limit).Find(&contents).Error; err == nil {
			return env.RESPONSE_SUCCESS, contents
		}
	}

	ids, ok := uuid.([]int64)
	if ok {
		if err := common.Db.Table("q_contents").Where("status = 1 and id not in(?)", ids).Select("id,uuid,mb_cover_picture,title,sub_title,search_count,description").Limit(limit).Find(&contents).Error; err == nil {
			return env.RESPONSE_SUCCESS, contents
		}
	}

	log.Printf("contents is:%f\v", contents)
	return env.RESPONSE_FAIL, nil
}
