package auth

import (
	"github.com/gin-gonic/gin"
	"design-api/common"
	"log"
	"design-api/model"
)

func Login(c *gin.Context) {

	rows, err := common.Db.Queryx("select id, target_url, carousel_url from q_slides where status = 1 and device = 1")
	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	slide := models.Slide{}
	var slides []models.Slide

	if rows == nil {
		log.Println("数据结果为空")
		return
	}

	if rows.Next() {
		err := rows.StructScan(&slide)
		if err != nil {
			log.Printf("数据结构化错误:%s\n", err.Error())
			return
		} else {
			slides = append(slides, slide)
		}
	}

	common.Format(c).SetData(slides).JsonResponse()
}
