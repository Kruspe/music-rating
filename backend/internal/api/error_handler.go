package api

import (
	. "github.com/kruspe/music-rating/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case MissingParameterError:
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	case UpdateNonExistingRatingError:
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	default:
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
