package config

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func InitAwsConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		slog.Error("error setting up AwsConfig", slog.Any("error", err))
		panic(err)
	}
	return cfg
}
