package validator

//import (
//	"github.com/gin-gonic/gin/binding"
//	"github.com/go-playground/locales/zh"
//	universal "github.com/go-playground/universal-translator"
//	"gopkg.in/go-playground/validator.v9" //"github.com/go-playground/validator" 没用，必须要gopkg.in下的包才可以，不懂为什么
//	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
//)

//var (
//	uni      *universal.UniversalTranslator
//	validate *validator.Validate
//	trans    universal.Translator
//)

/**
自定义汉化语言包 初始化
 */
//func init() {
//	//注册翻译器
//	zh := zh.New()
//	uni = universal.New(zh, zh)
//
//	trans, _ = uni.GetTranslator("zh")
//
//	//获取gin的校验器
//	validate := binding.Validator.Engine().(*validator.Validate)
//	//注册翻译器
//	zh_translations.RegisterDefaultTranslations(validate, trans)
//
//	//给自定义的tag 添加语言包提示信息
//	//Validate = validator.New() 这样用的是validator的实例，Validate.RegisterTranslation 不对自定义的phoneValidator tag生效，只对原生的required tag生效
//	//如果想要让 自定义的phoneValidator tag生效 ，需要 validate.RegisterTranslation，这是用 gin 的校验器 注册的
//	validate.RegisterTranslation("phoneValidator", trans, func(ut universal.Translator) error {
//		return ut.Add("phoneValidator", "{0}手机号格式不对", true) // see universal-translator for details
//	}, func(ut universal.Translator, fe validator.FieldError) string {
//		t, _ := ut.T("phoneValidator", fe.Field())
//		return t
//	})
//}

//Translate 翻译错误信息
//func Translate(err error) map[string][]string {
//	var result = make(map[string][]string)
//
//	errors := err.(validator.ValidationErrors)
//
//	for _, err := range errors {
//		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
//	}
//	return result
//}

//func Translate(err error) string {
//	var result string
//
//	errors := err.(validator.ValidationErrors)
//
//	for _, err := range errors {
//		result += err.Translate(trans)+","
//	}
//
//	return result[:len(result) -1]
//}
