package env

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	NO_AUTHED:      "请先授权登录",

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH_PARSE:               "Token解析错误",
	ERROR_AUTH_VALID:               "Token验证错误",
}
