package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	. "github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/api"
	. "github.com/kruspe/music-rating/internal/api/api_test_helper"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type unratedArtistResponse struct {
	ArtistName string `json:"artist_name"`
	ImageUrl   string `json:"image_url"`
}

type festivalHandlerSuite struct {
	suite.Suite
	repos *persistence.Repositories
	ph    *PersistenceHelper
}

func TestFestivalHandlerSuite(t *testing.T) {
	suite.Run(t, &festivalHandlerSuite{})
}

func (s *festivalHandlerSuite) BeforeTest(_, _ string) {
	s.ph = NewPersistenceHelper()
	s.repos = persistence.NewRepositories(s.ph.Dynamo, s.ph.TableName)
}

func (s *festivalHandlerSuite) Test_GetArtistsForFestival_Returns200AndAllArtists() {
	err := s.repos.RatingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.repos.RatingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)

	bloodbath := AnArtistWithName("Bloodbath")
	hypocrisy := AnArtistWithName("Hypocrisy")
	festivalStorage := persistence.NewFestivalStorage(s.ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			bloodbath,
			hypocrisy,
		},
	}))
	useCases := usecase.NewUseCases(s.repos, festivalStorage)
	festivalEndpoint := api.NewFestivalEndpoint(useCases.FestivalUseCase)
	ratingEndpoint := api.NewRatingEndpoint(s.repos.RatingRepo, useCases.FestivalUseCase)

	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/festivals/%s", AFestivalName), nil)
	recorder := httptest.NewRecorder()
	api.AuthMiddleware(api.NewRouter(festivalEndpoint, ratingEndpoint)).ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusOK, recorder.Result().StatusCode)

	var r []unratedArtistResponse
	err = json.NewDecoder(recorder.Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 2)
	require.Equal(s.T(), bloodbath.Name, r[0].ArtistName)
	require.Equal(s.T(), bloodbath.ImageUrl, r[0].ImageUrl)
	require.Equal(s.T(), hypocrisy.Name, r[1].ArtistName)
	require.Equal(s.T(), hypocrisy.ImageUrl, r[1].ImageUrl)
}

func (s *festivalHandlerSuite) Test_GetArtistsForFestival_Returns200AndAllUnratedArtists_WhenFiltering() {
	err := s.repos.RatingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.repos.RatingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)

	unratedArtist := AnArtistWithName("Benediction")
	festivalStorage := persistence.NewFestivalStorage(s.ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
			unratedArtist,
		},
	}))
	useCases := usecase.NewUseCases(s.repos, festivalStorage)
	festivalEndpoint := api.NewFestivalEndpoint(useCases.FestivalUseCase)
	ratingEndpoint := api.NewRatingEndpoint(s.repos.RatingRepo, useCases.FestivalUseCase)

	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/festivals/%s?filter=unrated", AFestivalName), nil)
	recorder := httptest.NewRecorder()
	api.AuthMiddleware(api.NewRouter(festivalEndpoint, ratingEndpoint)).ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusOK, recorder.Result().StatusCode)

	var r []unratedArtistResponse
	err = json.NewDecoder(recorder.Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 1)
	require.Equal(s.T(), unratedArtist.Name, r[0].ArtistName)
	require.Equal(s.T(), unratedArtist.ImageUrl, r[0].ImageUrl)
}

func (s *festivalHandlerSuite) Test_GetArtistsForFestival_Returns200AndEmptyList_WhenFilteringAndAllArtistsAreRated() {
	err := s.repos.RatingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.repos.RatingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)
	err = s.repos.RatingRepo.Save(context.Background(), AnUserId, ARatingForArtist("Benediction"))
	require.NoError(s.T(), err)

	festivalStorage := persistence.NewFestivalStorage(s.ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
			AnArtistWithName("Benediction"),
		},
	}))
	useCases := usecase.NewUseCases(s.repos, festivalStorage)
	festivalEndpoint := api.NewFestivalEndpoint(useCases.FestivalUseCase)
	ratingEndpoint := api.NewRatingEndpoint(s.repos.RatingRepo, useCases.FestivalUseCase)

	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/festivals/%s?filter=unrated", AFestivalName), nil)
	recorder := httptest.NewRecorder()
	api.AuthMiddleware(api.NewRouter(festivalEndpoint, ratingEndpoint)).ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusOK, recorder.Result().StatusCode)
	var r []unratedArtistResponse
	err = json.NewDecoder(recorder.Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 0)
}

func (s *festivalHandlerSuite) Test_GetArtistsForFestival_Returns404_WhenFestivalIsNotSupported() {
	festivalStorage := persistence.NewFestivalStorage(s.ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
		},
	}))
	useCases := usecase.NewUseCases(s.repos, festivalStorage)
	festivalEndpoint := api.NewFestivalEndpoint(useCases.FestivalUseCase)
	ratingEndpoint := api.NewRatingEndpoint(s.repos.RatingRepo, useCases.FestivalUseCase)

	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/festivals/%s", AnotherFestivalName), nil)
	recorder := httptest.NewRecorder()
	api.AuthMiddleware(api.NewRouter(festivalEndpoint, ratingEndpoint)).ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusNotFound, recorder.Result().StatusCode)
	var r errorResponse
	err := json.NewDecoder(recorder.Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Equal(s.T(), model.FestivalNotSupportedError{FestivalName: AnotherFestivalName}.Error(), r.Error)
}
