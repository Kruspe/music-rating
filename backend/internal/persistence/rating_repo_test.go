package persistence_test

import (
	"context"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/persistence"
	. "github.com/kruspe/music-rating/internal/persistence/persistence_test_helper"
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
	ph := NewPersistenceHelper()
	s.tableName = ph.TableName
	s.repo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)
}

func (s *ratingRepoSuite) Test_PersistsRatings() {
	rating1 := AnArtistRating("Bloodbath")
	rating2 := AnArtistRating("Hypocrisy")
	err := s.repo.Save(context.Background(), AnUserId, rating1)
	require.NoError(s.T(), err)
	err = s.repo.Save(context.Background(), AnUserId, rating2)
	require.NoError(s.T(), err)

	ratings, err := s.repo.GetAll(context.Background(), AnUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), model.Ratings{rating1.ArtistName: rating1, rating2.ArtistName: rating2}, ratings)
}

func (s *ratingRepoSuite) Test_UpdateRating() {
	rating := AnArtistRating("Bloodbath")
	err := s.repo.Save(context.Background(), AnUserId, rating)
	require.NoError(s.T(), err)

	newFestivalName := "new-festival"
	newYear := 666
	updatedArtistRating, err := model.NewArtistRating(rating.ArtistName, float64(2), &newFestivalName, &newYear, nil)
	require.NoError(s.T(), err)

	err = s.repo.Update(context.Background(), AnUserId, *updatedArtistRating)
	require.NoError(s.T(), err)

	ratings, err := s.repo.GetAll(context.Background(), AnUserId)
	require.NoError(s.T(), err)
	require.Len(s.T(), ratings, 1)
	require.Equal(s.T(), rating.ArtistName, ratings[rating.ArtistName].ArtistName)
	require.Equal(s.T(), updatedArtistRating.FestivalName, ratings[rating.ArtistName].FestivalName)
	require.Equal(s.T(), updatedArtistRating.Rating, ratings[rating.ArtistName].Rating)
	require.Equal(s.T(), updatedArtistRating.Year, ratings[rating.ArtistName].Year)
	require.Nil(s.T(), ratings[rating.ArtistName].Comment)
}

func (s *ratingRepoSuite) Test_UpdateRating_FailsWhenRatingDoesNotExist() {
	updatedRating := AnArtistRating("non_existing_artist")
	err := s.repo.Update(context.Background(), AnUserId, updatedRating)
	require.IsType(s.T(), &model.UpdateNonExistingRatingError{}, err)
}
