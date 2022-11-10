package main

import (
	"backend/internal/adapter/persistence"
	"backend/internal/handler"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"os"
)

func handle(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: 500}, err
	}
	repo := persistence.NewRatingRepo(dynamodb.NewFromConfig(cfg), os.Getenv("TABLE_NAME"))

	return handler.NewRatingHandler(repo).Handle(ctx, event)
}

func main() {
	lambda.Start(handle)
}
