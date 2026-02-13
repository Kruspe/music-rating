package persistence_test

import (
	"context"
	"testing"

	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/helper"
	"github.com/kruspe/music-rating/internal/persistence"
	. "github.com/kruspe/music-rating/internal/persistence/helper"
	"github.com/stretchr/testify/suite"
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
	s.Require().NoError(err)
	err = s.repo.Save(context.Background(), AnUserId, rating2)
	s.Require().NoError(err)

	ratings, err := s.repo.GetAll(context.Background(), AnUserId)
	s.Require().NoError(err)
	s.Equal(
		&model.Ratings{
			Keys:   []string{rating1.ArtistName, rating2.ArtistName},
			Values: map[string]model.ArtistRating{rating1.ArtistName: rating1, rating2.ArtistName: rating2},
		},
		ratings,
	)
}

func (s *ratingRepoSuite) Test_UpdateRating() {
	rating := AnArtistRating("Bloodbath")
	err := s.repo.Save(context.Background(), AnUserId, rating)
	s.Require().NoError(err)

	newFestivalName := "new-festival"
	newYear := 666
	updatedArtistRating, err := model.NewArtistRating(rating.ArtistName, float64(2), &newFestivalName, &newYear, nil)
	s.Require().NoError(err)

	err = s.repo.Update(context.Background(), AnUserId, *updatedArtistRating)
	s.Require().NoError(err)

	ratings, err := s.repo.GetAll(context.Background(), AnUserId)
	s.Require().NoError(err)
	s.Len(ratings.Keys, 1)
	s.Equal(rating.ArtistName, ratings.Keys[0])
	s.Len(ratings.Values, 1)
	s.Equal(rating.ArtistName, ratings.Values[rating.ArtistName].ArtistName)
	s.Equal(updatedArtistRating.FestivalName, ratings.Values[rating.ArtistName].FestivalName)
	s.InEpsilon(updatedArtistRating.Rating.Float64(), ratings.Values[rating.ArtistName].Rating.Float64(), 0.0001)
	s.Equal(updatedArtistRating.Year, ratings.Values[rating.ArtistName].Year)
	s.Nil(ratings.Values[rating.ArtistName].Comment)
}

func (s *ratingRepoSuite) Test_UpdateRating_FailsWhenRatingDoesNotExist() {
	updatedRating := AnArtistRating("non_existing_artist")
	err := s.repo.Update(context.Background(), AnUserId, updatedRating)
	s.ErrorAs(err, new(*model.UpdateNonExistingRatingError))
}
