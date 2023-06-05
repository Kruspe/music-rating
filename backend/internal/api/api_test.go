package api_test

import (
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/api"
	. "github.com/kruspe/music-rating/internal/api/api_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type apiSuite struct {
	suite.Suite
	api *api.Api
}

func Test_ApiSuite(t *testing.T) {
	suite.Run(t, &apiSuite{})
}

func (s *apiSuite) BeforeTest(_, _ string) {
	persistenceHelper := persistence_test_helper.NewPersistenceHelper()
	repos := persistence.NewRepositories(persistenceHelper.Dynamo, persistenceHelper.TableName)
	useCases := usecase.NewUseCases(repos, persistence.NewFestivalStorage(persistenceHelper.ReturnArtists(nil)))
	s.api = api.NewApi(useCases, repos)
}

func (s *apiSuite) Test_Returns404_WhenRequestPathDoesNotExist() {
	request := NewAuthenticatedRequest(http.MethodGet, "/not_existing", nil)
	recorder := httptest.NewRecorder()
	s.api.ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusNotFound, recorder.Result().StatusCode)
}

func (s *apiSuite) Test_Returns501_ratings_WhenMethodIsNotImplemented() {
	request := NewAuthenticatedRequest(http.MethodPut, "/ratings", nil)
	recorder := httptest.NewRecorder()
	s.api.ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusNotImplemented, recorder.Result().StatusCode)
}

func (s *apiSuite) Test_Returns401_WhenNoSubIsSet() {
	request, err := http.NewRequest(http.MethodGet, "/api/ratings", nil)
	require.NoError(s.T(), err)
	request.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjJ9.tbDepxpstvGdW8TC3G8zg4B6rUYAOvfzdceoH48wgRQ")
	recorder := httptest.NewRecorder()
	s.api.ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusUnauthorized, recorder.Result().StatusCode)
}
