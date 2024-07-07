package usecase

import (
	persistence2 "github.com/kruspe/music-rating/internal/persistence"
)

type UseCases struct {
	FestivalUseCase *FestivalUseCase
}

func NewUseCases(repos *persistence2.Repositories, festivalStorage *persistence2.FestivalStorage) *UseCases {
	return &UseCases{
		FestivalUseCase: NewFestivalUseCase(repos.RatingRepo, festivalStorage),
	}
}
