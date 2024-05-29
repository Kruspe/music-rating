package api

import (
	"encoding/json"
	. "github.com/kruspe/music-rating/internal/model"
	"log/slog"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *MissingParameterError:
		slog.Error("Missing parameter", slog.Any("error", err))
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		if err != nil {
			slog.Error("Error encoding response", slog.Any("error", err))
		}
	case *UpdateNonExistingRatingError:
		slog.Error("Update non existing rating", slog.Any("error", err))
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		if err != nil {
			slog.Error("Error encoding response", slog.Any("error", err))
		}
	case *FestivalNotSupportedError:
		slog.Error("Festival not supported", slog.Any("error", err))
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		if err != nil {
			slog.Error("Error encoding response", slog.Any("error", err))
		}
	default:
		slog.Error("Generic error", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(errorResponse{Error: "Something went wrong."})
		if err != nil {
			slog.Error("Error encoding response", slog.Any("error", err))
		}
	}
}
