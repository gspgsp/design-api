package service

import (
	"design-api/common"
	"design-api/model"
	"design-api/common/env"
	"strings"
	"log"
)

type FilterParam struct {
	St    string
	Sp    string
	Si    string
	Order string
}

func (f *FilterParam) ContentLIst() (int, interface{}) {

	whereMul := ""
	where := ""
	order := ""

	if len(f.St) > 0 {
		whereMul += "ca.abbreviation = " + `'` + f.St + `'` + " or "
	}

	if len(f.Sp) > 0 {
		whereMul += "ca.abbreviation = " + `'` + f.Sp + `'` + " or"
	}

	if len(f.Si) > 0 {

		size := strings.Split(f.Si, "|")

		log.Printf("size is:%v\n", size)
		where += "c.size >= " + size[0] + " and c.size <= " + size[1]
	}

	if len(f.Order) > 0 {
		if f.Order == "nw" {
			order += " order by c.updated_at desc, c.created_at desc "
		}

		if f.Order == "ht" {
			order += "order by c.search_count desc "
		}
	}

	sql := "select c.id, c.uuid, c.title, c.sub_title, c.size, c.mb_cover_picture, GROUP_CONCAT(ca.name order by ca.name desc separator '|') as category_name from q_contents as c inner join q_content_category as cc on c.id = cc.content_id inner join q_categories as ca on ca.id = cc.category_id "
	if len(whereMul) > 0 {

		index := strings.LastIndex(whereMul, "or")

		whereMul = whereMul[0:index]

		log.Printf("whereMul is:%f and index is:%d", whereMul, index)
		sql += " where ( " + whereMul + " )"
	}

	if len(where) > 0 {
		if len(whereMul) == 0 {
			sql += " where " + where
		} else {
			sql += " and " + where
		}
	}

	sql += " group by c.id "

	if len(order) > 0 {
		sql += order
	}

	log.Printf("sql is:%f\n", sql)
	contents := make([]models.Content, 0)
	common.Db.Raw(sql).Scan(&contents)

	return env.RESPONSE_SUCCESS, contents
}
