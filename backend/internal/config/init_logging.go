package config

import (
	"log/slog"
	"os"
)

func InitLogging() {
	logLevel := slog.LevelInfo
	if l, ok := os.LookupEnv("LOG_LEVEL"); ok {
		switch l {
		case "DEBUG":
			logLevel = slog.LevelDebug
		case "WARN":
			logLevel = slog.LevelWarn
		case "ERROR":
			logLevel = slog.LevelError
		}
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)
}
