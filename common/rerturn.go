package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回响应，默认返回成功无数据
func Format(ctx *gin.Context) *Response {
	return &Response{http.StatusOK, http.StatusText(http.StatusOK), []interface{}{}, ctx}
}

// 定义 Response 结构体
type Response struct {
	Code    int          `json:"code"`
	Message string       `json:"msg"`
	Data    interface{}  `json:"data"`
	ctx     *gin.Context `json:"-"`
}

// 设置错误码
func (res *Response) SetCode(code int) *Response {
	res.Code = code
	return res
}

// 设置错误信息
func (res *Response) SetMessage(msg string) *Response {
	res.Message = msg
	return res
}

// 设置返回数据
func (res *Response) SetData(data interface{}) *Response {
	res.Data = data
	return res
}

// json返回结果
func (res *Response) JsonResponse() {
	res.ctx.JSON(res.Code, res)
}
