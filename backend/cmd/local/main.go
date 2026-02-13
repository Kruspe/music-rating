package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/kruspe/music-rating/internal/handler"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/helper"
	"github.com/kruspe/music-rating/internal/persistence"
	. "github.com/kruspe/music-rating/internal/persistence/helper"
	"github.com/kruspe/music-rating/internal/usecase"
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
		w.WriteHeader(http.StatusOK)
	}))
	handler.Register(mux, &handler.Config{
		RatingRepo:      repos.RatingRepo,
		FestivalUseCase: useCases.FestivalUseCase,
	})
	//nolint:gosec // local server for testing purposes only
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
