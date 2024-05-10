package usecase_test

import (
	"context"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
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
	err := s.ratingRepo.Save(context.Background(), TestUserId, ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.ratingRepo.Save(context.Background(), TestUserId, ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)

	unratedArtists, err := s.ratingUseCase.GetUnratedArtistsForFestival(context.Background(), TestUserId, AFestivalName)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Artist{AnArtistWithName("Benediction")}, unratedArtists)
}
