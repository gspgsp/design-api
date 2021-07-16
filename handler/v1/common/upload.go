package common

import (
	"design-api/common"
	"design-api/config"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// Upload /**上传token
func Upload(c *gin.Context) {
	putPolicy := storage.PutPolicy{
		Scope: config.Config.QiNiu.Bucket,
	}
	mac := qbox.NewMac(config.Config.QiNiu.AccessKey, config.Config.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	common.Format(c).SetData(upToken).JsonResponse()
}
