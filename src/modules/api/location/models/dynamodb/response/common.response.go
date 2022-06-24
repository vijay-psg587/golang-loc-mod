package response

import (
	"encoding/json"
	"strconv"

	zlog "github.com/rs/zerolog/log"
	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	cerr "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
)

type CommonResponse struct {
	StatusMsg string
	Error     error
	Metadata  interface{}
	Data      interface{}
}

func (dt *CommonResponse) ToJSONString() (string, error) {
	if _, err := json.Marshal(dt); err == nil {

		if convertedStr, ok := dt.Data.(string); ok {
			if unquoteStr, err := strconv.Unquote(convertedStr); err != nil {
				dt.Data = convertedStr + "\n"
			} else {
				dt.Data = unquoteStr + "\n"
			}

		}
		marshalledErr, _ := json.Marshal(dt.Error)
		marshalledStatus, _ := json.Marshal(dt.StatusMsg)
		marshalledMetadata, _ := json.Marshal(dt.Metadata)
		return "{\n" + "Error:" + string(marshalledErr) + "," + "Status:" + string(marshalledStatus) + "," + "Metadata:" + string(marshalledMetadata) + "," + "Data:" + dt.Data.(string) + "}", nil
	} else {
		customErr := utilService.CreateCustomError(err, enums.CONV_ERROR, "500", enums.DEFAULT_LAYOUT, err.Error(), enums.INTERNAL_ERROR)
		// This Panics and stops the app -  This is the reason where Recover was inclued in main.go to recover from pancis if any
		convertedCustomErrModel, _ := customErr.(*cerr.CustomErrModel)
		zlog.Error().Msgf("Code: %v; Status: %v ;Timestamp: %v ErrorMessage: %v", convertedCustomErrModel.Code, convertedCustomErrModel.Status, convertedCustomErrModel.Timestamp, convertedCustomErrModel.Message)
		return "", customErr
	}
}
