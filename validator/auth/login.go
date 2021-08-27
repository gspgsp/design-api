package auth

import (
	"design-api/common/env"
	"github.com/go-playground/validator"
	"strconv"
)

type LoginParam struct {
	SmsParam
	OperateType int64 `validate:"oneof=1 2"`
	Password    string
	Code        string
	CodeKey     string
}

func (l *LoginParam) ParseParam(params map[string]interface{}) {
	for i, val := range params {
		if i == "mobile" {
			l.Mobile = val.(string)
		} else if i == "password" {
			l.Password = val.(string)
		} else if i == "operate_type" {
			t := val.(float64)
			l.OperateType = int64(t)
		} else if i == "code" {
			t, ok := val.(float64)
			if ok {
				l.Code = strconv.FormatFloat(t, 'E', -1, 64)
			} else {
				l.Code = val.(string)
			}
			l.Code = val.(string)
		} else if i == "code_key" {
			l.CodeKey = val.(string)
		}
	}
}

func (l *LoginParam) ValidateParam() (int, interface{}) {
	validate := validator.New()
	err := validate.Struct(l)
	if err != nil {
		return env.PARAM_REQUIRED, nil
	}

	if l.OperateType == 1 && len(l.Password) == 0 {
		return env.PARAM_REQUIRED, nil
	}

	if l.OperateType == 2 && (len(l.Code) == 0 || len(l.CodeKey) == 0) {
		return env.PARAM_REQUIRED, nil
	}

	return env.RESPONSE_SUCCESS, nil
}
