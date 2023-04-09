package main

import (
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	. "github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/api"
	. "github.com/kruspe/music-rating/internal/api/api_test_helper"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log := log.New()
	ph := NewPersistenceHelper()
	repos := persistence.NewRepositories(ph.Dynamo, ph.TableName)
	festivalStorage := persistence.NewFestivalStorage(ph.ReturnArtists([]model.Artist{
		AnArtistWithName("Bloodbath"),
		AnArtistWithName("Hypocrisy"),
		AnArtistWithName("Benediction"),
	}))
	useCases := usecase.NewUseCases(repos, festivalStorage)

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		r.Header.Set("Authorization", TestToken)
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
		} else {
			api.NewApi(useCases, repos, api.NewErrorHandler(log)).ServeHTTP(w, r)
		}
	}))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
