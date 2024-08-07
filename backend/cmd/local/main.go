package main

import (
	"github.com/kruspe/music-rating/internal/api"
	. "github.com/kruspe/music-rating/internal/api/api_test_helper"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/persistence"
	. "github.com/kruspe/music-rating/internal/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	ph := NewPersistenceHelper()
	repos := persistence.NewRepositories(ph.Dynamo, ph.TableName)
	festivalStorage := persistence.NewFestivalStorage(ph.MockFestivals(map[string][]model.Artist{
		"wacken": {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
			AnArtistWithName("Benediction"),
		},
		"dong": {
			AnArtistWithName("Deserted Fear"),
		},
	}))
	useCases := usecase.NewUseCases(repos, festivalStorage)

	mux := http.NewServeMux()
	mux.Handle("OPTIONS /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	router := api.NewRouter(api.NewFestivalEndpoint(useCases.FestivalUseCase), api.NewRatingEndpoint(repos.RatingRepo, useCases.FestivalUseCase))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		r.Header.Set("Authorization", TestToken)
		api.AuthMiddleware(router).ServeHTTP(w, r)
	}))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
