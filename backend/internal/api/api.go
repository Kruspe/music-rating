package api

import (
	"net/http"
)

func NewRouter(festivalEndpoint *FestivalEndpoint, ratingEndpoint *RatingEndpoint) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("POST /ratings", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(UserIdContextKey).(string)
		ratingEndpoint.create(w, r, userId)
	}))
	mux.Handle("GET /ratings", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(UserIdContextKey).(string)
		ratingEndpoint.getAll(w, r, userId)
	}))
	mux.Handle("PUT /ratings/{artistName}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(UserIdContextKey).(string)
		ratingEndpoint.put(w, r, userId, r.PathValue("artistName"))
	}))
	mux.Handle("GET /ratings/{festivalName}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(UserIdContextKey).(string)
		ratingEndpoint.getAllForFestival(w, r, userId, r.PathValue("festivalName"))
	}))
	mux.Handle("GET /festivals/{festivalName}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(UserIdContextKey).(string)
		festivalEndpoint.GetArtistsForFestival(w, r, userId, r.PathValue("festivalName"))
	}))

	return mux

}
