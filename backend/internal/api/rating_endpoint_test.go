package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	. "github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/api"
	. "github.com/kruspe/music-rating/internal/api/api_test_helper"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ratingResponse struct {
	ArtistName   string `json:"artist_name"`
	Comment      string `json:"comment"`
	FestivalName string `json:"festival_name"`
	Rating       int    `json:"rating"`
	Year         int    `json:"year"`
}

type ratingHandlerSuite struct {
	suite.Suite
	api *api.Api
}

func Test_RatingHandlerSuite(t *testing.T) {
	suite.Run(t, &ratingHandlerSuite{})
}

func (s *ratingHandlerSuite) BeforeTest(_ string, _ string) {
	ph := NewPersistenceHelper()
	repos := persistence.NewRepositories(ph.Dynamo, ph.TableName)
	s.api = api.NewApi(usecase.NewUseCases(repos, persistence.NewFestivalStorage(ph.ReturnArtists(nil))), repos, api.NewErrorHandler(log.New()))
}

func (s *ratingHandlerSuite) Test_PersistsRating() {
	rating := ARatingForArtist("Bloodbath")
	body, err := json.Marshal(map[string]interface{}{
		"artist_name":   rating.ArtistName,
		"comment":       rating.Comment,
		"festival_name": rating.FestivalName,
		"rating":        rating.Rating,
		"year":          rating.Year,
	})
	require.NoError(s.T(), err)
	put := NewAuthenticatedRequest(http.MethodPost, "/ratings", bytes.NewReader(body))
	require.NoError(s.T(), err)
	putRecorder := httptest.NewRecorder()

	s.api.ServeHTTP(putRecorder, put)
	require.Equal(s.T(), http.StatusCreated, putRecorder.Result().StatusCode)

	get := NewAuthenticatedRequest(http.MethodGet, "/ratings", nil)
	require.NoError(s.T(), err)
	getRecorder := httptest.NewRecorder()

	s.api.ServeHTTP(getRecorder, get)
	require.Equal(s.T(), http.StatusOK, getRecorder.Result().StatusCode)

	var r []ratingResponse
	err = json.NewDecoder(getRecorder.Result().Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 1)
	require.Equal(s.T(), rating.ArtistName, r[0].ArtistName)
	require.Equal(s.T(), rating.FestivalName, r[0].FestivalName)
	require.Equal(s.T(), rating.Rating, r[0].Rating)
	require.Equal(s.T(), rating.Year, r[0].Year)
	require.Equal(s.T(), rating.Comment, r[0].Comment)
}

func (s *ratingHandlerSuite) Test_UpdateRating() {
	rating := ARatingForArtist("Bloodbath")
	putBody, err := json.Marshal(map[string]interface{}{
		"artist_name":   rating.ArtistName,
		"comment":       rating.Comment,
		"festival_name": rating.FestivalName,
		"rating":        rating.Rating,
		"year":          rating.Year,
	})
	require.NoError(s.T(), err)
	put := NewAuthenticatedRequest(http.MethodPost, "/ratings", bytes.NewReader(putBody))
	putRecorder := httptest.NewRecorder()

	s.api.ServeHTTP(putRecorder, put)
	require.Equal(s.T(), http.StatusCreated, putRecorder.Result().StatusCode)

	patchBody, err := json.Marshal(map[string]interface{}{
		"comment":       AnotherComment,
		"festival_name": AnotherFestivalName,
		"rating":        AnotherRating,
		"year":          AnotherYear,
	})
	require.NoError(s.T(), err)
	patch := NewAuthenticatedRequest(http.MethodPatch, fmt.Sprintf("/ratings/%s", rating.ArtistName), bytes.NewReader(patchBody))
	patchRecorder := httptest.NewRecorder()

	s.api.ServeHTTP(patchRecorder, patch)
	require.Equal(s.T(), http.StatusOK, patchRecorder.Result().StatusCode)

	get := NewAuthenticatedRequest(http.MethodGet, "/ratings", nil)
	require.NoError(s.T(), err)
	getRecorder := httptest.NewRecorder()

	s.api.ServeHTTP(getRecorder, get)
	require.Equal(s.T(), http.StatusOK, getRecorder.Result().StatusCode)

	var r []ratingResponse
	err = json.NewDecoder(getRecorder.Result().Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 1)
	require.Equal(s.T(), rating.ArtistName, r[0].ArtistName)
	require.Equal(s.T(), AnotherFestivalName, r[0].FestivalName)
	require.Equal(s.T(), AnotherRating, r[0].Rating)
	require.Equal(s.T(), AnotherYear, r[0].Year)
	require.Equal(s.T(), AnotherComment, r[0].Comment)
}
