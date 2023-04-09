package api

import (
	"context"
	"encoding/json"
	"github.com/kruspe/music-rating/internal/model"
	"net/http"
)

type unratedArtistResponse struct {
	ArtistName string `json:"artist_name"`
	ImageUrl   string `json:"image_url"`
}

type festivalUseCase interface {
	GetUnratedArtistsForFestival(ctx context.Context, userId, festivalName string) ([]model.Artist, error)
}

type FestivalEndpoint struct {
	festivalUseCase festivalUseCase
	errorHandler    *ErrorHandler
}

func NewFestivalEndpoint(festivalUseCase festivalUseCase, errorHandler *ErrorHandler) *FestivalEndpoint {
	return &FestivalEndpoint{
		festivalUseCase: festivalUseCase,
		errorHandler:    errorHandler,
	}
}

func (e *FestivalEndpoint) GetUnratedArtistsForFestival(w http.ResponseWriter, r *http.Request, userId, festivalName string) {
	unratedArtists, err := e.festivalUseCase.GetUnratedArtistsForFestival(r.Context(), userId, festivalName)
	if err != nil {
		e.errorHandler.Handle(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(e.toUnratedArtistsResponse(unratedArtists))
	if err != nil {
		e.errorHandler.Handle(w, err)
		return
	}
}

func (e *FestivalEndpoint) toUnratedArtistsResponse(artists []model.Artist) []unratedArtistResponse {
	result := make([]unratedArtistResponse, 0)
	for _, artist := range artists {
		result = append(result, unratedArtistResponse{
			ArtistName: artist.Name,
			ImageUrl:   artist.ImageUrl,
		})
	}
	return result
}
