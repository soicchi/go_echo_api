package main

import (
	"log"

	"go_echo_api/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	log.Println("Starting Lambda")

	e := routes.SetupRoutes()
	echoLambda = echoadapter.New(e)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
