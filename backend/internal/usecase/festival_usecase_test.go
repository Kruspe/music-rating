package usecase_test

import (
	"context"
	"testing"

	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/helper"
	"github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/persistence/helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/suite"
)

type ratingUseCaseSuite struct {
	suite.Suite
	ratingRepo    *persistence.RatingRepo
	ratingUseCase *usecase.FestivalUseCase
}

func Test_RatingUseCaseSuite(t *testing.T) {
	suite.Run(t, &ratingUseCaseSuite{})
}

func (s *ratingUseCaseSuite) BeforeTest(_ string, _ string) {
	ph := helper.NewPersistenceHelper()
	s.ratingRepo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)

	s.ratingUseCase = usecase.NewFestivalUseCase(s.ratingRepo, persistence.NewFestivalStorage(ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
			AnArtistWithName("Benediction")},
	})))

}

func (s *ratingUseCaseSuite) Test_GetUnratedArtistsForFestival_ReturnsUnratedArtists() {
	err := s.ratingRepo.Save(context.Background(), AnUserId, AnArtistRating("Bloodbath"))
	s.Require().NoError(err)
	err = s.ratingRepo.Save(context.Background(), AnUserId, AnArtistRating("Hypocrisy"))
	s.Require().NoError(err)

	unratedArtists, err := s.ratingUseCase.GetUnratedArtistsForFestival(context.Background(), AnUserId, AFestivalName)
	s.Require().NoError(err)
	s.Equal([]model.Artist{AnArtistWithName("Benediction")}, unratedArtists)
}

func (s *ratingUseCaseSuite) Test_GetUnratedArtistsForFestival_ReturnsFestivalNotSupportedError() {
	_, err := s.ratingUseCase.GetUnratedArtistsForFestival(context.Background(), AnUserId, AnotherFestivalName)
	s.Require().Error(err)
	s.Require().ErrorAs(err, new(*model.FestivalNotSupportedError))
}

func (s *ratingUseCaseSuite) Test_GetArtistsForFestival_ReturnsAllArtists() {
	artists, err := s.ratingUseCase.GetArtistsForFestival(context.Background(), AFestivalName)
	s.Require().NoError(err)
	s.Equal([]model.Artist{AnArtistWithName("Bloodbath"), AnArtistWithName("Hypocrisy"), AnArtistWithName("Benediction")}, artists)

}

func (s *ratingUseCaseSuite) Test_GetArtistsForFestival_ReturnsFestivalNotSupportedError() {
	_, err := s.ratingUseCase.GetArtistsForFestival(context.Background(), AnotherFestivalName)
	s.Require().Error(err)
	s.Require().ErrorAs(err, new(*model.FestivalNotSupportedError))
}
