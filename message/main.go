package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/daesu/serverless-go/message/handler"
)

func main() {
	lambda.Start(handler.Handler)
}
