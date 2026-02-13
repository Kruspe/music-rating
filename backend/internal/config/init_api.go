package config

import (
	"net/http"

	"github.com/kruspe/music-rating/internal/handler"
	"github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/usecase"
)

func InitApi(useCases *usecase.UseCases, repos *persistence.Repositories) *http.ServeMux {
	mux := http.NewServeMux()
	handler.Register(mux, &handler.Config{
		RatingRepo:      repos.RatingRepo,
		FestivalUseCase: useCases.FestivalUseCase,
	})

	return mux
}
