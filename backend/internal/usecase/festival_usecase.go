package usecase

import (
	"context"
	"github.com/kruspe/music-rating/internal/model"
)

type FestivalStorage interface {
	GetArtists(ctx context.Context, festivalName string) ([]model.Artist, error)
}

type ratingRepo interface {
	GetAll(ctx context.Context, userId string) ([]model.Rating, error)
}

type FestivalUseCase struct {
	ratingRepo      ratingRepo
	festivalStorage FestivalStorage
}

func NewFestivalUseCase(ratingRepo ratingRepo, festivalStorage FestivalStorage) *FestivalUseCase {
	return &FestivalUseCase{
		ratingRepo:      ratingRepo,
		festivalStorage: festivalStorage,
	}
}

func (u *FestivalUseCase) GetUnratedArtistsForFestival(ctx context.Context, userId, festivalName string) ([]model.Artist, error) {
	artists, err := u.festivalStorage.GetArtists(ctx, festivalName)
	if err != nil {
		return nil, err
	}

	ratings, err := u.ratingRepo.GetAll(ctx, userId)
	if err != nil {
		return nil, err
	}
	ratedArtists := make(map[string]struct{}, 0)
	for _, rated := range ratings {
		ratedArtists[rated.ArtistName] = struct{}{}
	}

	unratedArtists := make([]model.Artist, 0)
	for _, artist := range artists {
		if _, found := ratedArtists[artist.Name]; !found {
			unratedArtists = append(unratedArtists, artist)
		}
	}
	return unratedArtists, nil
}
