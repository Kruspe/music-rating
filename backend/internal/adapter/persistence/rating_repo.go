package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kruspe/music-rating/internal/model"
	"strconv"
)

type RatingRecord struct {
	DbKey
	Type         string `dynamodbav:"type"`
	ArtistName   string `dynamodbav:"artist_name"`
	Comment      string `dynamodbav:"comment,omitempty"`
	FestivalName string `dynamodbav:"festival_name,omitempty"`
	Rating       string `dynamodbav:"rating"`
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

func (r *RatingRepo) Save(ctx context.Context, userId string, rating model.Rating) error {
	record := RatingRecord{
		DbKey: DbKey{
			PK: fmt.Sprintf("USER#%s", userId),
			SK: fmt.Sprintf("ARTIST#%s", rating.ArtistName),
		},
		Type:         RatingType,
		ArtistName:   rating.ArtistName,
		Comment:      rating.Comment,
		FestivalName: rating.FestivalName,
		Rating:       strconv.FormatFloat(rating.Rating, 'f', 1, 32),
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
		rating, err := strconv.ParseFloat(r.Rating, 32)
		if err != nil {
			return nil, err
		}
		result = append(result, model.Rating{
			ArtistName:   r.ArtistName,
			Comment:      r.Comment,
			FestivalName: r.FestivalName,
			Rating:       rating,
			Year:         r.Year,
		})
	}
	return result, nil
}

func (r *RatingRepo) Update(ctx context.Context, userId, artistName string, update model.RatingUpdate) error {
	expr, err := expression.NewBuilder().WithUpdate(
		expression.Set(expression.Name("festival_name"), expression.Value(update.FestivalName)).
			Set(expression.Name("rating"), expression.Value(update.Rating)).
			Set(expression.Name("year"), expression.Value(update.Year)).
			Set(expression.Name("comment"), expression.Value(update.Comment)),
	).WithCondition(expression.And(
		expression.Equal(expression.Name("PK"), expression.Value(fmt.Sprintf("USER#%s", userId))),
		expression.Equal(expression.Name("SK"), expression.Value(fmt.Sprintf("ARTIST#%s", artistName))),
	)).Build()
	if err != nil {
		return err
	}
	key, err := attributevalue.MarshalMap(DbKey{
		PK: fmt.Sprintf("USER#%s", userId),
		SK: fmt.Sprintf("ARTIST#%s", artistName),
	})
	if err != nil {
		return err
	}
	_, err = r.dynamo.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 aws.String(r.tableName),
		UpdateExpression:          expr.Update(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		ConditionExpression:       expr.Condition(),
	})
	var conditionalError *types.ConditionalCheckFailedException
	if errors.As(err, &conditionalError) {
		return model.UpdateNonExistingRatingError{ArtistName: artistName}
	}
	return err
}
