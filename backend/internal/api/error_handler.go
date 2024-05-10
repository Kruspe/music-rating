package api

import (
	"encoding/json"
	. "github.com/kruspe/music-rating/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case MissingParameterError:
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		if err != nil {
			log.Error(err)
		}
		break
	case UpdateNonExistingRatingError:
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		if err != nil {
			log.Error(err)
		}
		break
	case FestivalNotSupportedError:
		log.Error(err)
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		if err != nil {
			log.Error(err)
		}
		break
	default:
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(errorResponse{Error: "Something went wrong."})
		if err != nil {
			log.Error(err)
		}
	}
}
