package models

import (
	"encoding/json"

	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
)

type AppConfigModel struct {
	AppName          string
	AppServerConfig  AppServerConfigModel
	AWSConfig        AWSConfigModel
	HTTPClientConfig HTTPClientSettings
}

func (appConfig *AppConfigModel) ToJSONString() (string, error) {
	if appConfigStr, marshalErr := json.Marshal(appConfig); marshalErr == nil {
		return string(appConfigStr), nil
	} else {
		customErr := utilService.CreateCustomError(marshalErr, enums.CONV_ERROR, "500", enums.DEFAULT_LAYOUT, "Cannot convert app config model to string", enums.INTERNAL_ERROR)
		return "", customErr
	}
}

func (appConfig *AppConfigModel) ToConfigString() string {
	str, _ := appConfig.ToJSONString()
	return str
}
