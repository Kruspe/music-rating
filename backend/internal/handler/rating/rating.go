package rating

import (
	"context"
	"encoding/json"
	. "github.com/kruspe/music-rating/internal/handler/errors"
	"github.com/kruspe/music-rating/internal/middleware"
	"github.com/kruspe/music-rating/internal/model"
	"github.com/kruspe/music-rating/internal/persistence"
	"github.com/kruspe/music-rating/internal/usecase"
	"net/http"
	"slices"
	"strings"
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
	GetAll(ctx context.Context, userId string) (*model.Ratings, error)
	Save(ctx context.Context, userId string, rating model.ArtistRating) error
	Update(ctx context.Context, userId string, ratingUpdate model.ArtistRating) error
}

var _ ratingRepo = &persistence.RatingRepo{}

type festivalUseCase interface {
	GetArtistsForFestival(ctx context.Context, festivalName string) ([]model.Artist, error)
}

var _ festivalUseCase = &usecase.FestivalUseCase{}

type Rating struct {
	ratingRepo      ratingRepo
	festivalUseCase festivalUseCase
}

func NewRatingEndpoint(ratingRepo ratingRepo, festivalUseCase festivalUseCase) *Rating {
	return &Rating{
		ratingRepo:      ratingRepo,
		festivalUseCase: festivalUseCase,
	}
}

func (e *Rating) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.UserIdContextKey).(string)
		var ratingRequest ratingRequest
		err := json.NewDecoder(r.Body).Decode(&ratingRequest)
		if err != nil {
			HandleError(w, err)
			return
		}
		rating, err := model.NewArtistRating(ratingRequest.ArtistName, ratingRequest.Rating, &ratingRequest.FestivalName, &ratingRequest.Year, &ratingRequest.Comment)
		if err != nil {
			HandleError(w, err)
			return
		}

		err = e.ratingRepo.Save(r.Context(), userId, *rating)
		if err != nil {
			HandleError(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func (e *Rating) GetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.UserIdContextKey).(string)
		ratings, err := e.ratingRepo.GetAll(r.Context(), userId)
		if err != nil {
			HandleError(w, err)
			return
		}
		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(e.toRatingsResponse(*ratings))
		if err != nil {
			HandleError(w, err)
			return
		}
	})
}

func (e *Rating) Put() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.UserIdContextKey).(string)
		artistName := r.PathValue("artistName")
		var ratingUpdate updateRatingRequest
		err := json.NewDecoder(r.Body).Decode(&ratingUpdate)
		if err != nil {
			HandleError(w, err)
			return
		}
		rating, err := model.NewArtistRating(artistName, ratingUpdate.Rating, &ratingUpdate.FestivalName, &ratingUpdate.Year, &ratingUpdate.Comment)
		if err != nil {
			HandleError(w, err)
			return
		}

		err = e.ratingRepo.Update(r.Context(), userId, *rating)
		if err != nil {
			HandleError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func (e *Rating) GetAllForFestival() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.UserIdContextKey).(string)
		festivalName := r.PathValue("festivalName")

		artists, err := e.festivalUseCase.GetArtistsForFestival(r.Context(), festivalName)
		if err != nil {
			HandleError(w, err)
			return
		}
		slices.SortFunc(artists, func(i, j model.Artist) int {
			return strings.Compare(i.Name, j.Name)
		})
		ratings, err := e.ratingRepo.GetAll(r.Context(), userId)
		if err != nil {
			HandleError(w, err)
			return
		}

		festivalRatings := model.Ratings{
			Keys:   make([]string, 0, len(artists)),
			Values: make(map[string]model.ArtistRating),
		}
		var notRatedArtist []string
		for _, artist := range artists {
			if rating, found := ratings.Values[artist.Name]; found {
				festivalRatings.Keys = append(festivalRatings.Keys, artist.Name)
				festivalRatings.Values[artist.Name] = rating
			} else {
				notRatedArtist = append(notRatedArtist, artist.Name)
			}
		}
		for _, artist := range notRatedArtist {
			rating, err := model.NewArtistRating(artist, 0, nil, nil, nil)
			if err != nil {
				HandleError(w, err)
				return
			}
			festivalRatings.Keys = append(festivalRatings.Keys, artist)
			festivalRatings.Values[artist] = *rating
		}

		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(e.toRatingsResponse(festivalRatings))
		if err != nil {
			HandleError(w, err)
			return
		}
	})
}

func (e *Rating) toRatingsResponse(ratings model.Ratings) []ratingResponse {
	result := make([]ratingResponse, 0, len(ratings.Keys))
	for _, key := range ratings.Keys {
		var (
			festivalName string
			year         int
			comment      string
		)
		if ratings.Values[key].FestivalName != nil {
			festivalName = *ratings.Values[key].FestivalName
		}
		if ratings.Values[key].Year != nil {
			year = *ratings.Values[key].Year
		}
		if ratings.Values[key].Comment != nil {
			comment = *ratings.Values[key].Comment
		}
		result = append(result, ratingResponse{
			ArtistName:   key,
			Rating:       ratings.Values[key].Rating.Float64(),
			FestivalName: festivalName,
			Year:         year,
			Comment:      comment,
		})
	}
	return result
}
