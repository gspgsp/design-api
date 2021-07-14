package service

import (
	"design-api/model"
	"design-api/common"
	"design-api/common/env"
)

/**
设计师详情
 */
func GetDesignerDetail(uuid string) (int, interface{}) {
	sql := "select d.id, d.uuid, d.nick_name, d.photo, d.level, d.motto, d.description, ifnull(f.fans_count, 0) as fans_count, ifnull(cd.content_count, 0) as content_count from q_designers as d left join(select designer_id, count(id) as fans_count from q_fans group by designer_id) as f on d.id = f.designer_id left join(select designer_id, count(content_id) as content_count from q_content_designer group by designer_id) as cd on cd.designer_id = f.designer_id where d.uuid = " + `'` + uuid + `'`

	designer := models.Designer{}
	common.Db.Raw(sql).Scan(&designer)

	return env.RESPONSE_SUCCESS, designer
}
