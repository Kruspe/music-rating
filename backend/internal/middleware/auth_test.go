package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kruspe/music-rating/internal/handler/test"
	"github.com/kruspe/music-rating/internal/middleware"
	. "github.com/kruspe/music-rating/internal/model/helper"
	"github.com/stretchr/testify/suite"
)

type authMiddlewareSuite struct {
	suite.Suite
}

func Test_AuthMiddlewareSuite(t *testing.T) {
	suite.Run(t, &authMiddlewareSuite{})
}

func (s *authMiddlewareSuite) Test_AuthorizeMiddleware_CallsNextEndpoint() {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/ratings", nil)
	s.Require().NoError(err)
	request.Header.Set("Authorization", test.TestToken)
	recorder := httptest.NewRecorder()
	nextEndpointCalled := false

	middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextEndpointCalled = true
		s.Equal(AnUserId, r.Context().Value(middleware.UserIdContextKey))
	})).ServeHTTP(recorder, request)

	s.True(nextEndpointCalled)
	s.Equal(http.StatusOK, recorder.Result().StatusCode)
}

func (s *authMiddlewareSuite) Test_AuthorizeMiddleware_Returns401_WhenTokenIsInvalid() {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/ratings", nil)
	s.Require().NoError(err)
	request.Header.Set("Authorization", "Bearer not_a_valid_token")
	recorder := httptest.NewRecorder()
	nextEndpointCalled := false

	middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextEndpointCalled = true
	})).ServeHTTP(recorder, request)

	s.False(nextEndpointCalled)
	s.Equal(http.StatusUnauthorized, recorder.Result().StatusCode)
}

func (s *authMiddlewareSuite) Test_AuthorizeMiddleware_Returns401_WhenNoSubIsSet() {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/ratings", nil)
	s.Require().NoError(err)
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjJ9.tbDepxpstvGdW8TC3G8zg4B6rUYAOvfzdceoH48wgRQ")
	recorder := httptest.NewRecorder()
	nextEndpointCalled := false

	middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextEndpointCalled = true
	})).ServeHTTP(recorder, request)

	s.False(nextEndpointCalled)
	s.Equal(http.StatusUnauthorized, recorder.Result().StatusCode)
}
