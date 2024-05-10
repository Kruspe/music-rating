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
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/model"
	"github.com/kruspe/music-rating/scripts/setup"
	"io"
	"strings"
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
}

func NewPersistenceHelper() *PersistenceHelper {
	cfg, err := LocalDefaultConfig()
	if err != nil {
		panic(err)
	}
	dynamo := dynamodb.NewFromConfig(cfg)
	tableName := uuid.NewString()
	createTable(dynamo, tableName)
	return &PersistenceHelper{
		Dynamo:    dynamo,
		TableName: tableName,
	}
}

func (ph *PersistenceHelper) MockFestivals(festivals map[string][]model.Artist) persistence.S3Client {
	return MockS3Client{
		GetObjectMock: func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
			festivalName := strings.Split(*params.Key, ".json")[0]
			artists, ok := festivals[festivalName]
			if !ok {
				return nil, fmt.Errorf("festival %s not found", festivalName)
			}
			result := make([]persistence.ArtistRecord, 0)
			for _, artist := range artists {
				result = append(result, persistence.ArtistRecord{
					Artist: artist.Name,
					Image:  artist.ImageUrl,
				})
			}
			s3Body, err := json.Marshal(result)
			if err != nil {
				panic("s3 mock failed")
			}

			return &s3.GetObjectOutput{
				Body: io.NopCloser(bytes.NewReader(s3Body)),
			}, nil
		},
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
