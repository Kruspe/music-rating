package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kruspe/music-rating/internal/handler"
	. "github.com/kruspe/music-rating/internal/handler/test"
	"github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/persistence/helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/suite"
)

type apiSuite struct {
	suite.Suite
	mux *http.ServeMux
}

func Test_ApiSuite(t *testing.T) {
	suite.Run(t, &apiSuite{})
}

func (s *apiSuite) BeforeTest(_, _ string) {
	persistenceHelper := helper.NewPersistenceHelper()
	repos := persistence.NewRepositories(persistenceHelper.Dynamo, persistenceHelper.TableName)
	useCases := usecase.NewUseCases(repos, persistence.NewFestivalStorage(persistenceHelper.MockFestivals(nil)))
	s.mux = http.NewServeMux()
	handler.Register(s.mux, &handler.Config{
		RatingRepo:      repos.RatingRepo,
		FestivalUseCase: useCases.FestivalUseCase,
	})
}

func (s *apiSuite) Test_Returns404_WhenRequestPathDoesNotExist() {
	request := NewAuthenticatedRequest(http.MethodGet, "/not_existing", nil)
	recorder := httptest.NewRecorder()
	s.mux.ServeHTTP(recorder, request)

	s.Equal(http.StatusNotFound, recorder.Result().StatusCode)
}

func (s *apiSuite) Test_Returns501_ratings_WhenMethodIsNotImplemented() {
	request := NewAuthenticatedRequest(http.MethodPut, "/ratings", nil)
	recorder := httptest.NewRecorder()
	s.mux.ServeHTTP(recorder, request)

	s.Equal(http.StatusMethodNotAllowed, recorder.Result().StatusCode)
}
