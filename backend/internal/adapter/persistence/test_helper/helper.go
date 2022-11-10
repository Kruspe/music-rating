package test_helper

import (
	"backend/internal/adapter/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"os"
)

const (
	TestToken  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJNZXRhbExvdmVyNjY2IiwiaWF0IjoxNTE2MjM5MDIyfQ.JZ3R_3it-1K9ttH5NA80fpIsBhnW6DNsIzwB2zEFRmo7hgE-HhW3jJbArXNS0fw2Pcj-xrU-DMF8KoLr8_EJh2XdTDjaRqz859p0RJ1gPLovsQ8N1HeqeQXKi2mwDJe2rZhWILHdWZ1zmduCY5fF8jUYyBIqLRh1B44L_CBlgeEejKoJfw7V3WoZhxdLeW8SlS2PQ7kN0XIyzm-_TPq1j5QnNHRnXRIh8V7o9rBtdM7PVGEFTpzb1jC6bZ3W-aHuZEWk5e1kRTV8IOXiLf-xtPQ42Hn4j2F27mDg0h2PsgVWmNjr2eqc9y0izps-rmoXHnzmBzvbtGS2yytEFw_WAA"
	TestUserId = "MetalLover666"
)

var TestRatingDao = model.RatingDao{
	ArtistName:   "band",
	Comment:      "someComment",
	FestivalName: "Wacken",
	Rating:       aws.Int(1),
	Year:         aws.Int(123),
}

var TestRating = model.Rating{
	ArtistName:   "band",
	Comment:      "someComment",
	FestivalName: "Wacken",
	Rating:       1,
	Year:         123,
}

var TestRatingRecord = model.RatingRecord{
	DbKey: model.DbKey{
		PK: fmt.Sprintf("USER#%s", TestUserId),
		SK: fmt.Sprintf("ARTIST#%s", TestRating.ArtistName),
	},
	Type:         model.RatingType,
	ArtistName:   TestRating.ArtistName,
	Comment:      TestRating.Comment,
	FestivalName: TestRating.FestivalName,
	Rating:       TestRating.Rating,
	UserId:       TestUserId,
	Year:         TestRating.Year,
}

type PersistenceHelper struct {
	Dynamo    *dynamodb.Client
	TableName string
}

func NewPersistenceHelper() *PersistenceHelper {
	cfg, err := LocalDefaultConfig()
	if err != nil {
		panic(err)
	}

	dynamo := dynamodb.NewFromConfig(cfg)
	tableName := uuid.NewString()
	createTable(dynamo, tableName)

	err = os.Setenv("TABLE_NAME", tableName)

	return &PersistenceHelper{
		Dynamo:    dynamo,
		TableName: tableName,
	}
}

func LocalDefaultConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("local", "local", "")),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: "http://localhost:8095"}, nil
				},
			),
		),
	)
}

func createTable(dynamo *dynamodb.Client, tableName string) {
	if _, err := dynamo.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       types.KeyTypeRange,
			},
		},
		TableName: aws.String(tableName),
	}); err != nil {
		panic(err)
	}
}
