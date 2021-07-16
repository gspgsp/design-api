package validator

//import (
//	"github.com/gin-gonic/gin/binding"
//	"gopkg.in/go-playground/validator.v9" //"github.com/go-playground/validator" 没用，必须要gopkg.in下的包才可以，不懂为什么
//	"log"
//	"regexp"
//)
//
///**
//自定义验证规则 初始化
// */
//func init() {
//	//手机号验证
//	var phoneValidator validator.Func = func(fl validator.FieldLevel) bool {
//		phone, ok := fl.Field().Interface().(string)
//		if ok {
//			result, _ := regexp.MatchString(`^(1[3|4|5|6|7|8|9][0-9]\d{8})$`, phone)
//			if !result {
//				return false
//			}
//		}
//		return true
//	}
//
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		err := v.RegisterValidation("phoneValidator", phoneValidator)
//		if err != nil {
//			log.Printf("绑定验证器 err is:%v\n", err)
//		} else {
//			log.Printf("绑定验证器 success\n")
//		}
//	} else {
//		log.Printf("初始化验证器失败\n")
//	}
//}
