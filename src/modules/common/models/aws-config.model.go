package models

import (
	"encoding/json"

	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
)

type AWSConfigModel struct {
	Region          string `json:"AWS_REGION"`
	DYNAMO_DB_TABLE string `json:"DYNAMO_TABLE"`
	MAX_RETRIES     int    `json:"MAX_RETRIES"`
	RETRY_COST      uint   `json:"RETRY_COST"`
}

func (awsConfig *AWSConfigModel) ToJSONString() (string, error) {
	if awsConfigString, marshalErr := json.Marshal(awsConfig); marshalErr == nil {
		return string(awsConfigString), nil
	} else {
		customErr := utilService.CreateCustomError(marshalErr, enums.CONV_ERROR, "500", enums.DEFAULT_LAYOUT, "Cannot convert aws config model to string", enums.INTERNAL_ERROR)
		return "", customErr
	}
}

func (awsConfig *AWSConfigModel) ToConfigString() string {
	str, _ := awsConfig.ToJSONString()
	return str
}
