package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"log/slog"
)

func InitAwsConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		slog.Error("error setting up AwsConfig", slog.Any("error", err))
		panic(err)
	}
	return cfg
}
