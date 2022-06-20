package models

import (
	"encoding/json"

	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
)

type AppServerConfigModel struct {
	Port         int
	Addr         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

func (appServerConfig *AppServerConfigModel) ToJSONString() (string, error) {
	if appConfigStr, marshalErr := json.Marshal(appServerConfig); marshalErr == nil {
		return string(appConfigStr), nil
	} else {
		customErr := utilService.CreateCustomError(marshalErr, enums.CONV_ERROR, "500", enums.DEFAULT_LAYOUT, "Cannot convert app config model to string", enums.INTERNAL_ERROR)
		return "", customErr
	}
}

func (appServerConfig *AppServerConfigModel) ToConfigString() string {
	str, _ := appServerConfig.ToJSONString()
	return str
}
