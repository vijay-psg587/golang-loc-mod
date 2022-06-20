package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/utils"
)

func GetLocationValue() string {
	// TODO : Get values from dynamoDB
	return "location1"
}

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
		fmt.Println("Custom errir", err)
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
