package persistence

import (
	"backend/internal/adapter/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

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

func (r *RatingRepo) SaveRating(ctx context.Context, userId string, rating model.Rating) error {
	record := model.RatingRecord{
		DbKey: model.DbKey{
			PK: fmt.Sprintf("USER#%s", userId),
			SK: fmt.Sprintf("ARTIST#%s", rating.ArtistName),
		},
		Type:         model.RatingType,
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
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingRepo) GetRatings(ctx context.Context, userId string) ([]model.RatingRecord, error) {
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
	var ratings []model.RatingRecord
	err = attributevalue.UnmarshalListOfMaps(items.Items, &ratings)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}
