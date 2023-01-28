package main

import (
	"backend/internal/adapter/model/model_test_helper"
	"backend/internal/adapter/persistence"
	"backend/internal/adapter/persistence/persistence_test_helper"
	"backend/internal/handler"
	"backend/internal/usecase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

type local struct {
	logger          *log.Logger
	ratingRepo      *persistence.RatingRepo
	festivalStorage *persistence.Storage
}

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&log.JSONFormatter{})
	ph := persistence_test_helper.NewPersistenceHelper()
	repo := persistence.NewRatingRepo(ph.Dynamo, ph.TableName)
	festivalStorage := persistence.NewFestivalStorage(ph.S3Mock)

	l := local{
		logger:          logger,
		ratingRepo:      repo,
		festivalStorage: festivalStorage,
	}
	l.setupRatings()

	http.HandleFunc("/api/", l.handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal(err)
	}
}

func (l *local) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}

	ratingUseCase := usecase.NewRatingUseCase(l.ratingRepo, l.festivalStorage)
	h := handler.NewRatingHandler(ratingUseCase, l.logger)

	buffer := new(strings.Builder)
	_, err := io.Copy(buffer, r.Body)
	if err != nil {
		l.logger.Error(err)
	}
	event := events.APIGatewayV2HTTPRequest{
		Body: buffer.String(),
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", model_test_helper.TestToken),
		},
		RawPath: strings.SplitAfter(r.URL.EscapedPath(), "/api")[1],
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: r.Method,
			},
		},
	}
	response, err := h.Handle(context.Background(), event)
	if err != nil {
		log.Error(err)
	}

	w.WriteHeader(response.StatusCode)
	if response.Body != "" {
		var r interface{}
		err := json.Unmarshal([]byte(response.Body), &r)
		if err != nil {
			l.logger.Error(err)
		}
		body, err := json.Marshal(r)
		if err != nil {
			l.logger.Error(err)
		}
		_, err = w.Write(body)
		if err != nil {
			l.logger.Error(err)
		}
	}
}

func (l *local) setupRatings() {
	err := l.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Bloodbath"))
	if err != nil {
		l.logger.Fatal(err)
	}
	err = l.ratingRepo.SaveRating(context.Background(), model_test_helper.TestUserId, model_test_helper.ARatingForArtist("Hypocrisy"))
	if err != nil {
		l.logger.Fatal(err)
	}
}
