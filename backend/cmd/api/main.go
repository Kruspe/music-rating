package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/kruspe/music-rating/internal/config"
)

//nolint:gochecknoglobals // lambda start up optimization
var mux *http.ServeMux

//nolint:gochecknoinits // lambda start up optimization
func init() {
	config.InitLogging()
	cfg := config.InitAwsConfig()
	repos := config.InitRepos(cfg)
	storage := config.InitStorage(cfg)
	useCases := config.InitUseCases(repos, storage)
	mux = config.InitApi(useCases, repos)

}

func main() {
	lambda.Start(httpadapter.NewV2(mux).ProxyWithContext)
}
