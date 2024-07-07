package config

import (
	persistence2 "github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/usecase"
)

func InitUseCases(repos *persistence2.Repositories, festivalStorage *persistence2.FestivalStorage) *usecase.UseCases {
	return usecase.NewUseCases(repos, festivalStorage)
}
