package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/kruspe/music-rating/internal/config"
)

var api *httpadapter.HandlerAdapterV2

func init() {
	config.InitLogging()
	cfg := config.InitAwsConfig()
	repos := config.InitRepos(cfg)
	storage := config.InitStorage(cfg)
	useCases := config.InitUseCases(repos, storage)
	api = config.InitApi(useCases, repos)

}

func main() {
	lambda.Start(api.ProxyWithContext)
}
