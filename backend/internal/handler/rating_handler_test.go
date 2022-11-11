package handler_test

import (
	"backend/internal/adapter/model"
	"backend/internal/adapter/persistence"
	"backend/internal/adapter/persistence/test_helper"
	"backend/internal/handler"
	"backend/internal/usecase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ratingHandlerSuite struct {
	suite.Suite
	ratingRepo *persistence.RatingRepo
	handler    *handler.RatingHandler
}

func Test_RatingHandlerSuite(t *testing.T) {
	suite.Run(t, &ratingHandlerSuite{})
}

func (s *ratingHandlerSuite) BeforeTest(_ string, _ string) {
	ph := test_helper.NewPersistenceHelper()

	s.ratingRepo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)
	s.handler = handler.NewRatingHandler(usecase.NewRatingUseCase(s.ratingRepo))
}

func (s *ratingHandlerSuite) Test_Handle_CreateRating_Returns201() {
	rating, err := json.Marshal(test_helper.TestRatingDao)
	require.NoError(s.T(), err)

	response, err := s.handler.Handle(context.Background(), events.APIGatewayV2HTTPRequest{
		RawPath: "/ratings",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "POST",
			},
		},
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", test_helper.TestToken),
		},
		Body: string(rating),
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), http.StatusCreated, response.StatusCode)

	savedRating, err := s.ratingRepo.GetRatings(context.Background(), test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{test_helper.TestRating}, savedRating)
}

func (s *ratingHandlerSuite) Test_handle_CreateRating_Returns400WhenRatingIsMissingFields() {
	c := []struct {
		missingFieldName string
		rating           model.RatingDao
	}{
		{
			missingFieldName: "artist_name",
			rating: model.RatingDao{
				FestivalName: "Wacken",
				Rating:       aws.Int(5),
				Year:         aws.Int(666),
			},
		},
		{
			missingFieldName: "festival_name",
			rating: model.RatingDao{
				ArtistName: "Bloodbath",
				Rating:     aws.Int(5),
				Year:       aws.Int(666),
			},
		},
		{
			missingFieldName: "rating",
			rating: model.RatingDao{
				ArtistName:   "Bloodbath",
				FestivalName: "Wacken",
				Year:         aws.Int(666),
			},
		},
		{
			missingFieldName: "year",
			rating: model.RatingDao{
				ArtistName:   "Bloodbath",
				FestivalName: "Wacken",
				Rating:       aws.Int(5),
			},
		},
	}
	for _, testCase := range c {
		s.T().Run(fmt.Sprintf("missing %s", testCase.missingFieldName), func(t *testing.T) {
			rating, err := json.Marshal(testCase.rating)
			require.NoError(t, err)

			response, err := s.handler.Handle(context.Background(), events.APIGatewayV2HTTPRequest{
				RawPath: "/ratings",
				RequestContext: events.APIGatewayV2HTTPRequestContext{
					HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
						Method: "POST",
					},
				},
				Headers: map[string]string{
					"Authorization": fmt.Sprintf("Bearer %s", test_helper.TestToken),
				},
				Body: string(rating),
			})
			require.ErrorContains(t, err, fmt.Sprintf("missing %s", testCase.missingFieldName))
			require.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}

func (s *ratingHandlerSuite) Test_Handler_Returns401WhenSubjectIsMissingFromClaims() {
	rating, err := json.Marshal(model.RatingDao{
		ArtistName:   "Bloodbath",
		Comment:      "Amazing",
		FestivalName: "Wacken",
		Rating:       aws.Int(5),
		Year:         aws.Int(666),
	})
	require.NoError(s.T(), err)

	response, err := s.handler.Handle(context.Background(), events.APIGatewayV2HTTPRequest{
		RawPath: "/ratings",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "POST",
			},
		},
		Headers: map[string]string{
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjJ9.tbDepxpstvGdW8TC3G8zg4B6rUYAOvfzdceoH48wgRQ",
		},
		Body: string(rating),
	})
	require.ErrorContains(s.T(), err, "missing sub in token")
	require.Equal(s.T(), http.StatusUnauthorized, response.StatusCode)
}
