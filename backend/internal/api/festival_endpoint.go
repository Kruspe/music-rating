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
	GetArtistsForFestival(ctx context.Context, festivalName string) ([]model.Artist, error)
}

type FestivalEndpoint struct {
	festivalUseCase festivalUseCase
}

func NewFestivalEndpoint(festivalUseCase festivalUseCase) *FestivalEndpoint {
	return &FestivalEndpoint{
		festivalUseCase: festivalUseCase,
	}
}

func (e *FestivalEndpoint) GetArtistsForFestival(w http.ResponseWriter, r *http.Request, userId, festivalName string) {
	var result []model.Artist
	if r.URL.Query().Get("filter") == "unrated" {
		unratedArtists, err := e.festivalUseCase.GetUnratedArtistsForFestival(r.Context(), userId, festivalName)
		if err != nil {
			HandleError(w, err)
			return
		}
		result = unratedArtists
	} else {
		artists, err := e.festivalUseCase.GetArtistsForFestival(r.Context(), festivalName)
		if err != nil {
			HandleError(w, err)
			return
		}
		result = artists
	}

	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(e.toUnratedArtistsResponse(result))
	if err != nil {
		HandleError(w, err)
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
