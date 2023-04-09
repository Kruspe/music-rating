package api

import (
	"context"
	"encoding/json"
	"github.com/kruspe/music-rating/internal/model"
	"net/http"
)

type ratingResponse struct {
	ArtistName   string `json:"artist_name"`
	Comment      string `json:"comment"`
	FestivalName string `json:"festival_name"`
	Rating       int    `json:"rating"`
	Year         int    `json:"year"`
}

type ratingRequest struct {
	ArtistName   string `json:"artist_name"`
	Comment      string `json:"comment"`
	FestivalName string `json:"festival_name"`
	Rating       int    `json:"rating"`
	Year         int    `json:"year"`
}

type ratingRepo interface {
	GetAll(ctx context.Context, userId string) ([]model.Rating, error)
	Create(ctx context.Context, userId string, rating model.Rating) error
}

type RatingEndpoint struct {
	ratingRepo   ratingRepo
	errorHandler *ErrorHandler
}

func NewRatingEndpoint(ratingRepo ratingRepo, errorHandler *ErrorHandler) *RatingEndpoint {
	return &RatingEndpoint{
		ratingRepo:   ratingRepo,
		errorHandler: errorHandler,
	}
}

func (e *RatingEndpoint) create(w http.ResponseWriter, r *http.Request, userId string) {
	var rating ratingRequest
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		e.errorHandler.Handle(w, err)
		return
	}
	if rating.ArtistName == "" {
		e.errorHandler.Handle(w, model.MissingParameterError{ParameterName: "ArtistName"})
		return
	}
	if rating.Rating == 0 {
		e.errorHandler.Handle(w, model.MissingParameterError{ParameterName: "Rating"})
		return
	}

	err = e.ratingRepo.Create(r.Context(), userId, model.Rating{
		ArtistName:   rating.ArtistName,
		Comment:      rating.Comment,
		FestivalName: rating.FestivalName,
		Rating:       rating.Rating,
		Year:         rating.Year,
	})
	if err != nil {
		e.errorHandler.Handle(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (e *RatingEndpoint) getAll(w http.ResponseWriter, r *http.Request, userId string) {
	ratings, err := e.ratingRepo.GetAll(r.Context(), userId)
	if err != nil {
		e.errorHandler.Handle(w, err)
		return
	}
	err = json.NewEncoder(w).Encode(e.toRatingsResponse(ratings))
	if err != nil {
		e.errorHandler.Handle(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (e *RatingEndpoint) toRatingsResponse(r []model.Rating) []ratingResponse {
	result := make([]ratingResponse, 0)
	for _, rating := range r {
		result = append(result, ratingResponse{
			ArtistName:   rating.ArtistName,
			Comment:      rating.Comment,
			FestivalName: rating.FestivalName,
			Rating:       rating.Rating,
			Year:         rating.Year,
		})
	}
	return result
}
