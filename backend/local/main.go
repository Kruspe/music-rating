package main

import (
	"backend/internal/adapter/persistence"
	"backend/internal/adapter/persistence/test_helper"
	"backend/internal/handler"
	"backend/internal/usecase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

type local struct {
	logger    *log.Logger
	dynamo    *dynamodb.Client
	tableName string
}

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&log.JSONFormatter{})
	ph := test_helper.NewPersistenceHelper()
	setupRatings(logger, ph)
	l := local{
		logger:    logger,
		dynamo:    ph.Dynamo,
		tableName: ph.TableName,
	}

	http.HandleFunc("/api/", l.handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal(err)
	}
}

func (l *local) handle(w http.ResponseWriter, r *http.Request) {
	repo := persistence.NewRatingRepo(l.dynamo, l.tableName)
	ratingUseCase := usecase.NewRatingUseCase(repo)
	h := handler.NewRatingHandler(ratingUseCase, l.logger)

	buffer := new(strings.Builder)
	_, err := io.Copy(buffer, r.Body)
	if err != nil {
		l.logger.Error(err)
	}
	event := events.APIGatewayV2HTTPRequest{
		Body: buffer.String(),
		Headers: map[string]string{
			"authorization": fmt.Sprintf("Bearer %s", test_helper.TestToken),
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
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
	return
}

func setupRatings(logger *log.Logger, ph *test_helper.PersistenceHelper) {
	bloodbathItem, err := attributevalue.MarshalMap(test_helper.BloodbathRatingRecord)
	if err != nil {
		logger.Fatal(err)
	}
	_, err = ph.Dynamo.PutItem(context.Background(), &dynamodb.PutItemInput{
		Item:      bloodbathItem,
		TableName: aws.String(ph.TableName),
	})
	if err != nil {
		logger.Fatal(err)
	}

	hypocrisyItem, err := attributevalue.MarshalMap(test_helper.HypocrisyRatingRecord)
	if err != nil {
		logger.Fatal(err)
	}
	_, err = ph.Dynamo.PutItem(context.Background(), &dynamodb.PutItemInput{
		Item:      hypocrisyItem,
		TableName: aws.String(ph.TableName),
	})
	if err != nil {
		logger.Fatal(err)
	}
}
