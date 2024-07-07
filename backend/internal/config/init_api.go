package config

import (
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/kruspe/music-rating/internal/api"
	"github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/usecase"
	"net/http"
)

func InitApi(useCases *usecase.UseCases, repos *persistence.Repositories) *httpadapter.HandlerAdapterV2 {
	ratingEndpoint := api.NewRatingEndpoint(repos.RatingRepo, useCases.FestivalUseCase)
	festivalEndpoint := api.NewFestivalEndpoint(useCases.FestivalUseCase)
	mux := http.NewServeMux()
	mux.Handle("/", api.AuthMiddleware(api.NewRouter(festivalEndpoint, ratingEndpoint)))

	return httpadapter.NewV2(mux)
}
