package auth

import (
	"github.com/go-playground/validator"
	"design-api/common/env"
)

//发短信验证参数
type SmsParam struct {
	Mobile string `validate:"required"`
}

/**
验证方法
 */
func (s *SmsParam) ValidateParam() (int, interface{}) {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		return env.PARAM_REQUIRED, nil
	}

	return env.RESPONSE_SUCCESS, nil
}
