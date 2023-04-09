package config

import (
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/api"
	"github.com/kruspe/music-rating/internal/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func InitApi(useCases *usecase.UseCases, repos *persistence.Repositories, log *log.Logger) *httpadapter.HandlerAdapterV2 {
	mux := http.NewServeMux()
	mux.Handle("/", api.NewApi(useCases, repos, api.NewErrorHandler(log)))

	return httpadapter.NewV2(mux)
}
