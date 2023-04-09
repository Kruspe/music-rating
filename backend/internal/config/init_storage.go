package config

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
)

func InitStorage(cfg aws.Config) *persistence.FestivalStorage {
	return persistence.NewFestivalStorage(s3.NewFromConfig(cfg))
}
