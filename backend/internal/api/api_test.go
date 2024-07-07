package api_test

import (
	"github.com/kruspe/music-rating/internal/api"
	. "github.com/kruspe/music-rating/internal/api/api_test_helper"
	persistence2 "github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type apiSuite struct {
	suite.Suite
	router *http.ServeMux
}

func Test_ApiSuite(t *testing.T) {
	suite.Run(t, &apiSuite{})
}

func (s *apiSuite) BeforeTest(_, _ string) {
	persistenceHelper := persistence_test_helper.NewPersistenceHelper()
	repos := persistence2.NewRepositories(persistenceHelper.Dynamo, persistenceHelper.TableName)
	useCases := usecase.NewUseCases(repos, persistence2.NewFestivalStorage(persistenceHelper.MockFestivals(nil)))
	festivalEndpoint := api.NewFestivalEndpoint(useCases.FestivalUseCase)
	ratingEndpoint := api.NewRatingEndpoint(repos.RatingRepo, useCases.FestivalUseCase)
	s.router = api.NewRouter(festivalEndpoint, ratingEndpoint)
}

func (s *apiSuite) Test_Returns404_WhenRequestPathDoesNotExist() {
	request := NewAuthenticatedRequest(http.MethodGet, "/not_existing", nil)
	recorder := httptest.NewRecorder()
	s.router.ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusNotFound, recorder.Result().StatusCode)
}

func (s *apiSuite) Test_Returns501_ratings_WhenMethodIsNotImplemented() {
	request := NewAuthenticatedRequest(http.MethodPut, "/ratings", nil)
	recorder := httptest.NewRecorder()
	s.router.ServeHTTP(recorder, request)

	require.Equal(s.T(), http.StatusMethodNotAllowed, recorder.Result().StatusCode)
}
