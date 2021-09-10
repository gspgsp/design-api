package common

import (
	"design-api/common"
	"design-api/service"
	"github.com/gin-gonic/gin"
	"log"
)

// Captcha 获取随机图片
func Captcha(c *gin.Context) {
	captchaKey, base64Str := service.GetCaptcha()

	s, _ := common.Cache.Get(captchaKey)
	log.Printf("s is:%s", s)

	common.Format(c).SetData(map[string]interface{}{"captcha_key": captchaKey, "captcha_code": base64Str}).JsonResponse()
}
