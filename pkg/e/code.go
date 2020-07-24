package e

const (
	SUCCESS                        = 20000
	ERROR                          = 500
	INVALID_PARAMS                 = 400
	ERROR_EXIST_TAG                = 10001
	ERROR_NOT_EXIST_TAG            = 10002
	ERROR_NOT_EXIST_ARTICLE        = 10003
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	LOGIN_INVALID_PARAMS           = 30000
	LOGIN_FAILED                   = 30001
	DUPLICATE_USERNAME             = 30002
	USER_NOT_FOUND                 = 30003
	USERINFO_CHECK_JWT_FAILED      = 30004
	HEADER_NEED_TOKEN              = 30005
	MODIFY_VERIFY_PASSWORD_FAILED  = 30006
)
