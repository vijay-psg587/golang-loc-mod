package enums

type ErrorCodeEnum int

const (
	DB_ERROR         ErrorCodeEnum = 441
	INTERNAL         ErrorCodeEnum = 500
	SERVER_ERROR     ErrorCodeEnum = 404
	REQ_VALIDATION   ErrorCodeEnum = 442
	AWS_ERROR        ErrorCodeEnum = 443
	FILE_ERROR       ErrorCodeEnum = 444
	ACCESS_ERROR     ErrorCodeEnum = 445
	CONV_ERROR       ErrorCodeEnum = 446
	CONSTRAINT_ERROR ErrorCodeEnum = 447
	REDIS_ERROR      ErrorCodeEnum = 448
)
