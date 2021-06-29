package auth

import (
	"github.com/gin-gonic/gin"
	"design-api/util"
	"log"
	"design-api/common/env"
	"design-api/common"
)

func Register(c *gin.Context) {

	token, code := util.GenerateToken("101", "guo")

	claim, _ := c.Get("claim")
	log.Printf("claim is:%v", claim)

	if code != env.SUCCESS {
		common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()

		c.Abort()
		return
	}

	common.Format(c).SetData(map[string]string{"token_type": "Bearer", "access_token": token}).JsonResponse()
}
