package main

import (
	"log"
	"os"

	"go_echo_api/database"
	"go_echo_api/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	log.Println("Starting Lambda")

	// Database connection
	dbConfig := database.NewDBConfig(
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	dsn := dbConfig.CreateDSN()
	_, err := database.DBConnect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")

	e := routes.SetupRoutes()
	echoLambda = echoadapter.New(e)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
