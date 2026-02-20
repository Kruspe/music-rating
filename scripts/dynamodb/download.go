package main

import (
	"bytes"
	"context"

	"github.com/codeclysm/extract/v4"
	"github.com/kruspe/music-rating/scripts/setup"
)

func main() {
	b, err := setup.DownloadFile("https://d1ni2b6xgvw0s0.cloudfront.net/v2.x/dynamodb_local_latest.tar.gz")
	setup.FailOnError("download dynamodb-local", err)

	err = extract.Gz(context.TODO(), bytes.NewReader(b),
		setup.GolangDynamoDBLocalFolder, func(s string) string {
			return s
		})
	setup.FailOnError("download dynamodb-local", err)
}
