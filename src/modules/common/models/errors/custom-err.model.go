package errors

import (
	"encoding/json"
	"strconv"

	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
)

type CustomErrModel struct {
	Timestamp string
	Status    string
	Code      enums.ErrorCodeEnum
	Message   string
	Type      enums.ErrorTypeEnum
	Err       error
}

func (errModel *CustomErrModel) ToString() string {
	if str, err := json.Marshal(errModel); err == nil {
		convertedStr, _ := strconv.Unquote(string(str))
		return string(convertedStr)
	} else {
		panic(err)
	}
}

func (err *CustomErrModel) Error() string {
	if err.Code != 0 {
		return err.Err.Error()
	} else {
		return ""
	}
}
