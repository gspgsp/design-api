package auth

import (
	"github.com/gin-gonic/gin"
	"design-api/common"
	"log"
	"design-api/model"
	noSqlLog "design-api/common/log"
	"design-api/config"
	"go.mongodb.org/mongo-driver/bson"
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

	//insert mongodb
	noSqlLog.Init("mongodb://" + config.Config.Mongodb.MongodbUsername + ":" + config.Config.Mongodb.MongodbPassword + "@" + config.Config.Mongodb.MongodbHost + ":" + config.Config.Mongodb.MongodbPort)
	mgo := noSqlLog.NewMgo("design-api", "login")
	res := mgo.InsertOne(bson.D{{"user", "1000"}, {"action", "login"}})

	log.Printf("mongodb 插入返回ID:%s", res.InsertedID)
	//

	common.Format(c).SetData(slides).JsonResponse()
}
