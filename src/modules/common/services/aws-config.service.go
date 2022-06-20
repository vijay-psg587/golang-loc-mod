package services

import (
	"context"
	"errors"
	"sync"

	zlog "github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	serverUtils "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/server-utils"
)

var (
	awsDefinedConfig aws.Config = aws.Config{}
	awsErr           error      = nil
)

func CreateAwsConfig(env string) (*aws.Config, error) {

	var m sync.Mutex
	m.Lock()
	defer m.Unlock()
	if env == "dev_local" {
		awsDefinedConfig, awsErr = config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile("personal"),
			config.WithRetryer(serverUtils.GetRetryFunctionalityForAll))
	} else {
		awsDefinedConfig, awsErr = config.LoadDefaultConfig(context.Background(), config.WithRetryer(serverUtils.GetRetryFunctionalityForAll))
	}
	if awsErr == nil {
		return &awsDefinedConfig, nil
	} else {
		zlog.Panic().Msg("Cannot initialize aws configuration")
		return &aws.Config{}, errors.New("cannot initialize aws configuration")
	}
}
