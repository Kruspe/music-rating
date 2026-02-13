package handler

import (
	"net/http"

	"github.com/kruspe/music-rating/internal/handler/festival"
	"github.com/kruspe/music-rating/internal/handler/rating"
	"github.com/kruspe/music-rating/internal/middleware"
	"github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/usecase"
)

type Config struct {
	RatingRepo      *persistence.RatingRepo
	FestivalUseCase *usecase.FestivalUseCase
}

type handler struct {
	mux    *http.ServeMux
	config *Config
}

func Register(mux *http.ServeMux, config *Config) {
	h := &handler{
		mux:    mux,
		config: config,
	}

	h.registerRating()
	h.registerFestival()
}

func (h *handler) registerRating() {
	e := rating.NewRatingEndpoint(h.config.RatingRepo, h.config.FestivalUseCase)

	h.mux.Handle("POST /ratings", middleware.AuthMiddleware(e.Create()))
	h.mux.Handle("GET /ratings", middleware.AuthMiddleware(e.GetAll()))
	h.mux.Handle("PUT /ratings/{artistName}", middleware.AuthMiddleware(e.Put()))
	h.mux.Handle("GET /ratings/{festivalName}", middleware.AuthMiddleware(e.GetAllForFestival()))
}

func (h *handler) registerFestival() {
	e := festival.NewFestivalEndpoint(h.config.FestivalUseCase)

	h.mux.Handle("GET /festivals/{festivalName}", middleware.AuthMiddleware(e.GetArtistsForFestival()))
}
