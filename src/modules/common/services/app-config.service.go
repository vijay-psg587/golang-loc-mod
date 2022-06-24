package services

import (
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/utils"
)

// Declare global variables to be used with singleton pattern
var (
	appConfig       *models.AppConfigModel       = new(models.AppConfigModel)
	appServerConfig *models.AppServerConfigModel = new(models.AppServerConfigModel)
	lock                                         = new(sync.Mutex)
	awsConfig       *models.AWSConfigModel       = new(models.AWSConfigModel)
	clientConfig    *models.HTTPClientSettings   = new(models.HTTPClientSettings)
)

// Getting the AppConfig but setting values only one time, irrespective of the number of calls
// using mutex
func GetAppConfig() (*models.AppConfigModel, error) {
	if appConfig.AppName == "" {
		lock.Lock()
		defer lock.Unlock()
		appConfig.AppName = utilService.GetEnvWithFallback("APP_NAME", utils.APP_NAME).(string)
		// getting aws config
		if awsConfig, awsConfigErr := GetAwsAppConfig(); awsConfigErr == nil {
			appConfig.AWSConfig = *awsConfig
		} else {
			return &models.AppConfigModel{}, awsConfigErr
		}

		// get aws server config
		if appServerConfig, appServerConfigErr := GetAppServerConfig(); appServerConfigErr == nil {
			appConfig.AppServerConfig = *appServerConfig
		} else {
			return &models.AppConfigModel{}, appServerConfigErr
		}

		httpClientSettings := &models.HTTPClientSettings{}
		if connKeepAlive, convErr := strconv.Atoi(utilService.GetEnvWithFallback("HTTP_CONN_KEEP_ALIVE", "25000").(string)); convErr == nil {
			httpClientSettings.ConnKeepAlive = time.Duration(time.Duration(connKeepAlive).Seconds())
			return appConfig, nil
		} else {
			csErr := utilService.CreateCustomError(convErr, enums.CONV_ERROR, "500", enums.DEFAULT_LAYOUT, convErr.Error(), enums.INTERNAL_ERROR)
			convertedCSErr, _ := csErr.(*errors.CustomErrModel)
			log.Error().Msgf("Error: Code:%v : Timestamp:%v, Message: %v", convertedCSErr.Code, convertedCSErr.Timestamp, convertedCSErr.Message)
			return &models.AppConfigModel{}, csErr
		}
	} else {
		return appConfig, nil
	}

}

func GetAppServerConfig() (*models.AppServerConfigModel, error) {
	if appServerConfig.Addr == "" {
		port, _ := strconv.Atoi(utilService.GetEnvWithFallback("PORT", "5000").(string))
		appServerConfig.Port = int(port)
		appServerConfig.Addr = utilService.GetEnvWithFallback("ADDR", "0.0.0.0").(string)
		readTimeOut, _ := strconv.Atoi(utilService.GetEnvWithFallback("READ_TIMEOUT", "15").(string))
		writeTimeOut, _ := strconv.Atoi(utilService.GetEnvWithFallback("WRITE_TIMEOUT", "15").(string))
		idleTimeOut, _ := strconv.Atoi(utilService.GetEnvWithFallback("IDLE_TIMEOUT", "15").(string))
		appServerConfig.IdleTimeout = int(idleTimeOut)
		appServerConfig.ReadTimeout = int(readTimeOut)
		appServerConfig.WriteTimeout = int(writeTimeOut)
	}
	return appServerConfig, nil
}

func GetAwsAppConfig() (*models.AWSConfigModel, error) {

	if awsConfig.DYNAMO_DB_TABLE == "" {
		awsConfig.DYNAMO_DB_TABLE = utilService.GetEnvWithFallback("DYNAMO_TABLE", "test_table").(string)
		awsConfig.Region = utilService.GetEnvWithFallback("AWS_REGION", "us-east-1").(string)

	}
	return awsConfig, nil
}

// TODO! - set HttpClientSettings
