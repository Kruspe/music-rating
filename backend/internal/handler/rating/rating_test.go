package rating_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kruspe/music-rating/internal/handler"
	"github.com/kruspe/music-rating/internal/handler/errors"
	. "github.com/kruspe/music-rating/internal/handler/test"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/persistence"
	. "github.com/kruspe/music-rating/internal/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ratingResponse struct {
	ArtistName   string  `json:"artist_name"`
	Comment      string  `json:"comment"`
	FestivalName string  `json:"festival_name"`
	Rating       float64 `json:"rating"`
	Year         int     `json:"year"`
}

type ratingHandlerSuite struct {
	suite.Suite
	mux *http.ServeMux
}

func Test_RatingHandlerSuite(t *testing.T) {
	suite.Run(t, &ratingHandlerSuite{})
}

func (s *ratingHandlerSuite) BeforeTest(_ string, _ string) {
	ph := NewPersistenceHelper()
	repos := persistence.NewRepositories(ph.Dynamo, ph.TableName)
	useCases := usecase.NewUseCases(repos, persistence.NewFestivalStorage(ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
			AnArtistWithName("Deserted Fear"),
		},
	})))
	s.mux = http.NewServeMux()
	handler.Register(s.mux, &handler.Config{
		RatingRepo:      repos.RatingRepo,
		FestivalUseCase: useCases.FestivalUseCase,
	})
}

func (s *ratingHandlerSuite) Test_PersistsRating() {
	rating := AnArtistRating("Bloodbath")
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

	s.mux.ServeHTTP(putRecorder, put)
	require.Equal(s.T(), http.StatusCreated, putRecorder.Result().StatusCode)

	get := NewAuthenticatedRequest(http.MethodGet, "/ratings", nil)
	getRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(getRecorder, get)
	require.Equal(s.T(), http.StatusOK, getRecorder.Result().StatusCode)

	var r []ratingResponse
	err = json.NewDecoder(getRecorder.Result().Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 1)
	require.Equal(s.T(), rating.ArtistName, r[0].ArtistName)
	require.Equal(s.T(), rating.Rating.Float64(), r[0].Rating)
	require.Equal(s.T(), *rating.FestivalName, r[0].FestivalName)
	require.Equal(s.T(), *rating.Year, r[0].Year)
	require.Equal(s.T(), *rating.Comment, r[0].Comment)
}

func (s *ratingHandlerSuite) Test_UpdateRating() {
	rating := AnArtistRating("Bloodbath")
	putBody, err := json.Marshal(map[string]interface{}{
		"artist_name":   rating.ArtistName,
		"comment":       rating.Comment,
		"festival_name": rating.FestivalName,
		"rating":        rating.Rating,
		"year":          rating.Year,
	})
	require.NoError(s.T(), err)
	create := NewAuthenticatedRequest(http.MethodPost, "/ratings", bytes.NewReader(putBody))
	putRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(putRecorder, create)
	require.Equal(s.T(), http.StatusCreated, putRecorder.Result().StatusCode)

	updateBody, err := json.Marshal(map[string]interface{}{
		"comment":       AnotherComment,
		"festival_name": AnotherFestivalName,
		"rating":        AnotherRating,
		"year":          AnotherYear,
	})
	require.NoError(s.T(), err)
	update := NewAuthenticatedRequest(http.MethodPut, fmt.Sprintf("/ratings/%s", rating.ArtistName), bytes.NewReader(updateBody))
	updateRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(updateRecorder, update)
	require.Equal(s.T(), http.StatusOK, updateRecorder.Result().StatusCode)

	get := NewAuthenticatedRequest(http.MethodGet, "/ratings", nil)
	require.NoError(s.T(), err)
	getRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(getRecorder, get)
	require.Equal(s.T(), http.StatusOK, getRecorder.Result().StatusCode)

	var r []ratingResponse
	err = json.NewDecoder(getRecorder.Result().Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 1)
	require.Equal(s.T(), rating.ArtistName, r[0].ArtistName)
	require.Equal(s.T(), AnotherFestivalName, r[0].FestivalName)
	require.Equal(s.T(), AnotherRating.Float64(), r[0].Rating)
	require.Equal(s.T(), AnotherYear, r[0].Year)
	require.Equal(s.T(), AnotherComment, r[0].Comment)
}

func (s *ratingHandlerSuite) Test_GetAllForFestival() {
	hypocrisyRating := AnArtistRatingWithRating("Hypocrisy", ARating.Float64())
	s.saveRating(hypocrisyRating)
	desertedFearRating := AnArtistRatingWithRating("Deserted Fear", AnotherRating.Float64())
	s.saveRating(desertedFearRating)

	get := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/ratings/%s", AFestivalName), nil)
	getRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(getRecorder, get)
	resp := getRecorder.Result()
	require.Equal(s.T(), http.StatusOK, resp.StatusCode)

	var r []ratingResponse
	err := json.NewDecoder(resp.Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Len(s.T(), r, 3)
	require.Equal(s.T(), desertedFearRating.ArtistName, r[0].ArtistName)
	require.Equal(s.T(), hypocrisyRating.ArtistName, r[1].ArtistName)
	require.Equal(s.T(), "Bloodbath", r[2].ArtistName)
}

func (s *ratingHandlerSuite) Test_GetAllForFestival_Returns404_WhenFestivalIsNotSupported() {
	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/ratings/%s", AnotherFestivalName), nil)
	recorder := httptest.NewRecorder()
	s.mux.ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusNotFound, recorder.Result().StatusCode)
	var r errors.ErrorResponse
	err := json.NewDecoder(recorder.Body).Decode(&r)
	require.NoError(s.T(), err)
	require.Equal(s.T(), model.FestivalNotSupportedError{FestivalName: AnotherFestivalName}.Error(), r.Error)
}

func (s *ratingHandlerSuite) saveRating(rating model.ArtistRating) {
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

	s.mux.ServeHTTP(putRecorder, put)
	require.Equal(s.T(), http.StatusCreated, putRecorder.Result().StatusCode)
}
