package auth

import (
	"design-api/common"
	"design-api/common/env"
	mongo "design-api/common/log"
	"design-api/service"
	"design-api/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Login /**登录
func Login(c *gin.Context) {

	code, user := service.Login(c)

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()
	} else {
		token, _ := util.GenerateToken(*user.ID)

		mgo := mongo.NewMgo("login")
		mgo.InsertOne(bson.D{{"user", user.ID}, {"token", token.(map[string]interface{})["access_token"]}})

		common.Format(c).SetData(map[string]string{"token_type": "Bearer", "access_token": token.(map[string]interface{})["access_token"].(string), "expire_at":token.(map[string]interface{})["expire_at"].(string)}).JsonResponse()
	}

	////insert redis
	//rd := noSqlLog.NewRd(7)
	//rd.Set("test1024", 1025, -1)
	////
	//
	//common.Format(c).SetData(slides).JsonResponse()
}

/**
刷新token
*/
func Refresh(c *gin.Context) {
	refreshToken := c.Query("refresh_token")
	token, code := util.RefreshToken(refreshToken)

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()
	} else {
		common.Format(c).SetData(map[string]string{"token_type": "Bearer", "access_token": token.(map[string]interface{})["access_token"].(string), "expire_at":token.(map[string]interface{})["expire_at"].(string)}).JsonResponse()
	}
}
