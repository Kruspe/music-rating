package api_test

import (
	"errors"
	"github.com/kruspe/music-rating/internal/api"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
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
	parameterError := model.MissingParameterError{ParameterName: "some_parameter"}
	api.HandleError(recorder, parameterError)

	require.Equal(s.T(), http.StatusBadRequest, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, parameterError.Error())
}

func (s *errorHandlerSuite) Test_Returns400_UpdateNonExistingRatingError() {
	recorder := httptest.NewRecorder()
	updateError := model.UpdateNonExistingRatingError{ArtistName: "artist_name"}
	api.HandleError(recorder, updateError)

	require.Equal(s.T(), http.StatusBadRequest, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, updateError.Error())
}

func (s *errorHandlerSuite) Test_Returns404_WhenFestivalNotSupportedError() {
	recorder := httptest.NewRecorder()
	notSupportedError := model.FestivalNotSupportedError{FestivalName: AFestivalName}
	api.HandleError(recorder, notSupportedError)

	require.Equal(s.T(), http.StatusNotFound, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, notSupportedError.Error())
}

func (s *errorHandlerSuite) Test_Returns500_GenericError() {
	recorder := httptest.NewRecorder()
	api.HandleError(recorder, errors.New("random error"))

	require.Equal(s.T(), http.StatusInternalServerError, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "random error")
}
