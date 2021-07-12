package auth

import (
	"design-api/common/env"
	"github.com/go-playground/validator"
)

type LoginParam struct {
	SmsParam
	Password string `validate:"required"`
}

func (l *LoginParam) ParseParam(params map[string][]string) {

	for i, val := range params {
		if i == "mobile" {
			l.Mobile = val[0]
		} else if i == "password" {
			l.Password = val[0]
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
