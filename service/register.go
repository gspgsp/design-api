package service

import (
	"design-api/validator/auth"
	"design-api/common/env"
	"design-api/common"
	"design-api/model"
	"design-api/util"
	"log"
	uuid2 "github.com/satori/go.uuid"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) (int) {
	c.Request.ParseForm()
	values := c.Request.Form

	registerParam := &auth.RegisterParam{}
	registerParam.ParseParam(values)

	if code := registerParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
		smsCode, ok := common.Cache.Get(registerParam.CodeKey)
		if !ok {
			return env.SMS_CODE_EXPIRE_ERROR
		}

		if smsCode != registerParam.Code {
			return env.SMS_CODE_VERIFY_ERROR
		}

		//TODO::注册业务 要不要 DAO 层?
		password, _ := util.PasswordHash(registerParam.Password)
		uuid := fmt.Sprintf("%s", uuid2.NewV4())
		randStr := util.RandStr{}
		nickname := randStr.Generate(10)

		user := models.User{Mobile: registerParam.Mobile, Password: password, Uuid: uuid, NickName: nickname, RegisterAt: time.Now().Unix(), RegisterIp: util.ClientIp(c), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}

		res, err := common.Db.NamedExec("insert into q_users(uuid,mobile,password,nickname,register_at,register_ip,created_at,updated_at) values(:uuid,:mobile,:password,:nickname,:register_at,:register_ip,:created_at,:updated_at)", map[string]interface{}{
			"mobile":      user.Mobile,
			"password":    user.Password,
			"uuid":        user.Uuid,
			"nickname":    user.NickName,
			"register_at": user.RegisterAt,
			"register_ip": user.RegisterIp,
			"created_at":  user.CreatedAt,
			"updated_at":  user.UpdatedAt,
		})

		if err != nil {
			log.Printf("insert err is:%v\n", err.Error())
			return env.DB_INSERT_ERROR
		}

		userId, _ := res.LastInsertId()
		log.Printf("userId is:%d\n", userId)

		return env.RESPONSE_SUCCESS
	} else {
		return env.PARAM_REQUIRED
	}
}
