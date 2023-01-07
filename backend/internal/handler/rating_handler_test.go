package handler_test

import (
	"backend/internal/adapter/model"
	"backend/internal/adapter/model/model_test_helper"
	"backend/internal/adapter/persistence"
	"backend/internal/adapter/persistence/persistence_test_helper"
	"backend/internal/handler"
	"backend/internal/usecase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ratingHandlerSuite struct {
	suite.Suite
	ratingRepo      *persistence.RatingRepo
	festivalStorage usecase.FestivalStorage
	handler         *handler.RatingHandler
	logHook         *test.Hook
}

func Test_RatingHandlerSuite(t *testing.T) {
	suite.Run(t, &ratingHandlerSuite{})
}

func (s *ratingHandlerSuite) BeforeTest(_ string, _ string) {
	ph := persistence_test_helper.NewPersistenceHelper()

	s.ratingRepo = persistence.NewRatingRepo(ph.Dynamo, ph.TableName)
	s.festivalStorage = persistence.NewFestivalStorage(ph.S3Mock)
	logger, hook := test.NewNullLogger()
	s.logHook = hook
	s.handler = handler.NewRatingHandler(usecase.NewRatingUseCase(s.ratingRepo, s.festivalStorage), logger)
}

func (s *ratingHandlerSuite) Test_Handle_CreateRating_Returns201() {
	rating, err := json.Marshal(model_test_helper.BloodbathRatingDao)
	require.NoError(s.T(), err)

	response, err := s.handler.Handle(context.Background(), events.APIGatewayV2HTTPRequest{
		RawPath: "/ratings",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "POST",
			},
		},
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
		},
		Body: string(rating),
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), http.StatusCreated, response.StatusCode)

	savedRating, err := s.ratingRepo.GetRatings(context.Background(), model_test_helper.TestUserId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.Rating{model_test_helper.BloodbathRating}, savedRating)
}

func (s *ratingHandlerSuite) Test_Handle_CreateRating_Returns400WhenRatingIsMissingFields() {
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
					"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
				},
				Body: string(rating),
			})
			require.ErrorContains(t, err, fmt.Sprintf("missing %s", testCase.missingFieldName))
			require.Equal(t, http.StatusBadRequest, response.StatusCode)
			require.Equal(t, fmt.Sprintf("Request did not include %s", testCase.missingFieldName), s.logHook.LastEntry().Message)
		})
	}
}

func (s *ratingHandlerSuite) Test_Handle_CreateRating_Returns500WhenContextIsCanceled() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	rating, err := json.Marshal(model_test_helper.BloodbathRatingDao)
	require.NoError(s.T(), err)

	response, err := s.handler.Handle(ctx, events.APIGatewayV2HTTPRequest{
		RawPath: "/ratings",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "POST",
			},
		},
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
		},
		Body: string(rating),
	})
	require.ErrorContains(s.T(), err, "context canceled")
	require.Equal(s.T(), http.StatusInternalServerError, response.StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "context canceled")
}

func (s *ratingHandlerSuite) Test_Handle_GetRatings_Returns200AndAllRatings() {
	err := s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.BloodbathRating)
	require.NoError(s.T(), err)

	response, err := s.handler.Handle(context.Background(), events.APIGatewayV2HTTPRequest{
		RawPath: "/ratings",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
			},
		},
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
		},
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), http.StatusOK, response.StatusCode)
	var result []model.RatingDao
	err = json.Unmarshal([]byte(response.Body), &result)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.RatingDao{model_test_helper.BloodbathRatingDao}, result)
}

func (s *ratingHandlerSuite) Test_Handle_GetRatings_Returns500WhenContextIsCanceled() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	response, err := s.handler.Handle(ctx, events.APIGatewayV2HTTPRequest{
		RawPath: "/ratings",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
			},
		},
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
		},
	})
	require.ErrorContains(s.T(), err, "context canceled")
	require.Equal(s.T(), http.StatusInternalServerError, response.StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "context canceled")
}

func (s *ratingHandlerSuite) Test_Handle_GetUnratedArtistsForFestival_Returns200AndAllUnratedArtists() {
	err := s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.BloodbathRating)
	require.NoError(s.T(), err)
	err = s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.HypocrisyRating)
	require.NoError(s.T(), err)

	response, err := s.handler.Handle(context.Background(), events.APIGatewayV2HTTPRequest{
		RawPath: "/festivals/wacken",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
			},
		},
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
		},
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), http.StatusOK, response.StatusCode)
	var result []model.ArtistDao
	err = json.Unmarshal([]byte(response.Body), &result)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.ArtistDao{model_test_helper.UnratedArtistDao}, result)
}

func (s *ratingHandlerSuite) Test_Handle_GetUnratedArtistsForFestival_Returns200AndEmptyList_WhenAllArtistsAreRated() {
	err := s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.BloodbathRating)
	require.NoError(s.T(), err)
	err = s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.HypocrisyRating)
	require.NoError(s.T(), err)
	err = s.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model.Rating{
		ArtistName:   model_test_helper.UnratedArtist.ArtistName,
		FestivalName: "Concert",
		Rating:       2022,
		Year:         5,
	})
	require.NoError(s.T(), err)

	response, err := s.handler.Handle(context.Background(), events.APIGatewayV2HTTPRequest{
		RawPath: "/festivals/wacken",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
			},
		},
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
		},
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), http.StatusOK, response.StatusCode)
	var result []model.ArtistDao
	err = json.Unmarshal([]byte(response.Body), &result)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []model.ArtistDao{}, result)
}

func (s *ratingHandlerSuite) Test_Handle_GetUnratedArtistsForFestival_Returns500WhenContextIsCanceled() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	response, err := s.handler.Handle(ctx, events.APIGatewayV2HTTPRequest{
		RawPath: "/festivals/wacken",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
			},
		},
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
		},
	})
	require.ErrorContains(s.T(), err, "context canceled")
	require.Equal(s.T(), http.StatusInternalServerError, response.StatusCode)
	require.Contains(s.T(), s.logHook.LastEntry().Message, "context canceled")
}

func (s *ratingHandlerSuite) Test_Handler_Returns401WhenSubjectIsMissingFromClaims() {
	rating, err := json.Marshal(model_test_helper.BloodbathRatingDao)
	require.NoError(s.T(), err)

	response, err := s.handler.Handle(context.Background(), events.APIGatewayV2HTTPRequest{
		RawPath: "/ratings",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "POST",
			},
		},
		Headers: map[string]string{
			"authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjJ9.tbDepxpstvGdW8TC3G8zg4B6rUYAOvfzdceoH48wgRQ",
		},
		Body: string(rating),
	})
	require.ErrorContains(s.T(), err, "missing sub in token")
	require.Equal(s.T(), http.StatusUnauthorized, response.StatusCode)
}
