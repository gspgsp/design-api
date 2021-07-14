package service

import (
	"design-api/model"
	"design-api/common"
	"design-api/common/env"
	"strconv"
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

/**
设计师素材
 */
func GetContentList(uuid string) (int, interface{}) {
	id, err := getUserId(uuid)
	if err != nil {
		return env.INVALID_PARAMS, nil
	}

	sql := "select c.id, c.uuid, c.mb_cover_picture, c.title, c.sub_title, c.size from q_contents as c where c.id in(select content_id from q_content_designer as cd where cd.designer_id = " + `'` + strconv.FormatInt(id, 10) + `'` + ")"
	contents := make([]models.Content, 0)
	common.Db.Raw(sql).Scan(&contents)

	return env.RESPONSE_SUCCESS, contents
}

/**
设计师粉丝
 */
func GetFansList(uuid string) (int, interface{}) {
	id, err := getUserId(uuid)
	if err != nil {
		return env.INVALID_PARAMS, nil
	}

	sql := "select u.id, u.nickname, u.avatar from q_users as u where u.id in(select user_id from q_fans as f where designer_id = " + `'` + strconv.FormatInt(id, 10) + `'` + ")"
	users := make([]models.User, 0)
	common.Db.Raw(sql).Scan(&users)

	return env.RESPONSE_SUCCESS, users
}
