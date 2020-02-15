package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Print("Got request for '/.netlify/functions/goodbye', this message is dumpled by 'cmd/goodbye/main.go'")
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, I'm from 'cmd/goodbye/main.go'!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
