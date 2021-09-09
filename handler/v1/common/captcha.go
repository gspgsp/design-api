package common

import (
	"design-api/common"
	"design-api/service"
	"github.com/gin-gonic/gin"
)

// Captcha 获取随机图片
func Captcha(c *gin.Context) {
	base64Str := service.GetCaptcha()
	common.Format(c).SetData(map[string]interface{}{"captcha": base64Str}).JsonResponse()
}
