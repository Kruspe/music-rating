package usecase

import "github.com/kruspe/music-rating/internal/adapter/persistence"

type UseCases struct {
	FestivalUseCase *FestivalUseCase
}

func NewUseCases(repos *persistence.Repositories, festivalStorage *persistence.FestivalStorage) *UseCases {
	return &UseCases{
		FestivalUseCase: NewFestivalUseCase(repos.RatingRepo, festivalStorage),
	}
}
