package api_test

import (
	"context"
	"encoding/json"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	. "github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/api"
	. "github.com/kruspe/music-rating/internal/api/api_test_helper"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	log "github.com/sirupsen/logrus"
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

func (s *festivalHandlerSuite) Test_GetUnratedArtistsForFestival_Returns200AndAllUnratedArtists() {
	err := s.repos.RatingRepo.Create(context.Background(), TestUserId, ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.repos.RatingRepo.Create(context.Background(), TestUserId, ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)

	unratedArtist := AnArtistWithName("Benediction")
	festivalStorage := persistence.NewFestivalStorage(s.ph.ReturnArtists([]model.Artist{
		AnArtistWithName("Bloodbath"),
		AnArtistWithName("Hypocrisy"),
		unratedArtist,
	}))
	api := api.NewApi(usecase.NewUseCases(s.repos, festivalStorage), s.repos, api.NewErrorHandler(log.New()))

	request := NewAuthenticatedRequest(http.MethodGet, "/festivals/wacken", nil)
	recorder := httptest.NewRecorder()
	api.ServeHTTP(recorder, request)

	require.NoError(s.T(), err)
	require.Equal(s.T(), http.StatusOK, recorder.Result().StatusCode)

	var r []unratedArtistResponse
	err = json.NewDecoder(recorder.Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 1)
	require.Equal(s.T(), unratedArtist.Name, r[0].ArtistName)
	require.Equal(s.T(), unratedArtist.ImageUrl, r[0].ImageUrl)
}

func (s *festivalHandlerSuite) Test_GetUnratedArtistsForFestival_Returns200AndEmptyList_WhenAllArtistsAreRated() {
	err := s.repos.RatingRepo.Create(context.Background(), TestUserId, ARatingForArtist("Bloodbath"))
	require.NoError(s.T(), err)
	err = s.repos.RatingRepo.Create(context.Background(), TestUserId, ARatingForArtist("Hypocrisy"))
	require.NoError(s.T(), err)
	err = s.repos.RatingRepo.Create(context.Background(), TestUserId, ARatingForArtist("Benediction"))
	require.NoError(s.T(), err)

	festivalStorage := persistence.NewFestivalStorage(s.ph.ReturnArtists([]model.Artist{
		AnArtistWithName("Bloodbath"),
		AnArtistWithName("Hypocrisy"),
		AnArtistWithName("Benediction"),
	}))
	api := api.NewApi(usecase.NewUseCases(s.repos, festivalStorage), s.repos, api.NewErrorHandler(log.New()))

	request := NewAuthenticatedRequest(http.MethodGet, "/festivals/wacken", nil)
	recorder := httptest.NewRecorder()
	api.ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusOK, recorder.Result().StatusCode)
	var r []unratedArtistResponse
	err = json.NewDecoder(recorder.Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 0)
}
