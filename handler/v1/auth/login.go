package auth

import (
	"github.com/gin-gonic/gin"
	"design-api/common"
	"design-api/service"
	"design-api/common/env"
	"design-api/util"
)

/**
登录
 */
func Login(c *gin.Context) {

	code, user := service.Login(c)

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()
	} else {
		token, _ := util.GenerateToken(*user.ID)
		common.Format(c).SetData(map[string]string{"token_type": "Bearer", "access_token": token}).JsonResponse()
	}

	//rows, err := common.Db.Queryx("select id, target_url, carousel_url from q_slides where status = 1 and device = 1")
	//if err != nil {
	//	log.Printf("查询错误:%s", err.Error())
	//	return
	//}
	//
	//slide := models.Slide{}
	//var slides []models.Slide
	//
	//if rows == nil {
	//	log.Println("数据结果为空")
	//	return
	//}
	//
	//if rows.Next() {
	//	err := rows.StructScan(&slide)
	//	if err != nil {
	//		log.Printf("数据结构化错误:%s\n", err.Error())
	//		return
	//	} else {
	//		slides = append(slides, slide)
	//	}
	//}
	//
	////insert mongodb
	//mgo := noSqlLog.NewMgo("login")

	//res := mgo.InsertOne(bson.D{{"user", "1006"}, {"action", "login"}})
	//
	//log.Printf("mongodb 插入返回ID:%s", res.InsertedID)
	////
	//
	////insert redis
	//rd := noSqlLog.NewRd(7)
	//rd.Set("test1024", 1025, -1)
	////
	//
	//common.Format(c).SetData(slides).JsonResponse()
}
