package persistence_test

import (
	"backend/internal/adapter/model"
	"backend/internal/adapter/persistence"
	"backend/internal/adapter/persistence/test_helper"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ratingRepoSuite struct {
	suite.Suite
	ph        test_helper.PersistenceHelper
	repo      *persistence.RatingRepo
	tableName string
}

func Test_RatingHandlerSuite(t *testing.T) {
	suite.Run(t, &ratingRepoSuite{})
}

func (s *ratingRepoSuite) BeforeTest(_ string, _ string) {
	ph := test_helper.NewPersistenceHelper()
	s.tableName = ph.TableName
	s.repo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)
}

func (s *ratingRepoSuite) Test_SaveRating_SavesRating() {
	err := s.repo.SaveRating(context.Background(), test_helper.TestUserId, test_helper.BloodbathRating)
	require.NoError(s.T(), err)

	ratings, err := s.repo.GetRatings(context.Background(), test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{test_helper.BloodbathRating}, ratings)
}

func (s *ratingRepoSuite) Test_SaveRating_ReturnsError() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	err := s.repo.SaveRating(ctx, "me", test_helper.BloodbathRating)
	require.ErrorContains(s.T(), err, "context canceled")

	ratings, err := s.repo.GetRatings(context.Background(), "me")
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating(nil), ratings)
}

func (s *ratingRepoSuite) Test_GetRatings_ReturnsAllRatingsForUser() {
	err := s.repo.SaveRating(context.Background(), test_helper.TestUserId, test_helper.BloodbathRating)
	require.NoError(s.T(), err)
	err = s.repo.SaveRating(context.Background(), test_helper.TestUserId, test_helper.HypocrisyRating)
	require.NoError(s.T(), err)

	ratings, err := s.repo.GetRatings(context.Background(), test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{test_helper.HypocrisyRating, test_helper.BloodbathRating}, ratings)
}

func (s *ratingRepoSuite) Test_GetRatings_ReturnsError() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	ratings, err := s.repo.GetRatings(ctx, test_helper.TestUserId)
	require.ErrorContains(s.T(), err, "context canceled")
	require.Equal(s.T(), []model.Rating(nil), ratings)
}