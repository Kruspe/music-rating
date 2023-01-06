package main

import (
	"backend/internal/adapter/persistence"
	"backend/internal/handler"
	"backend/internal/usecase"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/sirupsen/logrus"
	"os"
)

func handle(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	logger := log.New()
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatal("Could not get log level from environment variable")
	}
	logger.SetLevel(level)
	logger.SetFormatter(&log.JSONFormatter{})

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: 500}, err
	}
	repo := persistence.NewRatingRepo(dynamodb.NewFromConfig(cfg), os.Getenv("TABLE_NAME"))
	festivalStorage := persistence.NewFestivalStorage(s3.NewFromConfig(cfg))

	ratingUseCase := usecase.NewRatingUseCase(repo, festivalStorage)

	return handler.NewRatingHandler(ratingUseCase, logger).Handle(ctx, event)
}

func main() {
	lambda.Start(handle)
}
