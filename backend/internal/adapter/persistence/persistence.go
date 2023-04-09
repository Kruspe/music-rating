package persistence

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DbKey struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`
}

const (
	RatingType = "RATING"
)

const (
	ServiceOfferingType      = "SERVICE_OFFERING"
	ServiceOfferingOwnerType = "SERVICE_OFFERING_OWNER"
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
