package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kruspe/music-rating/internal/config"
)

func main() {
	config.InitLogging()
	cfg := config.InitAwsConfig()
	repos := config.InitRepos(cfg)
	storage := config.InitStorage(cfg)
	useCases := config.InitUseCases(repos, storage)
	api := config.InitApi(useCases, repos)

	lambda.Start(api.ProxyWithContext)
}
