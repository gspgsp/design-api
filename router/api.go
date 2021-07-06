package router

import (
	"github.com/gin-gonic/gin"
	"design-api/handler/v1/auth"

	"design-api/handler/v1/common"
	"design-api/handler/v1/user"
	middleware "design-api/middleware/auth"
	"design-api/handler/v1/slide"
	"design-api/handler/v1/category"
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
			groupCommon.POST("sms", common.SendSms) //发短信
		}

		//用户登录注册路由
		groupV1Auth := groupV1.Group("auth")
		{
			groupV1Auth.POST("register", auth.Register) //注册
			groupV1Auth.POST("login", auth.Login)       //登录
		}

		//资源路由-需登录
		groupV1User := groupV1.Group("user").Use(middleware.Auth())
		{
			groupV1User.GET("user", user.UserInfo) //用户信息
		}

		//资源路由-不需登录
		groupV1.GET("slide", slide.Slide)          //幻灯片
		groupV1.GET("style", category.Style)       //风格列表
		groupV1.GET("category", category.Category) //分类列表
	}
}
