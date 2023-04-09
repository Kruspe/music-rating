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
	errorHandler *api.ErrorHandler
	logHook      *test.Hook
}

func Test_ErrorHandlerSuite(t *testing.T) {
	suite.Run(t, &errorHandlerSuite{})
}

func (s *errorHandlerSuite) BeforeTest(_, _ string) {
	logger, hook := test.NewNullLogger()
	s.logHook = hook
	s.errorHandler = api.NewErrorHandler(logger)
}

func (s *errorHandlerSuite) Test_Returns400_MissingParameterError() {
	recorder := httptest.NewRecorder()
	s.errorHandler.Handle(recorder, model.MissingParameterError{ParameterName: "some_parameter"})

	require.Equal(s.T(), http.StatusBadRequest, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "missing parameter 'some_parameter'")
}

func (s *errorHandlerSuite) Test_Returns500_GenericError() {
	recorder := httptest.NewRecorder()
	s.errorHandler.Handle(recorder, errors.New("random error"))

	require.Equal(s.T(), http.StatusInternalServerError, recorder.Result().StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "random error")
}
