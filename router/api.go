package router

import (
	"github.com/gin-gonic/gin"
	"design-api/handler/v1/auth"
	authMiddleware "design-api/middleware/auth"
)

/**
初始化路由
 */
func InitRouter(r *gin.Engine) {
	//v1
	groupV1 := r.Group("v1")
	{
		//不需要登录的路由
		groupV1NAuth := groupV1.Group("auth")
		{
			groupV1NAuth.Any("login", auth.Login)
		}

		//需要授权登录的路由
		groupV1Auth := groupV1.Group("auth")
		{
			groupV1Auth.Use(authMiddleware.Auth()).Any("register", auth.Register)
		}
	}
}
