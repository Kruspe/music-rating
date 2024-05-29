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
		DbKey:        ratingDbKey(userId, rating.ArtistName),
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

func (r *RatingRepo) GetAll(ctx context.Context, userId string) (*model.Ratings, error) {
	expr, err := expression.NewBuilder().WithKeyCondition(expression.KeyEqual(expression.Key("PK"), expression.Value(fmt.Sprintf("USER#%s", userId)))).Build()
	if err != nil {
		return nil, err
	}
	paginator := dynamodb.NewQueryPaginator(r.dynamo, &dynamodb.QueryInput{
		TableName:                 aws.String(r.tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	var ratings []RatingRecord
	for paginator.HasMorePages() {
		items, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		var r []RatingRecord
		err = attributevalue.UnmarshalListOfMaps(items.Items, &r)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, r...)
	}

	result := &model.Ratings{
		Keys:   make([]string, len(ratings)),
		Values: make(map[string]model.Rating),
	}
	for i, record := range ratings {
		rating, err := strconv.ParseFloat(record.Rating, 32)
		if err != nil {
			return nil, err
		}
		result.Keys[i] = record.ArtistName
		result.Values[record.ArtistName] = model.Rating{
			ArtistName:   record.ArtistName,
			Comment:      record.Comment,
			FestivalName: record.FestivalName,
			Rating:       rating,
			Year:         record.Year,
		}
	}
	return result, nil
}

func (r *RatingRepo) Update(ctx context.Context, userId string, rating model.Rating) error {
	updateExpr := expression.Set(expression.Name("rating"), expression.Value(rating.Rating))
	if rating.Year == 0 {
		updateExpr.Remove(expression.Name("year"))
	} else {
		updateExpr.Set(expression.Name("year"), expression.Value(rating.Year))
	}
	if rating.FestivalName == "" {
		updateExpr.Remove(expression.Name("festival_name"))
	} else {
		updateExpr.Set(expression.Name("festival_name"), expression.Value(rating.FestivalName))
	}
	if rating.Comment == "" {
		updateExpr.Remove(expression.Name("comment"))
	} else {
		updateExpr.Set(expression.Name("comment"), expression.Value(rating.Comment))
	}
	expr, err := expression.NewBuilder().WithUpdate(updateExpr).WithCondition(expression.And(
		expression.Equal(expression.Name("PK"), expression.Value(fmt.Sprintf("USER#%s", userId))),
		expression.Equal(expression.Name("SK"), expression.Value(fmt.Sprintf("ARTIST#%s", rating.ArtistName))),
	)).Build()
	if err != nil {
		return err
	}
	key, err := attributevalue.MarshalMap(ratingDbKey(userId, rating.ArtistName))
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
		return &model.UpdateNonExistingRatingError{ArtistName: rating.ArtistName}
	}
	return err
}
