package common

import (
	"github.com/gin-gonic/gin"
	"design-api/common/env"
)

// 返回响应，默认返回成功无数据
func Format(ctx *gin.Context) *Response {
	return &Response{env.SUCCESS, env.RESPONSE_SUCCESS, env.MsgFlags[env.RESPONSE_SUCCESS], []interface{}{}, ctx}
}

// 定义 Response 结构体
type Response struct {
	Status  int          `json:"-"`
	Code    int          `json:"code"`
	Message string       `json:"msg"`
	Data    interface{}  `json:"data"`
	ctx     *gin.Context `json:"-"`
}

//设置响应状态码
func (res *Response) SetStatus(status int) *Response {
	res.Status = status
	return res
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
	res.ctx.JSON(res.Status, res)
}
