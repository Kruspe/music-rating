package middleware_test

import (
	"github.com/kruspe/music-rating/internal/handler/test"
	"github.com/kruspe/music-rating/internal/middleware"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type authMiddlewareSuite struct {
	suite.Suite
}

func Test_AuthMiddlewareSuite(t *testing.T) {
	suite.Run(t, &authMiddlewareSuite{})
}

func (s *authMiddlewareSuite) Test_AuthorizeMiddleware_CallsNextEndpoint() {
	request, err := http.NewRequest(http.MethodGet, "/api/ratings", nil)
	require.NoError(s.T(), err)
	request.Header.Set("authorization", test.TestToken)
	recorder := httptest.NewRecorder()
	nextEndpointCalled := false

	middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextEndpointCalled = true
		require.Equal(s.T(), AnUserId, r.Context().Value(middleware.UserIdContextKey))
	})).ServeHTTP(recorder, request)

	require.True(s.T(), nextEndpointCalled)
	require.Equal(s.T(), http.StatusOK, recorder.Result().StatusCode)
}

func (s *authMiddlewareSuite) Test_AuthorizeMiddleware_Returns401_WhenTokenIsInvalid() {
	request, err := http.NewRequest(http.MethodGet, "/api/ratings", nil)
	require.NoError(s.T(), err)
	request.Header.Set("authorization", "Bearer not_a_valid_token")
	recorder := httptest.NewRecorder()
	nextEndpointCalled := false

	middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextEndpointCalled = true
	})).ServeHTTP(recorder, request)

	require.False(s.T(), nextEndpointCalled)
	require.Equal(s.T(), http.StatusUnauthorized, recorder.Result().StatusCode)
}

func (s *authMiddlewareSuite) Test_AuthorizeMiddleware_Returns401_WhenNoSubIsSet() {
	request, err := http.NewRequest(http.MethodGet, "/api/ratings", nil)
	require.NoError(s.T(), err)
	request.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjJ9.tbDepxpstvGdW8TC3G8zg4B6rUYAOvfzdceoH48wgRQ")
	recorder := httptest.NewRecorder()
	nextEndpointCalled := false

	middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextEndpointCalled = true
	})).ServeHTTP(recorder, request)

	require.False(s.T(), nextEndpointCalled)
	require.Equal(s.T(), http.StatusUnauthorized, recorder.Result().StatusCode)
}
