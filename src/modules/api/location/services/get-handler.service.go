package location

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go/middleware"
	"github.com/gofiber/fiber/v2"
	zlog "github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
	apiMiddleware "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/api/location/middlewares"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/api/location/models/dynamodb/response"
	apiResp "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/api/location/models/dynamodb/response"
	appUtils "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	cmodels "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	serverUtils "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/server-utils"
	cservices "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/services"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/utils"
)

func GetAllLocations(ctx *fiber.Ctx) error {
	//TODO: connect to dynamodb and get all locations
	env := appUtils.GetEnvWithFallback("APP_GO_ENV", "dev").(string)
	if awsDefinedConfig, awsErr := cservices.CreateAwsConfig(env); awsErr == nil {

		printDynamoAction := middleware.InitializeMiddlewareFunc("PrintIngress", apiMiddleware.InitialzierMiddleware)
		awsDefinedConfig.APIOptions = append(awsDefinedConfig.APIOptions, func(stack *middleware.Stack) error {
			// Attach the custom middleware to the beginning of the Initialize step
			return stack.Initialize.Add(printDynamoAction, middleware.Before)
		})

		// MOST IMP!!!! Every handler implementation needs to have this , this is the only way we can recover in case of panics
		defer func() {
			if r := recover(); r != nil {
				zlog.Info().Msgf("Code: %v; Status: %v ;Timestamp: %v Message: Recovered from panic in method - %v", 200, 200, appUtils.GetTimeStamp(enums.DEFAULT_LAYOUT), "GetAllLocations")
			}
		}()
		// dynamodb specific configs
		opt := dynamodb.EndpointDiscoveryOptions{
			EnableEndpointDiscovery: aws.EndpointDiscoveryAuto, // THis takes care of endpoint discovery for streams
		}
		// Its an array of option functions
		// TODO: best
		dynamoDBClient := dynamodb.NewFromConfig(*awsDefinedConfig, func(o *dynamodb.Options) {
			o.EndpointDiscovery = opt
		}, func(o *dynamodb.Options) {
			o.Retryer = serverUtils.GetRetryFunctionalityForAll()
		})

		// First make sure that the table is available
		var tbName string
		if ctx.Locals(utils.CONTEXT_VARS.APPCONFIG).(cmodels.AppConfigModel).AppName != "" {
			tbName = (ctx.Locals(utils.CONTEXT_VARS.APPCONFIG).(cmodels.AppConfigModel)).AWSConfig.DYNAMO_DB_TABLE
		} else {
			tbName = ""
		}
		if op, err := dynamoDBClient.DescribeTable(context.Background(), &dynamodb.DescribeTableInput{
			TableName: &tbName,
		}); err == nil {
			tableModel := apiResp.DescribeTableModel(*op)

			// means the table is available in dynamodb and now list all items
			if scannedItemMap, scannedErr := dynamoDBClient.Scan(context.Background(), &dynamodb.ScanInput{
				TableName:      tableModel.Table.TableName,
				ConsistentRead: aws.Bool(true),
				Limit:          aws.Int32(1500),
			}); scannedErr == nil {
				commonRespModel := response.CommonResponse{}
				commonRespModel.StatusMsg = "Success"
				commonRespModel.Error = nil
				commonRespModel.Metadata = scannedItemMap.ResultMetadata
				var records []string

				for _, itemMap := range scannedItemMap.Items {
					for key, val := range itemMap {
						marshalledVal, _ := json.Marshal(val)
						fmt.Println("marshalledVal", string(marshalledVal))
						// By default this has a json string like {"Value": xxxx} -  so decoding  it
						resultVal := gjson.Get(string(marshalledVal), "Value")
						fmt.Println("vv:", resultVal)
						resultItem := map[string]interface{}{
							key: resultVal.Str,
						}

						record, _ := json.Marshal(resultItem)
						fmt.Println("ree:", string(record))
						records = append(records, string(record))

					}

				}
				commonRespModel.Data = strings.Join(records, ",")
				respStr, _ := commonRespModel.ToJSONString()
				return ctx.Status(200).Send([]byte(respStr))
			} else {
				fmt.Println("err:", scannedErr)
				return ctx.Status(500).Send([]byte(scannedErr.Error()))
			}

		} else {
			fmt.Println("error occurs", err)
			return ctx.Status(500).Send([]byte(err.Error()))
		}

	} else {
		// TODO: create a better log message and then panic
		return ctx.Status(500).Send([]byte(awsErr.Error()))
	}
	// First get the awsConfig and load it

}

func GetLocationByCity(ctx *fiber.Ctx) {

}

func GetLocationByCityAndCounty(ctx *fiber.Ctx) {

}

func GetLocationByCityCountyAndCountry(ctx *fiber.Ctx) {

}

func GetLocationByCityPinCode(ctx *fiber.Ctx) {

}
