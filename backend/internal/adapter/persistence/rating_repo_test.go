package persistence_test

import (
	"context"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	. "github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
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
	rating1 := ARatingForArtist("Bloodbath")
	rating2 := ARatingForArtist("Hypocrisy")
	err := s.repo.Save(context.Background(), TestUserId, rating1)
	require.NoError(s.T(), err)
	err = s.repo.Save(context.Background(), TestUserId, rating2)
	require.NoError(s.T(), err)

	ratings, err := s.repo.GetAll(context.Background(), TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{rating1, rating2}, ratings)
}

func (s *ratingRepoSuite) Test_Patch_UpdatesRating() {
	rating := ARatingForArtist("Bloodbath")
	err := s.repo.Save(context.Background(), TestUserId, rating)
	require.NoError(s.T(), err)

	ratingUpdate := ARatingUpdateForArtist("Bloodbath")
	err = s.repo.Update(context.Background(), TestUserId, "Bloodbath", ratingUpdate)
	require.NoError(s.T(), err)

	ratings, err := s.repo.GetAll(context.Background(), TestUserId)
	require.NoError(s.T(), err)
	require.Len(s.T(), ratings, 1)
	require.Equal(s.T(), rating.ArtistName, ratings[0].ArtistName)
	require.Equal(s.T(), ratingUpdate.FestivalName, ratings[0].FestivalName)
	require.Equal(s.T(), ratingUpdate.Rating, ratings[0].Rating)
	require.Equal(s.T(), ratingUpdate.Year, ratings[0].Year)
	require.Equal(s.T(), ratingUpdate.Comment, ratings[0].Comment)
}

func (s *ratingRepoSuite) Test_UpdateRating_FailsWhenRatingDoesNotExist() {
	err := s.repo.Update(context.Background(), TestUserId, "non_existing_artist", model.RatingUpdate{
		Comment:      AComment,
		FestivalName: AFestivalName,
		Rating:       ARating,
		Year:         AYear,
	})
	require.ErrorIs(s.T(), err, model.UpdateNonExistingRatingError{ArtistName: "non_existing_artist"})
}
