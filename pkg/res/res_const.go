package res

const (
	SUCCESS        = 2000
	ERROR          = 5000
	INVALID_PARAMS = 4000
	FAIL_VALIDATE  = 4100
	FAIL_AUTH      = 4200
	FAIL_OPER      = 8000

	//用户相关
	ERROR_EXIST_USER     = 10001
	ERROR_NOT_EXIST_USER = 10002
	ERROR_PASS_USER      = 10003
	ERROR_CAPTCHA_USER   = 10004
	FAIL_LOGOUT_USER     = 10005

	//token相关
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_AUTH_CHECK_FAIL          = 20005

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003

	FAIL_SEND_EMAIl    = 40001
	FAIL_VALIDATE_CODE = 40002
)
