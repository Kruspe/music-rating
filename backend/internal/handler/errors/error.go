package errors

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	. "github.com/kruspe/music-rating/internal/model"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.As(err, new(*MissingParameterError)):
		slog.Error("Missing parameter", slog.Any("error", err))
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		if err != nil {
			slog.Error("Error encoding response", slog.Any("error", err))
		}
	case errors.As(err, new(*UpdateNonExistingRatingError)):
		slog.Error("Update non existing rating", slog.Any("error", err))
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		if err != nil {
			slog.Error("Error encoding response", slog.Any("error", err))
		}
	case errors.As(err, new(*FestivalNotSupportedError)):
		slog.Error("Festival not supported", slog.Any("error", err))
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		if err != nil {
			slog.Error("Error encoding response", slog.Any("error", err))
		}
	default:
		slog.Error("Generic error", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Something went wrong."})
		if err != nil {
			slog.Error("Error encoding response", slog.Any("error", err))
		}
	}
}
