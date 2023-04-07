package persistence_test

import (
	"context"
	"github.com/kruspe/music-rating/internal/adapter/model"
	"github.com/kruspe/music-rating/internal/adapter/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ratingRepoSuite struct {
	suite.Suite
	repo      *persistence.RatingRepo
	tableName string
}

func Test_RatingHandlerSuite(t *testing.T) {
	suite.Run(t, &ratingRepoSuite{})
}

func (s *ratingRepoSuite) BeforeTest(_ string, _ string) {
	ph := persistence_test_helper.NewPersistenceHelper()
	s.tableName = ph.TableName
	s.repo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)
}

func (s *ratingRepoSuite) Test_SaveRating_SavesRating() {
	rating := model_test_helper.ARatingForArtist("Bloodbath")
	err := s.repo.SaveRating(context.Background(), model_test_helper.TestUserId, rating)
	require.NoError(s.T(), err)

	ratings, err := s.repo.GetRatings(context.Background(), model_test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{rating}, ratings)
}

func (s *ratingRepoSuite) Test_SaveRating_ReturnsError() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	err := s.repo.SaveRating(ctx, "me", model_test_helper.ARatingForArtist("Bloodbath"))
	require.ErrorContains(s.T(), err, "context canceled")

	ratings, err := s.repo.GetRatings(context.Background(), "me")
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating(nil), ratings)
}

func (s *ratingRepoSuite) Test_GetRatings_ReturnsAllRatingsForUser() {
	rating1 := model_test_helper.ARatingForArtist("Bloodbath")
	rating2 := model_test_helper.ARatingForArtist("Hypocrisy")
	err := s.repo.SaveRating(context.Background(), model_test_helper.TestUserId, rating1)
	require.NoError(s.T(), err)
	err = s.repo.SaveRating(context.Background(), model_test_helper.TestUserId, rating2)
	require.NoError(s.T(), err)

	ratings, err := s.repo.GetRatings(context.Background(), model_test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{rating1, rating2}, ratings)
}

func (s *ratingRepoSuite) Test_GetRatings_ReturnsError() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	ratings, err := s.repo.GetRatings(ctx, model_test_helper.TestUserId)
	require.ErrorContains(s.T(), err, "context canceled")
	require.Equal(s.T(), []model.Rating(nil), ratings)
}
