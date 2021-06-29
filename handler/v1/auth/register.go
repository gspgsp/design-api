package auth

import (
	"github.com/gin-gonic/gin"
	//"design-api/util"
	//"log"
	"design-api/common/env"
	"design-api/common"
	"design-api/validator/auth"
	"log"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	//注册操作
	c.Request.ParseForm()
	values := c.Request.Form

	registerParam := &auth.RegisterParam{}
	registerParam.ParseParam(values)

	if code := registerParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
		//TODO::注册业务
		log.Printf("registerParam is:%v\n", registerParam)

		bytes, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		log.Printf("密码为:%s\n", string(bytes))

		err := bcrypt.CompareHashAndPassword([]byte("$2y$10$hyLCq9sR3lu98F3OI.VKpOuI.M/CLvZSTWTf6VzaVgxilrqmvy3Yy"), []byte("123456"))
		if err == nil {
			log.Printf("密码验证通过\n")
		}



	} else {
		common.Format(c).SetStatus(env.ERROR).SetCode(env.PARAM_REQUIRED).SetMessage(env.MsgFlags[env.PARAM_REQUIRED]).JsonResponse()
	}

	//登录操作
	//token, code := util.GenerateToken("101", "guo")
	//
	//claim, _ := c.Get("claim")
	//log.Printf("claim is:%v", claim)
	//
	//if code != env.SUCCESS {
	//	common.Format(c).SetStatus(env.ERROR).SetCode(code).SetMessage(env.MsgFlags[code]).JsonResponse()
	//
	//	c.Abort()
	//	return
	//}
	//
	//common.Format(c).SetData(map[string]string{"token_type": "Bearer", "access_token": token}).JsonResponse()
}
