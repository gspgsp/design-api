package service

import (
	"design-api/common"
	"design-api/common/env"
	"design-api/model"
	"design-api/util"
)

/**
风格列表
 */
func StyleList() (int, interface{}) {
	w := "belong = 'style' and status = 1"
	s := "id,belong,name,mb_cover_picture"
	code, data := category(w, s)

	return code, data
}

/**
分类列表
 */
func CategoryList() (int, interface{}) {
	w := "status = 1"
	s := "id,abbreviation,name,belong"
	code, data := category(w, s)

	return code, data
}

/**
分类
 */
func category(w, s string) (int, interface{}) {
	categories := make([]models.Category, 1)

	if err := common.Db.Table("q_categories").Where(w).Select(s).Order("created_at desc").Find(&categories).Error; err != nil {
		return env.ACCOUNT_ERROR, categories
	}

	data := util.SplitSlice(categories)

	return env.RESPONSE_SUCCESS, data
}
