package usecase_test

import (
	"context"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
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
	ph := persistence_test_helper.NewPersistenceHelper()
	s.ratingRepo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)

	s.ratingUseCase = usecase.NewFestivalUseCase(s.ratingRepo, persistence.NewFestivalStorage(ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
			AnArtistWithName("Benediction")},
	})))

}

func (s *ratingUseCaseSuite) Test_GetUnratedArtistsForFestival_ReturnsUnratedArtists() {
	err := s.ratingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.ratingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)

	unratedArtists, err := s.ratingUseCase.GetUnratedArtistsForFestival(context.Background(), AnUserId, AFestivalName)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Artist{AnArtistWithName("Benediction")}, unratedArtists)
}

func (s *ratingUseCaseSuite) Test_GetUnratedArtistsForFestival_ReturnsFestivalNotSupportedError() {
	_, err := s.ratingUseCase.GetUnratedArtistsForFestival(context.Background(), AnUserId, AnotherFestivalName)
	require.Error(s.T(), err)
	require.IsType(s.T(), &model.FestivalNotSupportedError{}, err)
}

func (s *ratingUseCaseSuite) Test_GetArtistsForFestival_ReturnsAllArtists() {
	artists, err := s.ratingUseCase.GetArtistsForFestival(context.Background(), AFestivalName)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Artist{AnArtistWithName("Bloodbath"), AnArtistWithName("Hypocrisy"), AnArtistWithName("Benediction")}, artists)

}

func (s *ratingUseCaseSuite) Test_GetArtistsForFestival_ReturnsFestivalNotSupportedError() {
	_, err := s.ratingUseCase.GetArtistsForFestival(context.Background(), AnotherFestivalName)
	require.Error(s.T(), err)
	require.IsType(s.T(), &model.FestivalNotSupportedError{}, err)
}
