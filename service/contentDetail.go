package service

import (
	"design-api/common"
	"design-api/model"
	"design-api/common/env"
)

/**
内容详情
 */
func Detail(uuid string) (int, interface{}) {
	sql := "select c.id, c.uuid, c.title, c.sub_title, c.search_count, c.description, c.carousel, c.city, c.size, group_concat((case when (ca.belong = 'style' or ca.belong = 'space') then ca.name else null end) order by ca.name desc separator '|') as category_name from q_contents as c inner join q_content_category as cc on c.id = cc.content_id inner join q_categories as ca on ca.id = cc.category_id where c.uuid = " + `'` + uuid + `'`

	content := models.Content{}
	common.Db.Raw(sql).Scan(&content)

	if content.Id > 0 {
		designers := make([]models.Designer, 0)
		sqlD := "select d.id, d.uuid, d.name from q_designers as d inner join q_content_designer as cd on d.id = cd.designer_id inner join q_contents as c on c.id = cd.content_id where c.uuid = " + `'` + content.Uuid + `'`
		common.Db.Raw(sqlD).Scan(&designers)

		content.DesignerList = designers
	}

	return env.RESPONSE_SUCCESS, content
}
