package usecase

import (
	"context"
	"github.com/kruspe/music-rating/internal/adapter/model"
)

type ratingRepo interface {
	GetRatings(ctx context.Context, userId string) ([]model.Rating, error)
	SaveRating(ctx context.Context, userId string, rating model.Rating) error
}

type FestivalStorage interface {
	GetFestival(ctx context.Context, festivalName string) (model.Festival, error)
}

type RatingUseCase struct {
	ratingRepo      ratingRepo
	festivalStorage FestivalStorage
}

func NewRatingUseCase(ratingRepo ratingRepo, festivalStorage FestivalStorage) *RatingUseCase {
	return &RatingUseCase{
		ratingRepo:      ratingRepo,
		festivalStorage: festivalStorage,
	}
}

func (u *RatingUseCase) GetRatings(ctx context.Context, userId string) ([]model.Rating, error) {
	return u.ratingRepo.GetRatings(ctx, userId)
}

func (u *RatingUseCase) SaveRating(ctx context.Context, userId string, rating model.Rating) error {
	return u.ratingRepo.SaveRating(ctx, userId, rating)
}

func (u *RatingUseCase) GetUnratedArtistsForFestival(ctx context.Context, userId, festivalName string) ([]model.Artist, error) {
	ratings, err := u.ratingRepo.GetRatings(ctx, userId)
	if err != nil {
		return []model.Artist{}, err
	}
	ratedArtists := make(map[string]struct{}, len(ratings))
	for _, rated := range ratings {
		ratedArtists[rated.ArtistName] = struct{}{}
	}

	festival, err := u.festivalStorage.GetFestival(ctx, festivalName)
	if err != nil {
		return []model.Artist{}, err
	}

	var unratedArtists []model.Artist
	for _, artist := range festival.Artists {
		if _, found := ratedArtists[artist.ArtistName]; !found {
			unratedArtists = append(unratedArtists, artist)
		}
	}
	return unratedArtists, nil
}
