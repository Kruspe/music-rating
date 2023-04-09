package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kruspe/music-rating/internal/config"
)

func main() {
	log := config.InitLogging()
	cfg := config.InitAwsConfig(log)
	repos := config.InitRepos(cfg, log)
	storage := config.InitStorage(cfg)
	useCases := config.InitUseCases(repos, storage)
	api := config.InitApi(useCases, repos, log)

	lambda.Start(api.ProxyWithContext)
}
