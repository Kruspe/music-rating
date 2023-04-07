package persistence_test_helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/kruspe/music-rating/internal/adapter/model"
	"github.com/kruspe/music-rating/internal/adapter/model/model_test_helper"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/scripts/setup"
	"io"
	"os"
)

type MockS3Client struct {
	GetObjectMock func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

func (m MockS3Client) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return m.GetObjectMock(ctx, params, optFns...)
}

type PersistenceHelper struct {
	Dynamo    *dynamodb.Client
	TableName string
	S3Mock    persistence.S3Client
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
	if err != nil {
		panic(err)
	}

	s3Mock := func() persistence.S3Client {
		artist1 := model_test_helper.AnArtistWithName("Bloodbath")
		artist2 := model_test_helper.AnArtistWithName("Hypocrisy")
		artist3 := model_test_helper.AnArtistWithName("Benediction")
		return MockS3Client{
			GetObjectMock: func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
				s3Body, err := json.Marshal([]model.ArtistRecord{
					{
						Artist: artist1.ArtistName,
						Image:  artist1.ImageUrl,
					},
					{
						Artist: artist2.ArtistName,
						Image:  artist2.ImageUrl,
					},
					{
						Artist: artist3.ArtistName,
						Image:  artist3.ImageUrl,
					},
				})
				if err != nil {
					panic("s3 mock failed")
				}

				return &s3.GetObjectOutput{
					Body: io.NopCloser(bytes.NewReader(s3Body)),
				}, nil
			},
		}
	}

	return &PersistenceHelper{
		Dynamo:    dynamo,
		TableName: tableName,
		S3Mock:    s3Mock(),
	}
}

func LocalDefaultConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("local", "local", "")),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: fmt.Sprintf("http://localhost:%d", setup.DynamoDBPort)}, nil
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
