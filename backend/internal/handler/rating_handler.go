package handler

import (
	"backend/internal/adapter/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type ratingUseCase interface {
	GetRatings(ctx context.Context, userId string) ([]model.Rating, error)
	SaveRating(ctx context.Context, userId string, rating model.Rating) error
}

type RatingHandler struct {
	ratingUseCase ratingUseCase
}

func NewRatingHandler(ratingUseCase ratingUseCase) *RatingHandler {
	return &RatingHandler{
		ratingUseCase: ratingUseCase,
	}
}

func (h *RatingHandler) Handle(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	userId, err := getUserId(strings.SplitAfter(event.Headers["Authorization"], "Bearer ")[1])
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusUnauthorized}, err
	}

	switch {
	case event.RawPath == "/ratings" && event.RequestContext.HTTP.Method == http.MethodPost:
		return h.createRating(ctx, userId, event.Body)
	case event.RawPath == "/ratings" && event.RequestContext.HTTP.Method == http.MethodGet:
		return h.getRatings(ctx, userId)
	}
	return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusNotFound}, nil
}

func (h *RatingHandler) createRating(ctx context.Context, userId string, body string) (events.APIGatewayV2HTTPResponse, error) {
	var rating model.RatingDao
	err := json.Unmarshal([]byte(body), &rating)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, err
	}
	if rating.ArtistName == "" {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest}, errors.New("missing artist_name from rating")
	}
	if rating.FestivalName == "" {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest}, errors.New("missing festival_name from rating")
	}
	if rating.Rating == nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest}, errors.New("missing rating from rating")
	}
	if rating.Year == nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest}, errors.New("missing year from rating")
	}

	err = h.ratingUseCase.SaveRating(ctx, userId, model.Rating{
		ArtistName:   rating.ArtistName,
		Comment:      rating.Comment,
		FestivalName: rating.FestivalName,
		Rating:       *rating.Rating,
		Year:         *rating.Year,
	})
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusCreated}, nil
}

func (h *RatingHandler) getRatings(ctx context.Context, userId string) (events.APIGatewayV2HTTPResponse, error) {
	ratings, err := h.ratingUseCase.GetRatings(ctx, userId)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, err
	}
	var ratingDaos []model.RatingDao
	for _, r := range ratings {
		ratingDaos = append(ratingDaos, toRatingDao(r))
	}
	result, err := json.Marshal(ratingDaos)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, err
	}
	return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusOK, Body: string(result)}, nil
}

func toRatingDao(rating model.Rating) model.RatingDao {
	return model.RatingDao{
		ArtistName:   rating.ArtistName,
		Comment:      rating.Comment,
		FestivalName: rating.FestivalName,
		Rating:       &rating.Rating,
		Year:         &rating.Year,
	}
}

func getUserId(token string) (string, error) {
	var claims jwt.RegisteredClaims
	_, _, err := jwt.NewParser().ParseUnverified(token, &claims)
	if err != nil {
		return "", err
	}
	if claims.Subject == "" {
		return "", errors.New("missing sub in token")
	}
	return claims.Subject, nil
}
