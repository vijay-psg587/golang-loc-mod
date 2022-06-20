package aws_utils

import (
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models"
)

func CreateCustomRetry(appConfig models.AppConfigModel) *retry.Standard {

	return retry.NewStandard(func(so *retry.StandardOptions) {
		so.MaxAttempts = 2
	})
}
