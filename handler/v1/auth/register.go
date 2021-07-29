package auth

import (
	"design-api/common"
	"design-api/common/env"
	"design-api/service"
	"design-api/util"
	"github.com/gin-gonic/gin"
)

/**
注册
 */
func Register(c *gin.Context) {

	//注册操作
	code, userId := service.Register(c)

	if code != env.RESPONSE_SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	token, _ := util.GenerateToken(userId)
	common.Format(c).SetData(map[string]string{"token_type": "Bearer", "access_token": token.(map[string]interface{})["access_token"].(string), "expire_at":token.(map[string]interface{})["expire_at"].(string)}).JsonResponse()
}
