package config

import (
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/usecase"
)

func InitUseCases(repos *persistence.Repositories, festivalStorage *persistence.FestivalStorage) *usecase.UseCases {
	return usecase.NewUseCases(repos, festivalStorage)
}
