package models

import (
	"encoding/json"
	"time"

	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
)

type HTTPClientSettings struct {
	Connect          time.Duration
	ConnKeepAlive    time.Duration
	ExpectContinue   time.Duration
	IdleConn         time.Duration
	MaxAllIdleConns  int
	MaxHostIdleConns int
	ResponseHeader   time.Duration
	TLSHandshake     time.Duration
}

func (httpClientConfig *HTTPClientSettings) ToJSONString() (string, error) {
	if httpClientConfigStr, marshalErr := json.Marshal(httpClientConfig); marshalErr == nil {
		return string(httpClientConfigStr), nil
	} else {
		customErr := utilService.CreateCustomError(marshalErr, enums.CONV_ERROR, "500", enums.DEFAULT_LAYOUT, "Cannot convert HttpClient settings to string", enums.INTERNAL_ERROR)
		return "", customErr
	}
}

func (httpClientConfig *HTTPClientSettings) ToConfigString() string {
	str, _ := httpClientConfig.ToJSONString()
	return str
}
