package api

import (
	. "github.com/kruspe/music-rating/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ErrorHandler struct {
	log *log.Logger
}

func NewErrorHandler(log *log.Logger) *ErrorHandler {
	return &ErrorHandler{log: log}
}

func (e *ErrorHandler) Handle(w http.ResponseWriter, err error) {
	switch err.(type) {
	case MissingParameterError:
		e.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	default:
		e.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
