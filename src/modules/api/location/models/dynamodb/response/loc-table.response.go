package response

import (
	"encoding/json"

	zlog "github.com/rs/zerolog/log"
	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	cerr "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
)

type LocTableSchema struct {
	B_LOC_COUNTRY string `json:b_loc_country; required:true`
	B_CENTER_ID   string `json:b_center_id; required:true`
}

func (dt *LocTableSchema) ToJSONString() (string, error) {
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
