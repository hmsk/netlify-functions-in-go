package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/[your org/account]/netlify-functions-in-go/internal/pkg/utils" can provide common utilities
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Print("Got request for '/.netlify/functions/hello', this message is dumpled by 'cmd/hello/main.go'")
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, I'm from 'cmd/hello/main.go'!",
		//Body:       utils.IntroductionYourself("cmd/goodbye/main.go"),
	}, nil
}

func main() {
	lambda.Start(handler)
}
