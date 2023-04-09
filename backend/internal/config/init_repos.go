package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitRepos(cfg aws.Config, log *log.Logger) *persistence.Repositories {
	tableName, present := os.LookupEnv("TABLE_NAME")
	if !present {
		err := fmt.Errorf("missing table name in environment variables")
		log.Fatal(err)
		panic(err)
	}
	dynamo := dynamodb.NewFromConfig(cfg)

	return persistence.NewRepositories(dynamo, tableName)
}
