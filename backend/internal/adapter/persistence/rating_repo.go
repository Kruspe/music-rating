package persistence

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kruspe/music-rating/internal/model"
)

type RatingRecord struct {
	DbKey
	Type         string `dynamodbav:"type"`
	ArtistName   string `dynamodbav:"artist_name"`
	Comment      string `dynamodbav:"comment,omitempty"`
	FestivalName string `dynamodbav:"festival_name,omitempty"`
	Rating       int    `dynamodbav:"rating"`
	UserId       string `dynamodbav:"user_id"`
	Year         int    `dynamodbav:"year,omitempty"`
}

type RatingRepo struct {
	dynamo    *dynamodb.Client
	tableName string
}

func NewRatingRepo(dynamo *dynamodb.Client, tableName string) *RatingRepo {
	return &RatingRepo{
		dynamo:    dynamo,
		tableName: tableName,
	}
}

func (r *RatingRepo) Create(ctx context.Context, userId string, rating model.Rating) error {
	record := RatingRecord{
		DbKey: DbKey{
			PK: fmt.Sprintf("USER#%s", userId),
			SK: fmt.Sprintf("ARTIST#%s", rating.ArtistName),
		},
		Type:         RatingType,
		ArtistName:   rating.ArtistName,
		Comment:      rating.Comment,
		FestivalName: rating.FestivalName,
		Rating:       rating.Rating,
		UserId:       userId,
		Year:         rating.Year,
	}
	item, err := attributevalue.MarshalMap(record)
	if err != nil {
		return err
	}
	_, err = r.dynamo.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.tableName),
	})
	return err
}

func (r *RatingRepo) GetAll(ctx context.Context, userId string) ([]model.Rating, error) {
	expr, err := expression.NewBuilder().WithKeyCondition(expression.KeyEqual(expression.Key("PK"), expression.Value(fmt.Sprintf("USER#%s", userId)))).Build()
	if err != nil {
		return nil, err
	}
	items, err := r.dynamo.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(r.tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		return nil, err
	}
	var ratings []RatingRecord
	err = attributevalue.UnmarshalListOfMaps(items.Items, &ratings)
	if err != nil {
		return nil, err
	}

	var result []model.Rating
	for _, r := range ratings {
		result = append(result, toRating(r))
	}
	return result, nil
}

func toRating(r RatingRecord) model.Rating {
	return model.Rating{
		ArtistName:   r.ArtistName,
		Comment:      r.Comment,
		FestivalName: r.FestivalName,
		Rating:       r.Rating,
		Year:         r.Year,
	}
}
