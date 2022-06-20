package server_utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	zlog "github.com/rs/zerolog/log"
	appUtils "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	cerr "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/utils"
)

//"github.com/rs/zerolog/log"

// func FiberErrorHandler(ctx fiber.Ctx, customError error) fiber.ErrorHandler {
// 	return func(c *fiber.Ctx, err error) error {
// 		log.Error().Msg("Inside fiber Error Handler")
// 		if reflect.TypeOf(err).Name() == "CustomErrorModel" {
// 			ctx.GetRes()
// 		}
// 	}
// }

func ConfigureHooks(app *fiber.App) {

	// shutdown hook
	app.Hooks().OnShutdown(func() error {
		fmt.Println("Inside post shutdown")
		// TODO: Any post shutdown activities to be carried out, like closing any monitoring & observability agents
		// Need to test this more
		return errors.New("None")
	})
}

func ConfigMiddlewares(app *fiber.App) {

	var wg sync.WaitGroup
	// Using cache middleware
	cacheStatusCh := make(chan string, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		cacheStatusCh <- RegisterCacheMiddleware(app)
	}()

	// Using cors middleware
	wg.Add(1)
	go func() {
		defer wg.Done()
		cacheStatusCh <- RegisterCorsMiddleware(app)
	}()
	// using requestId middleware
	wg.Add(1)
	go func() {
		defer wg.Done()
		cacheStatusCh <- RegisterReqIdMiddleware(app)
	}()

	//custom middleware for logger
	wg.Add(1)
	go func() {
		defer wg.Done()
		cacheStatusCh <- RegisterLoggerMiddleware(app)
	}()

	// TODO: Implement  custom authentication middleware - possibly try the paigw imp
	go func() {
		wg.Wait()
		close(cacheStatusCh)
	}()

	for val := range cacheStatusCh {

		fmt.Println("DDDD", val)
		zlog.Info().Msgf("Code: %v; Status: %v ;Timestamp: %v; Message: %v - Completed", 200, "200", appUtils.GetTimeStamp(enums.DEFAULT_LAYOUT), val)
	}
}

func RegisterReqIdMiddleware(app *fiber.App) string {
	app.Use(requestid.New(requestid.Config{
		ContextKey: app.Config().AppName + "requestid",
	}))
	return "RegisterReqId"
}

func RegisterLoggerMiddleware(app *fiber.App) string {
	app.Use(func(ctx *fiber.Ctx) error {
		host := string(ctx.Request().Host())
		zlog.Info().Msgf("Req Host:%v", host)

		headers := ctx.GetReqHeaders()
		var strVal string
		for k, v := range headers {
			if strVal == "" {
				strVal = "HeaderKey:" + k + utils.COLON_CHAR + "HeaderVal:" + v + "\n"
			} else {
				strVal = strVal + "HeaderKey:" + k + utils.COLON_CHAR + "HeaderVal:" + v + "\n"
			}

		}
		zlog.Info().Msgf("Req Headers:%v", strVal)

		zlog.Info().Msgf("Req path:%v", ctx.Path())
		return ctx.Next()
	})
	return "RegisterLoggerMiddleware"
}

func RegisterCacheMiddleware(app *fiber.App) string {
	app.Use(cache.New(cache.Config{
		ExpirationGenerator: func(c1 *fiber.Ctx, c2 *cache.Config) time.Duration {
			expiration, _ := strconv.Atoi(c1.GetRespHeader("Cache-Time", "500"))
			c2.Expiration = time.Duration(expiration)
			return c2.Expiration
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			if c.Route().Method == http.MethodGet {
				return c.Route().Path
			} else {
				return ""
			}
		},
	}))

	return "RegisterCacheMiddleware"
}

func RegisterCorsMiddleware(app *fiber.App) string {
	allowCred, _ := appUtils.GetEnvWithFallback("ALLOW_CREDS", "false").(bool)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     strings.Join(utils.CORS.ACCEPT_METHODS, utils.COMMA_CHAR),
		AllowCredentials: allowCred,
		AllowHeaders:     strings.Join(utils.CORS.ACCEPT_HEADERS, utils.COMMA_CHAR),
	}))
	return "RegisterCorsMiddleware"
}

func GracefullShutDown(app *fiber.App, ch chan os.Signal) {
	// TODO: Shutdown or close any db connections opened
	// TODO: Close any streams opened
	<-ch
	fmt.Println("Gracefully Shutting down...")
	_ = app.Shutdown()

}

func GetRetryFunctionalityForAll() aws.Retryer {
	r := retry.AddWithErrorCodes(retry.NewStandard(), "424", "433")
	retry.AddWithMaxAttempts(r, 5)                                // TODO - get from env
	retry.AddWithMaxBackoffDelay(r, time.Duration(time.Second*2)) // TODO: get this from env
	return r
}

func GetAppFiberConfig(appConfig *models.AppConfigModel) *fiber.Config {
	return &fiber.Config{
		CaseSensitive:        true,
		UnescapePath:         true,
		ErrorHandler:         FiberCustomErrorHandler,
		CompressedFileSuffix: appConfig.AppName + ".gz",
		ServerHeader:         appConfig.AppName + utils.HYPEN_CHAR + "server",
		StreamRequestBody:    true,
		AppName:              appConfig.AppName,
		ReadTimeout:          time.Second * time.Duration(appConfig.AppServerConfig.ReadTimeout),
		WriteTimeout:         time.Second * time.Duration(appConfig.AppServerConfig.WriteTimeout),
		IdleTimeout:          time.Second * time.Duration(appConfig.AppServerConfig.IdleTimeout),
	}
}

func FiberCustomErrorHandler(ctx *fiber.Ctx, err error) error {

	if e, ok := err.(*fiber.Error); ok {
		//Its a fiber Error so continue sending error response
		// TODO: create a seperate error message for template hosting
		// Identifying the correspodning ErrorCodeNum
		errorEnumCodes := []enums.ErrorCodeEnum{enums.ACCESS_ERROR, enums.AWS_ERROR, enums.CONSTRAINT_ERROR, enums.CONV_ERROR, enums.DB_ERROR, enums.FILE_ERROR, enums.INTERNAL, enums.REDIS_ERROR, enums.REQ_VALIDATION, enums.SERVER_ERROR}
		var code enums.ErrorCodeEnum
		for errorEnumVal := range errorEnumCodes {
			if e.Code == int(errorEnumVal) {
				code = errorEnumCodes[errorEnumVal]
			} else {
				code = enums.INTERNAL
			}
		}

		customErr := appUtils.CreateCustomError(e, code, "500", enums.DEFAULT_LAYOUT, e.Message, enums.INTERNAL_ERROR)
		convertedCustomErrModel, _ := customErr.(*cerr.CustomErrModel)
		zlog.Error().Msgf("Code: %v; Status: %v ;Timestamp: %v ErrorMessage: %v", convertedCustomErrModel.Code, convertedCustomErrModel.Status, convertedCustomErrModel.Timestamp, convertedCustomErrModel.Message)

		return ctx.Send([]byte(convertedCustomErrModel.ToString()))
	} else if cErr, ok := err.(*cerr.CustomErrModel); ok {
		return ctx.Send([]byte(cErr.ToString()))
	} else {

		customErr := appUtils.CreateCustomError(err, enums.SERVER_ERROR, "500", enums.DEFAULT_LAYOUT, err.Error(), enums.INTERNAL_ERROR)
		convertedCustomErrModel, _ := customErr.(*cerr.CustomErrModel)
		return ctx.Status(500).Send([]byte(convertedCustomErrModel.ToString()))
	}

}
