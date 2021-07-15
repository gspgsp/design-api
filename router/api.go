package router

import (
	"github.com/gin-gonic/gin"
	"design-api/handler/v1/auth"

	"design-api/handler/v1/common"
	"design-api/handler/v1/user"
	middleware "design-api/middleware/auth"
	"design-api/handler/v1/slide"
	"design-api/handler/v1/category"
	"design-api/handler/v1/content"
	"design-api/handler/v1/designer"
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
			groupCommon.POST("sms", common.SendSms)   //发短信
			//groupCommon.POST("upload", common.Upload) //上传
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
			groupV1User.POST("quote", user.Quote)  //报价
		}

		//资源路由-不需登录
		{
			groupV1.GET("slide", slide.Slide)                       //幻灯片
			groupV1.GET("style", category.Style)                    //风格列表
			groupV1.GET("category", category.Category)              //分类列表
			groupV1.GET("content", content.List)                    //内容列表
			groupV1.GET("content/detail/:uuid", content.Detail)     //内容详情
			groupV1.GET("content/relative/:uuid", content.Relative) //相关列表
			groupV1.GET("designer/detail/:uuid", designer.Detail)   //设计师详情
			groupV1.GET("designer/content/:uuid", designer.Content) //设计师素材
			groupV1.GET("designer/fans/:uuid", designer.Fans)       //设计师粉丝
		}
	}
}
