package api_test

import (
	"encoding/json"
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

type errorResponse struct {
	Error string `json:"error"`
}

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
	resp := recorder.Result()

	require.Contains(s.T(), s.logHook.LastEntry().Message, parameterError.Error())
	require.Equal(s.T(), http.StatusBadRequest, resp.StatusCode)

	var respBody errorResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	require.NoError(s.T(), err)
	require.Equal(s.T(), errorResponse{Error: parameterError.Error()}, respBody)
}

func (s *errorHandlerSuite) Test_Returns400_UpdateNonExistingRatingError() {
	recorder := httptest.NewRecorder()
	updateError := model.UpdateNonExistingRatingError{ArtistName: "artist_name"}
	api.HandleError(recorder, updateError)
	resp := recorder.Result()

	require.Contains(s.T(), s.logHook.LastEntry().Message, updateError.Error())
	require.Equal(s.T(), http.StatusBadRequest, resp.StatusCode)

	var respBody errorResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	require.NoError(s.T(), err)
	require.Equal(s.T(), errorResponse{Error: updateError.Error()}, respBody)
}

func (s *errorHandlerSuite) Test_Returns404_WhenFestivalNotSupportedError() {
	recorder := httptest.NewRecorder()
	notSupportedError := &model.FestivalNotSupportedError{FestivalName: AFestivalName}
	api.HandleError(recorder, notSupportedError)
	resp := recorder.Result()

	require.Contains(s.T(), s.logHook.LastEntry().Message, notSupportedError.Error())
	require.Equal(s.T(), http.StatusNotFound, resp.StatusCode)

	var respBody errorResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	require.NoError(s.T(), err)
	require.Equal(s.T(), errorResponse{Error: notSupportedError.Error()}, respBody)
}

func (s *errorHandlerSuite) Test_Returns500_GenericError() {
	recorder := httptest.NewRecorder()
	api.HandleError(recorder, errors.New("random error"))
	resp := recorder.Result()

	require.Contains(s.T(), s.logHook.LastEntry().Message, "random error")
	require.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)

	var respBody errorResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	require.NoError(s.T(), err)
	require.Equal(s.T(), errorResponse{Error: "Something went wrong."}, respBody)
}
