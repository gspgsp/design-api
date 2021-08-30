package service

import (
	"design-api/common"
	"design-api/common/env"
	mongo "design-api/common/log"
	"design-api/model"
	"design-api/util"
	"design-api/validator/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

const DEFAULT_PASSWORD = "123456"

func Login(c *gin.Context) (int, models.User) {
	//postman可以传form，也可以传json，对于php来讲 框架都可以解析，但是gin的话，必须区别对待.直接form是获取不到参数的
	//c.Request.ParseForm()
	//values := c.Request.PostForm
	var user models.User
	//
	method := c.Request.Method
	if method == "POST" {
		values, err := parseJsonParams(c.Request.Body)
		if err != nil {
			return env.PARAM_REQUIRED, user
		}

		loginParam := &auth.LoginParam{}
		loginParam.ParseParam(values)

		if code, _ := loginParam.ValidateParam(); code == env.RESPONSE_SUCCESS {
			err := common.Db.Where("mobile = "+loginParam.Mobile).Select("id", "name", "nickname", "avatar", "mobile", "email", "password").Find(&user).Error
			if err != nil {
				return env.ACCOUNT_ERROR, user
			}

			if loginParam.OperateType == 1 {
				err = util.PasswordCheck([]byte(user.Password), []byte(loginParam.Password))
				if err != nil {
					return env.ACCOUNT_ERROR, user
				}
			} else if loginParam.OperateType == 2 {
				//
				var d mongo.SmsMongoInfo
				mongo.NewMgo("sms_code").GetOne(bson.M{"mobile": user.Mobile}, &d)
				if d.Code != loginParam.Code {
					return env.SMS_CODE_VERIFY_ERROR, user
				}

				if d.CodeKey != loginParam.CodeKey {
					return env.SMS_CODE_KEY_INVALID, user
				}

				if d.ExpireAt < time.Now().Unix() {

					log.Printf("tt is:%v", time.Now().Unix() + 3600)

					return env.SMS_CODE_EXPIRE_ERROR, user
				}

				if d.Mobile != loginParam.Mobile {
					return env.SMS_CODE_INVALID_MOBILE, user
				}

				if (user == models.User{}) {
					password, _ := util.PasswordHash(DEFAULT_PASSWORD)
					uuid := fmt.Sprintf("%s", uuid2.NewV4())
					randStr := util.RandStr{}
					nickname := randStr.Generate(10)

					user = models.User{Mobile: loginParam.Mobile, Password: password, Uuid: uuid, Nickname: nickname, RegisterAt: time.Now().Unix(), RegisterIp: util.ClientIp(c), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
					err = common.Db.Create(&user).Error
					if err != nil {
						return env.DATABASE_OPERATE_ERROR, user
					}
				}
			}

			return env.RESPONSE_SUCCESS, user
		} else {
			return env.PARAM_REQUIRED, user
		}
	}

	return env.INVALID_METHOD, user
}
