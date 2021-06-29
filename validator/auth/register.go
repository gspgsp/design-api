package auth

import (
	"github.com/go-playground/validator"
	"design-api/common/env"
)

//注册证参数
type RegisterParam struct {
	SmsParam
	CodeKey  string `validate:"required"`
	Code     string `validate:"required"`
	Password string `validate:"required,validateLength"`
}

/**
解析POST参数
 */
func (r *RegisterParam) ParseParam(params map[string][]string) {

	for i, val := range params {
		if i == "mobile" {
			r.Mobile = val[0]
		} else if i == "code_key" {
			r.CodeKey = val[0]
		} else if i == "code" {
			r.Code = val[0]
		} else if i == "password" {
			r.Password = val[0]
		}
	}
}

/**
验证POST参数
 */
func (r *RegisterParam) ValidateParam() (int) {
	validate := validator.New()

	validate.RegisterValidation("validateLength", ValidateLengthFunc)

	err := validate.Struct(r)
	if err != nil {
		return env.PARAM_REQUIRED
	}

	return env.RESPONSE_SUCCESS
}

/**
自定义验证方法 | 验证字符串长度
 */
func ValidateLengthFunc(field validator.FieldLevel) bool {
	codeLen := len(field.Field().String())
	if codeLen < 6 || codeLen > 20 {
		return false
	}

	return true
}
