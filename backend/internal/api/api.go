package api

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/usecase"
	"net/http"
	"strings"
)

type Api struct {
	festivalEndpoint *FestivalEndpoint
	ratingEndpoint   *RatingEndpoint
}

func NewApi(useCases *usecase.UseCases, repos *persistence.Repositories) *Api {
	return &Api{
		ratingEndpoint:   NewRatingEndpoint(repos.RatingRepo),
		festivalEndpoint: NewFestivalEndpoint(useCases.FestivalUseCase),
	}
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId, err := a.getUserId(strings.SplitAfter(r.Header.Get("authorization"), "Bearer ")[1])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	var festival string
	var artistName string
	switch {
	case match(r.URL.Path, "/ratings"):
		switch r.Method {
		case http.MethodPost:
			a.ratingEndpoint.create(w, r, userId)
		case http.MethodGet:
			a.ratingEndpoint.getAll(w, r, userId)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	case match(r.URL.Path, "/ratings/+", &artistName) && r.Method == http.MethodPut:
		a.ratingEndpoint.put(w, r, userId, artistName)
	// TODO handle not implemented festivals
	case match(r.URL.Path, "/festivals/+", &festival):
		a.festivalEndpoint.GetUnratedArtistsForFestival(w, r, userId, festival)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (a *Api) getUserId(token string) (string, error) {
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
