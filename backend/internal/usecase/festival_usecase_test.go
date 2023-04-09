package usecase_test

import (
	"context"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/model"
	"github.com/kruspe/music-rating/internal/model/model_test_helper"
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

	s.ratingUseCase = usecase.NewFestivalUseCase(s.ratingRepo, persistence.NewFestivalStorage(ph.ReturnArtists(
		[]model.Artist{
			model_test_helper.AnArtistWithName("Bloodbath"),
			model_test_helper.AnArtistWithName("Hypocrisy"),
			model_test_helper.AnArtistWithName("Benediction")})))
}

func (s *ratingUseCaseSuite) Test_GetUnratedArtistsForFestival_ReturnsUnratedArtists() {
	err := s.ratingRepo.Save(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.ratingRepo.Save(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)

	unratedArtists, err := s.ratingUseCase.GetUnratedArtistsForFestival(context.Background(), model_test_helper.TestUserId, "festival-name")
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Artist{model_test_helper.AnArtistWithName("Benediction")}, unratedArtists)
}
