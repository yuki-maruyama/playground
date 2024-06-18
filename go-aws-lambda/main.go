package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	String string `json:"string"`
}

func LambdaHandler(ctx context.Context, event *myEvent) (string, error) {
	fmt.Println("function invoke")
	var err error
	fmt.Printf("string: %s\n", event.String)
	return "", err
}

func main() {
	lambda.Start(LambdaHandler)
}
