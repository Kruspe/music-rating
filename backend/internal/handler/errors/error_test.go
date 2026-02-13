package errors_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/kruspe/music-rating/internal/handler/errors"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/helper"
	"github.com/stretchr/testify/suite"
)

type errorResponse struct {
	Error string `json:"error"`
}

type errorHandlerSuite struct {
	suite.Suite
}

func Test_ErrorHandlerSuite(t *testing.T) {
	suite.Run(t, &errorHandlerSuite{})
}

func (s *errorHandlerSuite) Test_Returns400_MissingParameterError() {
	recorder := httptest.NewRecorder()
	parameterError := &model.MissingParameterError{ParameterName: "some_parameter"}
	HandleError(recorder, parameterError)
	resp := recorder.Result()

	s.Equal(http.StatusBadRequest, resp.StatusCode)

	var respBody errorResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	s.Require().NoError(err)
	s.Equal(errorResponse{Error: parameterError.Error()}, respBody)
}

func (s *errorHandlerSuite) Test_Returns400_UpdateNonExistingRatingError() {
	recorder := httptest.NewRecorder()
	updateError := &model.UpdateNonExistingRatingError{ArtistName: "artist_name"}
	HandleError(recorder, updateError)
	resp := recorder.Result()

	s.Equal(http.StatusBadRequest, resp.StatusCode)

	var respBody errorResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	s.Require().NoError(err)
	s.Equal(errorResponse{Error: updateError.Error()}, respBody)
}

func (s *errorHandlerSuite) Test_Returns404_WhenFestivalNotSupportedError() {
	recorder := httptest.NewRecorder()
	notSupportedError := &model.FestivalNotSupportedError{FestivalName: AFestivalName}
	HandleError(recorder, notSupportedError)
	resp := recorder.Result()

	s.Equal(http.StatusNotFound, resp.StatusCode)

	var respBody errorResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	s.Require().NoError(err)
	s.Equal(errorResponse{Error: notSupportedError.Error()}, respBody)
}

func (s *errorHandlerSuite) Test_Returns500_GenericError() {
	recorder := httptest.NewRecorder()
	HandleError(recorder, errors.New("random error"))
	resp := recorder.Result()

	s.Equal(http.StatusInternalServerError, resp.StatusCode)

	var respBody errorResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	s.Require().NoError(err)
	s.Equal(errorResponse{Error: "Something went wrong."}, respBody)
}
