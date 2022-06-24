package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	apiService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/api/location/services"
	utilService "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	cerr "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/interfaces"
	serverUtils "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/server-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/services"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/utils"
)

func TestHttpHandlerFunc(ctx context.Context, w *http.ResponseWriter, r http.Request) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func main() {
	fmt.Println("Echoing...")
	//loading the log config
	loadLogConfig()
	// load the env -  TODO - best to do it in the build step , lets see what can be done about that
	loadEnv()
	// load the appconfiguration
	appConfig := getAppConfiguration().(*models.AppConfigModel)
	zlog.Info().Msg("INfo:" + appConfig.ToConfigString())

	app := fiber.New(*serverUtils.GetAppFiberConfig(appConfig))

	app.Server().MaxConnsPerIP = 2

	app.Route("api", func(router fiber.Router) {
		// Hooks
		serverUtils.ConfigureHooks(app)
		// REGISTER MIDDLEWARES
		api := app.Group("/api")

		serverUtils.ConfigMiddlewares(app, appConfig)

		// registering routes
		apiService.RegisterRoutes(api)

	})

	// TODO: There should be validators, we can work on that later

	//route handlers

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		serverUtils.GracefullShutDown(app, signalCh)
	}()

	httpErr := app.Listen(":3001")
	if httpErr != nil {
		customErr := utilService.CreateCustomError(httpErr, enums.SERVER_ERROR, "500", enums.DEFAULT_LAYOUT, "Server cannot be started", enums.INTERNAL_ERROR)
		zlog.Panic().MsgFunc(func() string {
			convertedErr, _ := customErr.(*cerr.CustomErrModel)
			return convertedErr.ToString()
		})
	}

	// Loading the app & aws configuration with go routines and sync grp
	//var wg sync.WaitGroup
	//configChannel := make(chan interfaces.IConfig)

	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		configChannel <- getAppConfiguration()
	// 	}()
	// }

	// go func() {
	// 	log.Info().Msg("Befre calling waitgrp done")
	// 	wg.Wait()
	// 	close(configChannel)
	// 	log.Info().Msg("After calling waitgrp done")
	// }()

	// for channelVal := range configChannel {
	// 	fmt.Printf("configChannel Val: %v\n", channelVal.(*models.AppConfigModel))
	// }

}

func getAppConfiguration() interfaces.IConfig {
	appConfigPassed, _ := services.GetAppConfig()
	return appConfigPassed

}

func loadEnv() {
	env := utilService.GetEnvWithFallback("APP_GO_ENV", "dev").(string)
	if env == strings.ToLower("CCDEV") {
		// load from environemnt
		getCurrentDir, _ := os.Getwd()
		if err := godotenv.Load(path.Join(getCurrentDir, "src/config/development/.env")); err != nil {
			customErr := utilService.CreateCustomError(err, enums.FILE_ERROR, "500", enums.DEFAULT_LAYOUT, err.Error(), enums.INTERNAL_ERROR)
			if convertedCustomErrModel, ok := customErr.(*cerr.CustomErrModel); ok {
				zlog.Panic().Msgf("Code: %v; Status: %v ;Timestamp: %v ErrorMessage: %v", convertedCustomErrModel.Code, convertedCustomErrModel.Status, convertedCustomErrModel.Timestamp, convertedCustomErrModel.Message)
			} else {
				zlog.Panic().Msgf("Panic: %v", customErr)
			}

		}
	} else {
		// Get the environment from aws - to load something from local system - aws profiile using config.WithSharedConfigProfile

		ssmParamPath := utils.AWS_SSM_PARAM_PATH_PREFIX + "dev" + utils.SLASH
		zlog.Info().Msgf("Param path: %v", ssmParamPath)
		// Configured the retryable configuration for the aws resource/service call
		// if its local, we need to have a defined "personal" (you can set ur own name if so use value from env) aws profile and we need to pick it from there
		// if its within aws itself, no need to create and use any aws profile

		awsDefinedConfig, awsErr := services.CreateAwsConfig(env)

		if awsErr == nil {
			// get the values from aws ssm secret service
			services.GetSSMParamStoreDefaults(context.Background(), awsDefinedConfig, ssmParamPath)
		}
	}

}

func loadLogConfig() {
	zerolog.TimeFieldFormat = zerolog.TimestampFunc().Format(string(enums.DEFAULT_LAYOUT))
	if isDebug, flag := os.LookupEnv("APP_DEBUG_ENABLED"); flag {
		if strings.EqualFold(isDebug, "true") {
			// then debug is enabled
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		customErr := utilService.CreateCustomError(errors.New("test err"), enums.FILE_ERROR, "500", enums.DEFAULT_LAYOUT, errors.New("test err").Error(), enums.INTERNAL_ERROR)
		// Testing zlog - TODO:REMOVE this
		if convertedCustomErrModel, ok := customErr.(*cerr.CustomErrModel); ok {
			zlog.Info().Msgf("Code: %v; Status: %v ;Timestamp: %v ErrorMessage: %v", convertedCustomErrModel.Code, convertedCustomErrModel.Status, convertedCustomErrModel.Timestamp, convertedCustomErrModel.Message)
		} else {
			zlog.Info().Str("test", customErr.Error()).Send()
		}
	}
}
