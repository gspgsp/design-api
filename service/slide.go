package service

import (
	"design-api/common/env"
	"design-api/model"
	"design-api/common"
)

/**
幻灯片
 */
func SlideList() (int, interface{}) {
	slides := make([]models.Slide, 1)

	if err := common.Db.Table("q_slides").Where("status = 1 and device = 2").Select("id", "target_url", "carousel_url").Order("created_at desc").Find(&slides).Error; err !=nil {
		return env.ACCOUNT_ERROR, slides
	}

	return env.RESPONSE_SUCCESS, slides
}
