package usecase_test

import (
	"backend/internal/adapter/model"
	"backend/internal/adapter/persistence"
	"backend/internal/adapter/persistence/test_helper"
	"backend/internal/usecase"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ratingUseCaseSuite struct {
	suite.Suite
	ratingRepo    *persistence.RatingRepo
	ratingUseCase *usecase.RatingUseCase
}

func Test_RatingUseCaseSuite(t *testing.T) {
	suite.Run(t, &ratingUseCaseSuite{})
}

func (s *ratingUseCaseSuite) BeforeTest(_ string, _ string) {
	ph := test_helper.NewPersistenceHelper()

	s.ratingRepo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)
	s.ratingUseCase = usecase.NewRatingUseCase(s.ratingRepo)
}

func (s *ratingUseCaseSuite) Test_GetRatings_ReturnsAllRatings() {
	err := s.ratingRepo.SaveRating(context.Background(), test_helper.TestUserId, test_helper.BloodbathRating)
	require.NoError(s.T(), err)

	ratings, err := s.ratingUseCase.GetRatings(context.Background(), test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{test_helper.BloodbathRating}, ratings)
}

func (s *ratingUseCaseSuite) Test_GetRatings_ReturnsError() {
	err := s.ratingRepo.SaveRating(context.Background(), test_helper.TestUserId, test_helper.BloodbathRating)
	require.NoError(s.T(), err)

	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()
	ratings, err := s.ratingUseCase.GetRatings(ctx, test_helper.TestUserId)
	require.ErrorContains(s.T(), err, "context canceled")
	require.Equal(s.T(), []model.Rating(nil), ratings)
}

func (s *ratingUseCaseSuite) Test_SaveRating_SavesRating() {
	err := s.ratingUseCase.SaveRating(context.Background(), test_helper.TestUserId, test_helper.BloodbathRating)
	require.NoError(s.T(), err)

	ratings, err := s.ratingRepo.GetRatings(context.Background(), test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{test_helper.BloodbathRating}, ratings)
}

func (s *ratingUseCaseSuite) Test_SaveRating_ReturnsError() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()
	err := s.ratingUseCase.SaveRating(ctx, test_helper.TestUserId, test_helper.BloodbathRating)
	require.ErrorContains(s.T(), err, "context canceled")
}
