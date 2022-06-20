package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	zlog "github.com/rs/zerolog/log"
	appUtils "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	cerr "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/utils"
)

func GetSSMParamStoreDefaults(ctx context.Context, config *aws.Config, paramPath string) {
	client := ssm.NewFromConfig(*config)
	//strArr := [2]string{"DYNAMO_TABLE"}

	ssmParamOp, err := client.GetParametersByPath(ctx, &ssm.GetParametersByPathInput{
		Path:           &paramPath,
		MaxResults:     10,
		Recursive:      true,
		WithDecryption: true,
	})
	if err != nil {
		// Create the custom error, and throw the error
		customErr := appUtils.CreateCustomError(err, enums.AWS_ERROR, "500", enums.DEFAULT_LAYOUT, err.Error(), enums.AWS_INTERNAL_ERROR)
		if convertedCustomErrModel, ok := customErr.(*cerr.CustomErrModel); ok {
			zlog.Panic().Msgf("Code: %v; Status: %v ;Timestamp: %v ErrorMessage: %v", convertedCustomErrModel.Code, convertedCustomErrModel.Status, convertedCustomErrModel.Timestamp, convertedCustomErrModel.Message)
		}
	} else {
		// Fetch the results to be populated in env

		for _, val := range ssmParamOp.Parameters {
			strKeyFull := strings.Split(*val.Name, utils.SLASH)
			fmt.Println("val:", strKeyFull[len(strKeyFull)-1], *val.Value)
			os.Setenv(strKeyFull[len(strKeyFull)-1], *val.Value)
		}

		// Modify the keys to
	}
}
