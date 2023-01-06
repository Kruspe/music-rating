package persistence

import (
	"backend/internal/adapter/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

type S3Client interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type Storage struct {
	s3 S3Client
}

func NewFestivalStorage(s3Client S3Client) *Storage {
	return &Storage{
		s3: s3Client,
	}
}

func (s *Storage) GetFestival(ctx context.Context, festivalName string) (model.Festival, error) {
	object, err := s.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("FESTIVAL_ARTIST_BUCKET_NAME")),
		Key:    aws.String(fmt.Sprintf("%s.json", festivalName)),
	})
	if err != nil {
		return model.Festival{}, err
	}

	var artists []model.ArtistRecord
	if err := json.NewDecoder(object.Body).Decode(&artists); err != nil {
		return model.Festival{}, err
	}

	var festival model.Festival
	for _, artist := range artists {
		festival.Artists = append(festival.Artists, model.Artist{
			ArtistName: artist.Artist,
			ImageUrl:   artist.Image,
		})
	}
	return festival, nil
}
