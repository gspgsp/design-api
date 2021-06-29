package auth

import (
	"github.com/go-playground/validator"
	"design-api/common/env"
)

type SmsParam struct {
	Mobile string `validate:"required"`
}

func (s *SmsParam) ValidateParam() (int) {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		return env.PARAM_REQUIRED
	}

	return env.RESPONSE_SUCCESS
}
