package enums

type ErrorTypeEnum string

const (
	DB_ACCESS_ERROR    ErrorTypeEnum = "DATABASE Access error"
	AWS_INTERNAL_ERROR ErrorTypeEnum = "AWS Service internal error"
	INTERNAL_ERROR     ErrorTypeEnum = "Service internal error"
	REQ_ERROR          ErrorTypeEnum = "Input Request Error"
	REQ_VAL_ERROR      ErrorTypeEnum = "Input Request Schema is an invalid one"
	RESPONSE_ERROR     ErrorTypeEnum = "Internal error. Request cannot be processed"
)
