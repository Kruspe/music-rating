package persistence

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dbKey struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`
}

func ratingDbKey(userId, artistName string) dbKey {
	return dbKey{
		PK: fmt.Sprintf("USER#%s", userId),
		SK: fmt.Sprintf("ARTIST#%s", artistName),
	}
}

const (
	RatingType = "RATING"
)

const (
	SkPkGsi = "SK-PK-GSI"
)

type Repositories struct {
	RatingRepo *RatingRepo
}

func NewRepositories(dynamo *dynamodb.Client, tableName string) *Repositories {
	return &Repositories{
		RatingRepo: NewRatingRepo(dynamo, tableName),
	}
}
