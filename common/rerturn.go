package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回响应，默认返回成功无数据
func Response(c *gin.Context) *Return {
	return &Return{http.StatusOK, http.StatusText(http.StatusOK), []interface{}{}, c}
}

// 定义 Return 结构体
type Return struct {
	Code    int          `json:"code"`
	Message string       `json:"msg"`
	Data    interface{}  `json:"data"`
	ctx     *gin.Context `json:"-"`
}

// 设置错误码
func (res *Return) SetCode(code int) *Return {
	res.Code = code
	return res
}

// 设置错误信息
func (res *Return) SetMessage(msg string) *Return {
	res.Message = msg
	return res
}

// 设置返回数据
func (res *Return) SetData(data interface{}) *Return {
	res.Data = data
	return res
}

// json返回结果
func (res *Return) JsonReturn() {
	res.ctx.JSON(res.Code, res)
}
