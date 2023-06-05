package api

import (
	"context"
	"encoding/json"
	"github.com/kruspe/music-rating/internal/model"
	"net/http"
)

type ratingResponse struct {
	ArtistName   string  `json:"artist_name"`
	Comment      string  `json:"comment"`
	FestivalName string  `json:"festival_name"`
	Rating       float64 `json:"rating"`
	Year         int     `json:"year,omitempty"`
}

type ratingRequest struct {
	ArtistName   string  `json:"artist_name"`
	Comment      string  `json:"comment"`
	FestivalName string  `json:"festival_name"`
	Rating       float64 `json:"rating"`
	Year         int     `json:"year"`
}

type updateRatingRequest struct {
	Comment      string  `json:"comment"`
	FestivalName string  `json:"festival_name"`
	Rating       float64 `json:"rating"`
	Year         int     `json:"year"`
}

type ratingRepo interface {
	GetAll(ctx context.Context, userId string) ([]model.Rating, error)
	Save(ctx context.Context, userId string, rating model.Rating) error
	Update(ctx context.Context, userId string, ratingUpdate model.Rating) error
}

type RatingEndpoint struct {
	ratingRepo ratingRepo
}

func NewRatingEndpoint(ratingRepo ratingRepo) *RatingEndpoint {
	return &RatingEndpoint{
		ratingRepo: ratingRepo,
	}
}

func (e *RatingEndpoint) create(w http.ResponseWriter, r *http.Request, userId string) {
	var rating ratingRequest
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		HandleError(w, err)
		return
	}
	if rating.ArtistName == "" {
		HandleError(w, model.MissingParameterError{ParameterName: "ArtistName"})
		return
	}
	if rating.Rating == 0 {
		HandleError(w, model.MissingParameterError{ParameterName: "Rating"})
		return
	}

	err = e.ratingRepo.Save(r.Context(), userId, model.Rating{
		ArtistName:   rating.ArtistName,
		Comment:      rating.Comment,
		FestivalName: rating.FestivalName,
		Rating:       rating.Rating,
		Year:         rating.Year,
	})
	if err != nil {
		HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (e *RatingEndpoint) getAll(w http.ResponseWriter, r *http.Request, userId string) {
	ratings, err := e.ratingRepo.GetAll(r.Context(), userId)
	if err != nil {
		HandleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(e.toRatingsResponse(ratings))
	if err != nil {
		HandleError(w, err)
		return
	}
}

func (e *RatingEndpoint) put(w http.ResponseWriter, r *http.Request, userId, artistName string) {
	var ratingUpdate updateRatingRequest
	err := json.NewDecoder(r.Body).Decode(&ratingUpdate)
	if err != nil {
		HandleError(w, err)
		return
	}

	err = e.ratingRepo.Update(r.Context(), userId, model.Rating{
		ArtistName:   artistName,
		Comment:      ratingUpdate.Comment,
		FestivalName: ratingUpdate.FestivalName,
		Rating:       ratingUpdate.Rating,
		Year:         ratingUpdate.Year,
	})
	if err != nil {
		HandleError(w, err)
		return
	}
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
