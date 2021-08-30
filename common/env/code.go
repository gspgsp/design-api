package env

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	NO_AUTHED      = 419
	INVALID_METHOD = 404

	RESPONSE_SUCCESS = 0
	RESPONSE_FAIL    = 1

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20001
	ERROR_AUTH_TOKEN               = 20002
	ERROR_AUTH_PARSE               = 20003
	ERROR_AUTH_VALID               = 20004

	PARAM_REQUIRED = 30001

	SMS_CODE_SEND_ERROR   = 40001
	SMS_CODE_VERIFY_ERROR = 40002
	SMS_CODE_EXPIRE_ERROR = 40003

	DB_INSERT_ERROR = 50001

	ACCOUNT_ERROR = 60001

	CODE_EXPIRED     = 70001
	CODE_INVALID     = 70002
	CODE_KEY_INVALID = 70003

	DATABASE_OPERATE_ERROR = 8001
)
