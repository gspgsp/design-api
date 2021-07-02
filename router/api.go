package router

import (
	"github.com/gin-gonic/gin"
	"design-api/handler/v1/auth"

	"design-api/handler/v1/common"
	"design-api/handler/v1/user"
	middleware "design-api/middleware/auth"
)

/**
初始化路由
 */
func InitRouter(r *gin.Engine) {
	//v1
	groupV1 := r.Group("v1")
	{
		//公共路由
		groupCommon := groupV1.Group("common")
		{
			groupCommon.POST("sms", common.SendSms)
		}

		//不需要登录的路由
		groupV1NAuth := groupV1.Group("auth")
		{
			groupV1NAuth.POST("login", auth.Login)
		}

		//需要授权登录的路由
		groupV1Auth := groupV1.Group("auth")
		{
			groupV1Auth.POST("register", auth.Register)
			groupV1Auth.Use(middleware.Auth()).GET("user", user.UserInfo)
		}
	}
}
