package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithBaseEndpoint("http://localhost:4566"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := lambda.NewFromConfig(cfg)

	payload, _ := json.Marshal(map[string]string{"name": "Rafael"})

	result, err := client.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName: aws.String("hello-lambda"),
		Payload:      payload,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %d\nResposta: %s\n", result.StatusCode, string(result.Payload))
}
