package usecase

import (
	"backend/internal/adapter/model"
	"context"
)

type ratingRepo interface {
	GetRatings(ctx context.Context, userId string) ([]model.Rating, error)
	SaveRating(ctx context.Context, userId string, rating model.Rating) error
}

type RatingUseCase struct {
	ratingRepo ratingRepo
}

func NewRatingUseCase(ratingRepo ratingRepo) *RatingUseCase {
	return &RatingUseCase{
		ratingRepo: ratingRepo,
	}
}

func (u *RatingUseCase) GetRatings(ctx context.Context, userId string) ([]model.Rating, error) {
	return u.ratingRepo.GetRatings(ctx, userId)
}

func (u *RatingUseCase) SaveRating(ctx context.Context, userId string, rating model.Rating) error {
	return u.ratingRepo.SaveRating(ctx, userId, rating)
}
