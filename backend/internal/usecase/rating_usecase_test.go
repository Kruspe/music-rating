package usecase_test

import (
	"backend/internal/adapter/model"
	"backend/internal/adapter/model/model_test_helper"
	"backend/internal/adapter/persistence"
	"backend/internal/adapter/persistence/persistence_test_helper"
	"backend/internal/usecase"
	"context"
	"errors"
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

type mockFestivalStorage struct {
	getFestival func(ctx context.Context, festivalName string) (model.Festival, error)
}

func (m mockFestivalStorage) GetFestival(ctx context.Context, festivalName string) (model.Festival, error) {
	return m.getFestival(ctx, festivalName)
}

func (s *ratingUseCaseSuite) BeforeTest(_ string, _ string) {
	ph := persistence_test_helper.NewPersistenceHelper()
	s.ratingRepo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)

	festivalStorage := func() usecase.FestivalStorage {
		return mockFestivalStorage{
			getFestival: func(ctx context.Context, festivalName string) (model.Festival, error) {
				return model_test_helper.AFestivalWithArtists([]model.Artist{
					model_test_helper.AnArtistWithName("Bloodbath"),
					model_test_helper.AnArtistWithName("Hypocrisy"),
					model_test_helper.AnArtistWithName("Benediction"),
				}), nil
			},
		}
	}

	s.ratingUseCase = usecase.NewRatingUseCase(s.ratingRepo, festivalStorage())
}

func (s *ratingUseCaseSuite) Test_GetRatings_ReturnsAllRatings() {
	rating := model_test_helper.ARatingForArtist("Bloodbath")
	err := s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, rating)
	require.NoError(s.T(), err)

	ratings, err := s.ratingUseCase.GetRatings(context.Background(), model_test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{rating}, ratings)
}

func (s *ratingUseCaseSuite) Test_GetRatings_ReturnsError() {
	err := s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)

	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()
	ratings, err := s.ratingUseCase.GetRatings(ctx, model_test_helper.TestUserId)
	require.ErrorContains(s.T(), err, "context canceled")
	require.Equal(s.T(), []model.Rating(nil), ratings)
}

func (s *ratingUseCaseSuite) Test_SaveRating_SavesRating() {
	rating := model_test_helper.ARatingForArtist("Bloodbath")
	err := s.ratingUseCase.SaveRating(context.Background(), model_test_helper.TestUserId, rating)
	require.NoError(s.T(), err)

	ratings, err := s.ratingRepo.GetRatings(context.Background(), model_test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{rating}, ratings)
}

func (s *ratingUseCaseSuite) Test_SaveRating_ReturnsError() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()
	err := s.ratingUseCase.SaveRating(ctx, model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Bloodbath"))
	require.ErrorContains(s.T(), err, "context canceled")
}

func (s *ratingUseCaseSuite) Test_GetUnratedArtistsForFestival_ReturnsUnratedArtists() {
	err := s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)

	unratedArtists, err := s.ratingUseCase.GetUnratedArtistsForFestival(context.Background(), model_test_helper.TestUserId, "festival-name")
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Artist{model_test_helper.AnArtistWithName("Benediction")}, unratedArtists)
}

func (s *ratingUseCaseSuite) Test_GetUnratedArtistsForFestival_ReturnsWhenGettingRatingsFails() {
	err := s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	_, err = s.ratingUseCase.GetUnratedArtistsForFestival(ctx, model_test_helper.TestUserId, "festival-name")
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "context canceled")
}

func (s *ratingUseCaseSuite) Test_GetUnratedArtistsForFestival_ReturnsWhenGettingFestivalFails() {
	err := s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)

	festivalStorage := func() usecase.FestivalStorage {
		return mockFestivalStorage{
			getFestival: func(ctx context.Context, festivalName string) (model.Festival, error) {
				return model.Festival{}, errors.New("fetching festival failed")
			},
		}
	}
	ratingUseCase := usecase.NewRatingUseCase(s.ratingRepo, festivalStorage())
	_, err = ratingUseCase.GetUnratedArtistsForFestival(context.Background(), model_test_helper.TestUserId, "festival-name")
	require.ErrorContains(s.T(), err, "fetching festival failed")
}
