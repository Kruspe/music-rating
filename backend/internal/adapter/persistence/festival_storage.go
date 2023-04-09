package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kruspe/music-rating/internal/model"
	"os"
)

type S3Client interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type ArtistRecord struct {
	Artist string `json:"artist"`
	Image  string `json:"image"`
}

type FestivalStorage struct {
	s3 S3Client
}

func NewFestivalStorage(s3Client S3Client) *FestivalStorage {
	return &FestivalStorage{
		s3: s3Client,
	}
}

func (s *FestivalStorage) GetArtists(ctx context.Context, festivalName string) ([]model.Artist, error) {
	object, err := s.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("FESTIVAL_ARTIST_BUCKET_NAME")),
		Key:    aws.String(fmt.Sprintf("%s.json", festivalName)),
	})
	if err != nil {
		return nil, err
	}

	var artists []ArtistRecord
	if err := json.NewDecoder(object.Body).Decode(&artists); err != nil {
		return nil, err
	}

	result := make([]model.Artist, 0)
	for _, artist := range artists {
		result = append(result, model.Artist{
			Name:     artist.Artist,
			ImageUrl: artist.Image,
		})
	}
	return result, nil
}
