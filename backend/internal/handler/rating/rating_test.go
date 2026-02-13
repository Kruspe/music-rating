package rating_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kruspe/music-rating/internal/handler"
	"github.com/kruspe/music-rating/internal/handler/errors"
	. "github.com/kruspe/music-rating/internal/handler/test"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/helper"
	"github.com/kruspe/music-rating/internal/persistence"
	. "github.com/kruspe/music-rating/internal/persistence/helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/suite"
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
	body, err := json.Marshal(map[string]any{
		"artist_name":   rating.ArtistName,
		"comment":       rating.Comment,
		"festival_name": rating.FestivalName,
		"rating":        rating.Rating,
		"year":          rating.Year,
	})
	s.Require().NoError(err)
	put := NewAuthenticatedRequest(http.MethodPost, "/ratings", bytes.NewReader(body))
	s.Require().NoError(err)
	putRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(putRecorder, put)
	s.Equal(http.StatusCreated, putRecorder.Result().StatusCode)

	get := NewAuthenticatedRequest(http.MethodGet, "/ratings", nil)
	getRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(getRecorder, get)
	s.Equal(http.StatusOK, getRecorder.Result().StatusCode)

	var r []ratingResponse
	err = json.NewDecoder(getRecorder.Result().Body).Decode(&r)
	s.Require().NoError(err)
	s.Len(r, 1)
	s.Equal(rating.ArtistName, r[0].ArtistName)
	s.InEpsilon(rating.Rating.Float64(), r[0].Rating, 0.0001)
	s.Equal(*rating.FestivalName, r[0].FestivalName)
	s.Equal(*rating.Year, r[0].Year)
	s.Equal(*rating.Comment, r[0].Comment)
}

func (s *ratingHandlerSuite) Test_UpdateRating() {
	rating := AnArtistRating("Bloodbath")
	putBody, err := json.Marshal(map[string]any{
		"artist_name":   rating.ArtistName,
		"comment":       rating.Comment,
		"festival_name": rating.FestivalName,
		"rating":        rating.Rating,
		"year":          rating.Year,
	})
	s.Require().NoError(err)
	create := NewAuthenticatedRequest(http.MethodPost, "/ratings", bytes.NewReader(putBody))
	putRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(putRecorder, create)
	s.Equal(http.StatusCreated, putRecorder.Result().StatusCode)

	updateBody, err := json.Marshal(map[string]any{
		"comment":       AnotherComment,
		"festival_name": AnotherFestivalName,
		"rating":        AnotherRating,
		"year":          AnotherYear,
	})
	s.Require().NoError(err)
	update := NewAuthenticatedRequest(http.MethodPut, fmt.Sprintf("/ratings/%s", rating.ArtistName), bytes.NewReader(updateBody))
	updateRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(updateRecorder, update)
	s.Equal(http.StatusOK, updateRecorder.Result().StatusCode)

	get := NewAuthenticatedRequest(http.MethodGet, "/ratings", nil)
	s.Require().NoError(err)
	getRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(getRecorder, get)
	s.Equal(http.StatusOK, getRecorder.Result().StatusCode)

	var r []ratingResponse
	err = json.NewDecoder(getRecorder.Result().Body).Decode(&r)
	s.Require().NoError(err)
	s.Len(r, 1)
	s.Equal(rating.ArtistName, r[0].ArtistName)
	s.Equal(AnotherFestivalName, r[0].FestivalName)
	s.InEpsilon(AnotherRating.Float64(), r[0].Rating, 0.0001)
	s.Equal(AnotherYear, r[0].Year)
	s.Equal(AnotherComment, r[0].Comment)
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
	s.Equal(http.StatusOK, resp.StatusCode)

	var r []ratingResponse
	err := json.NewDecoder(resp.Body).Decode(&r)
	s.Require().NoError(err)
	s.Len(r, 3)
	s.Equal(desertedFearRating.ArtistName, r[0].ArtistName)
	s.Equal(hypocrisyRating.ArtistName, r[1].ArtistName)
	s.Equal("Bloodbath", r[2].ArtistName)
}

func (s *ratingHandlerSuite) Test_GetAllForFestival_Returns404_WhenFestivalIsNotSupported() {
	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/ratings/%s", AnotherFestivalName), nil)
	recorder := httptest.NewRecorder()
	s.mux.ServeHTTP(recorder, request)

	s.Equal(http.StatusNotFound, recorder.Result().StatusCode)
	var r errors.ErrorResponse
	err := json.NewDecoder(recorder.Body).Decode(&r)
	s.Require().NoError(err)
	s.Equal(model.FestivalNotSupportedError{FestivalName: AnotherFestivalName}.Error(), r.Error)
}

func (s *ratingHandlerSuite) saveRating(rating model.ArtistRating) {
	body, err := json.Marshal(map[string]any{
		"artist_name":   rating.ArtistName,
		"comment":       rating.Comment,
		"festival_name": rating.FestivalName,
		"rating":        rating.Rating,
		"year":          rating.Year,
	})
	s.Require().NoError(err)
	put := NewAuthenticatedRequest(http.MethodPost, "/ratings", bytes.NewReader(body))
	s.Require().NoError(err)
	putRecorder := httptest.NewRecorder()

	s.mux.ServeHTTP(putRecorder, put)
	s.Equal(http.StatusCreated, putRecorder.Result().StatusCode)
}
