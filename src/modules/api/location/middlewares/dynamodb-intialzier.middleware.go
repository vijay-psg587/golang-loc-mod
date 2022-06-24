package middlewares

import (
	"context"
	"fmt"

	"github.com/aws/smithy-go/middleware"
)

func InitialzierMiddleware(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (out middleware.InitializeOutput, metadata middleware.Metadata, err error) {

	// TODO: THere are advaned use cases for using middlewares here

	fmt.Println("Just printing it in middleware", in)

	// Middleware must call the next middleware to be executed in order to continue execution of the stack.
	// If an error occurs, you can return to prevent further execution.
	return next.HandleInitialize(ctx, in)
}
