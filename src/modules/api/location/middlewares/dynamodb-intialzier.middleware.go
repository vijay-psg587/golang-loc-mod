package middlewares

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go/middleware"
)

func InitialzierMiddleware(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (out middleware.InitializeOutput, metadata middleware.Metadata, err error) {

	// TODO: THere are advaned use cases for using middlewares here

	fmt.Println("Just printing it in middleware", in.Parameters, out.Result)

	// Middleware must call the next middleware to be executed in order to continue execution of the stack.
	// If an error occurs, you can return to prevent further execution.
	return next.HandleInitialize(ctx, in)
}

// This is the common intialzie middleware that can be attached to all dynamoDB clients
// For now this just logs the input and output
func AttachCommonInitializerMiddleware(awsConfig *aws.Config) {

	initializeMiddleware := middleware.InitializeMiddlewareFunc("Common Initializer Middleware", InitialzierMiddleware)
	awsConfig.APIOptions = append(awsConfig.APIOptions, func(s *middleware.Stack) error {
		return s.Initialize.Add(initializeMiddleware, middleware.After)
	})
}

// This is to get the dynamoDb configuration
func GetDyanmoDbClient(awsConfig *aws.Config, optFns ...func(*dynamodb.Options)) *dynamodb.Client {
	//Configure the req middlewares first
	AttachCommonInitializerMiddleware(awsConfig)
	if optFns == nil {
		return dynamodb.NewFromConfig(*awsConfig)
	} else {
		return dynamodb.NewFromConfig(*awsConfig, optFns...)
	}

}
