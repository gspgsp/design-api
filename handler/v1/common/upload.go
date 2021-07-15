package common

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/qiniu/go-sdk/auth/qbox"
//	"github.com/qiniu/go-sdk/storage"
//	"design-api/common"
//)
//
//func Upload(c *gin.Context) {
//	accessKey := "BtTK6_zEwlHRbZwFOJQ9yGfSzhIkOTVeKfkS-2oM"
//	secretKey := "KPIY6tgA3Ankt2-j3S2IMUbo81Ou7G4WV9J_4CjM"
//
//	bucket := "dsss"
//	putPolicy := storage.PutPolicy{
//		Scope: bucket,
//	}
//	mac := qbox.NewMac(accessKey, secretKey)
//	upToken := putPolicy.UploadToken(mac)
//
//	common.Format(c).SetData(upToken).JsonResponse()
//}
