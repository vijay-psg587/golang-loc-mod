package app_utils

import (
	"os"
	"reflect"
	"runtime"
	"time"

	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	models "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/utils"
)

// This method is to format time with the input layout
func GetTimeStamp(layout enums.TimeLayoutEnum) string {
	switch layout {
	case enums.ISO_LAYOUT:
		return time.Now().Format(utils.LAYOUT_US)
	case enums.US_LAYOUT:
		return time.Now().Format(utils.LAYOUT_US)
	case enums.DEFAULT_LAYOUT:
		return time.Now().Format(utils.LAYOUT_DEFAULT)
	default:
		return time.Now().Format(utils.LAYOUT_DEFAULT)
	}

}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// This method create a custom error given the list of inputs
func CreateCustomError(err error, code enums.ErrorCodeEnum, status string, layout enums.TimeLayoutEnum, message string, errType enums.ErrorTypeEnum) error {
	if layout == "" {
		layout = enums.DEFAULT_LAYOUT
	}
	customErr := &models.CustomErrModel{
		Err:       err,
		Message:   message,
		Timestamp: GetTimeStamp(layout),
		Code:      code,
		Status:    status,
		Type:      errType,
	}
	return customErr
}

// This method returns the environment variable or fallback if not found
func GetEnvWithFallback(key string, fallback interface{}) interface{} {
	if val, boolFlag := os.LookupEnv(key); boolFlag {
		// means the key is found
		return val
	} else {
		return fallback
	}
}
