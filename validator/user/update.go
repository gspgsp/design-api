package user

import (
	"design-api/common/env"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UpdateParam 更新参数
type UpdateParam struct {
	C           *gin.Context
	OperateType int
	Name        string
	NickName    string
}

// ValidateParam /**验证更新参数
func (u *UpdateParam) ValidateParam() (int, interface{}) {
	u.C.Request.ParseMultipartForm(128)
	values := u.C.Request.Form

	for i, val := range values {
		if i == "type" {
			if oType, err := strconv.Atoi(val[0]); err != nil {
				return env.INVALID_PARAMS, nil
			} else {
				u.OperateType = oType
			}
		} else if i == "name" {
			u.Name = val[0]
		} else if i == "nick_name" {
			u.NickName = val[0]
		}
	}

	return env.RESPONSE_SUCCESS, u
}
