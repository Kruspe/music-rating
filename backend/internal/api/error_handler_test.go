package api_test

import (
	"errors"
	"github.com/kruspe/music-rating/internal/api"
	"github.com/kruspe/music-rating/internal/model"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type errorHandlerSuite struct {
	suite.Suite
	logHook *test.Hook
}

func Test_ErrorHandlerSuite(t *testing.T) {
	suite.Run(t, &errorHandlerSuite{})
}

func (s *errorHandlerSuite) BeforeTest(_, _ string) {
	s.logHook = test.NewGlobal()
}

func (s *errorHandlerSuite) Test_Returns400_MissingParameterError() {
	recorder := httptest.NewRecorder()
	api.HandleError(recorder, model.MissingParameterError{ParameterName: "some_parameter"})

	require.Equal(s.T(), http.StatusBadRequest, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "missing parameter 'some_parameter'")
}

func (s *errorHandlerSuite) Test_Returns400_UpdateNonExistingRatingError() {
	recorder := httptest.NewRecorder()
	api.HandleError(recorder, model.UpdateNonExistingRatingError{ArtistName: "artist_name"})

	require.Equal(s.T(), http.StatusBadRequest, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "trying to update non existing rating for 'artist_name'")
}

func (s *errorHandlerSuite) Test_Returns500_GenericError() {
	recorder := httptest.NewRecorder()
	api.HandleError(recorder, errors.New("random error"))

	require.Equal(s.T(), http.StatusInternalServerError, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "random error")
}
