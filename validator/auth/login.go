package auth

import (
	"design-api/common/env"
	"github.com/go-playground/validator"
)

type LoginParam struct {
	SmsParam
	Password string `validate:"required"`
}

func (l *LoginParam) ParseParam(params map[string]interface{}) {

	for i, val := range params {
		if i == "mobile" {
			l.Mobile = val.(string)
		} else if i == "password" {
			l.Password = val.(string)
		}
	}
}

func (l *LoginParam) ValidateParam() (int, interface{}) {
	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		return env.PARAM_REQUIRED, nil
	}

	return env.RESPONSE_SUCCESS, nil
}
