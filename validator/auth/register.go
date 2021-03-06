package auth

import (
	"design-api/common/env"
	"github.com/go-playground/validator"
)

// RegisterParam 注册证参数
type RegisterParam struct {
	SmsParam
	CodeKey  string `validate:"required"`
	Code     string `validate:"required"`
	Password string `validate:"required,validateLength"`
}

// ParseParam /**解析POST参数
func (r *RegisterParam) ParseParam(params map[string]string) {
	for i, val := range params {
		if i == "mobile" {
			r.Mobile = val
		} else if i == "code_key" {
			r.CodeKey = val
		} else if i == "code" {
			r.Code = val
		} else if i == "password" {
			r.Password = val
		}
	}
}

// ValidateParam /**验证POST参数
func (r *RegisterParam) ValidateParam() (int, interface{}) {
	validate := validator.New()

	validate.RegisterValidation("validateLength", ValidateLengthFunc)

	err := validate.Struct(r)
	if err != nil {
		return env.PARAM_REQUIRED, nil
	}

	return env.RESPONSE_SUCCESS, nil
}

// ValidateLengthFunc /**自定义验证方法 | 验证字符串长度
func ValidateLengthFunc(field validator.FieldLevel) bool {
	codeLen := len(field.Field().String())
	if codeLen < 6 || codeLen > 20 {
		return false
	}

	return true
}
