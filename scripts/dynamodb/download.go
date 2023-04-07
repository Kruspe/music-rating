package main

import (
	"bytes"
	"context"
	"github.com/codeclysm/extract/v3"
	"github.com/kruspe/music-rating/scripts/setup"
)

func main() {
	b, err := setup.DownloadFile("https://s3.eu-central-1.amazonaws.com/dynamodb-local-frankfurt/dynamodb_local_latest.tar.gz")
	setup.FailOnError("download dynamodb-local", err)

	err = extract.Gz(context.TODO(), bytes.NewReader(b),
		setup.GolangDynamoDBLocalFolder, func(s string) string {
			return s
		})
	setup.FailOnError("download dynamodb-local", err)
}
