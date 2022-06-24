package request

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	zlog "github.com/rs/zerolog/log"
	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	cerr "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
)

type ScanInputRequest struct {
	Input dynamodb.ScanInput
}

func (dt *ScanInputRequest) ToJSONString() (string, error) {
	if dtModelByte, err := json.Marshal(dt); err == nil {
		return string(dtModelByte), nil
	} else {
		customErr := utilService.CreateCustomError(err, enums.CONV_ERROR, "500", enums.DEFAULT_LAYOUT, err.Error(), enums.INTERNAL_ERROR)
		// This Panics and stops the app -  This is the reason where Recover was inclued in main.go to recover from pancis if any
		convertedCustomErrModel, _ := customErr.(*cerr.CustomErrModel)
		zlog.Error().Msgf("Code: %v; Status: %v ;Timestamp: %v ErrorMessage: %v", convertedCustomErrModel.Code, convertedCustomErrModel.Status, convertedCustomErrModel.Timestamp, convertedCustomErrModel.Message)
		return "", customErr
	}
}
