package router

import (
	"github.com/gin-gonic/gin"
	"design-api/handler/v1/auth"

	"design-api/handler/v1/common"
	"design-api/handler/v1/user"
	middleware "design-api/middleware/auth"
	"design-api/handler/v1/slide"
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

		//用户登录注册路由
		groupV1Auth := groupV1.Group("auth")
		{
			groupV1Auth.POST("register", auth.Register)
			groupV1Auth.POST("login", auth.Login)
		}

		//资源路由-不需登录
		groupV1Slide := groupV1.Group("slide").Use(middleware.WithAccessToken())
		{
			groupV1Slide.GET("slide", slide.Slide)
		}

		//资源路由-需登录
		groupV1User := groupV1.Group("user").Use(middleware.Auth())
		{
			groupV1User.GET("user", user.UserInfo)
		}
	}
}
