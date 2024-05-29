package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"log/slog"
	"os"
)

func InitRepos(cfg aws.Config) *persistence.Repositories {
	tableName, present := os.LookupEnv("TABLE_NAME")
	if !present {
		err := fmt.Errorf("missing table name in environment variables")
		slog.Error("missing env variable", slog.Any("error", err))
		panic(err)
	}
	dynamo := dynamodb.NewFromConfig(cfg)

	return persistence.NewRepositories(dynamo, tableName)
}
