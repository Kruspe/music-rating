package config

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogging() {
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatal("could not get log level from environment variable")
	}
	log.SetLevel(level)
	log.SetFormatter(&log.JSONFormatter{})
}
