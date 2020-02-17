package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandlerReturnsMessage(t *testing.T) {
	message, _ := handler(events.APIGatewayProxyRequest{})

	if message.StatusCode != 200 {
		t.Errorf("Unintentional status code: %d", message.StatusCode)
	}

	if message.Body != "Hello, I'm from 'cmd/goodbye/main.go'!" {
		t.Errorf("Unintentional message: %s", message.Body)
	}
}
