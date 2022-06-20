package models

import (
	"encoding/json"

	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
)

type AWSSSMParamModel map[string]string

func (m *AWSSSMParamModel) ToJSONString() (string, error) {
	if marshalledData, err := json.Marshal(m); err == nil {
		return string(marshalledData), nil
	} else {
		customErr := utilService.CreateCustomError(err, enums.CONV_ERROR, "500", enums.DEFAULT_LAYOUT, "Cannot convert ssm param map to string", enums.INTERNAL_ERROR)
		return "", customErr
	}
}
