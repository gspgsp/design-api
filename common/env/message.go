package env

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	NO_AUTHED:      "请先授权登录",
	INVALID_METHOD: "非法请求方式",

	RESPONSE_SUCCESS: "请求成功",
	RESPONSE_FAIL:    "请求失败",

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH_PARSE:               "Token解析错误",
	ERROR_AUTH_VALID:               "Token验证错误",

	PARAM_REQUIRED: "参数必须",

	SMS_CODE_SEND_ERROR:   "验证码发送失败",
	SMS_CODE_VERIFY_ERROR: "验证码验证错误",
	SMS_CODE_EXPIRE_ERROR: "验证码已过期",

	DB_INSERT_ERROR: "数据插入错误",

	ACCOUNT_ERROR: "账号或密码错误",

	CODE_EXPIRED:     "验证码已失效",
	CODE_INVALID:     "验证码无效",
	CODE_KEY_INVALID: "验证码Key无效",

	DATABASE_OPERATE_ERROR: "数据库操作失败",
}
