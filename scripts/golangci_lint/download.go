package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/codeclysm/extract/v3"
	"github.com/kruspe/music-rating/scripts/setup"
	"runtime"
	"strings"
)

func main() {
	b, err := setup.DownloadFile(fmt.Sprintf("https://github.com/golangci/golangci-lint/releases/download/v%s/golangci-lint-%s-%s-%s.tar.gz", setup.GolangCiLintVersion, setup.GolangCiLintVersion, runtime.GOOS, runtime.GOARCH))
	setup.FailOnError("download golangci-lint", err)

	err = extract.Gz(context.TODO(), bytes.NewReader(b), setup.GolangCiLintFolder, func(s string) string {
		parts := strings.Split(s, "/")
		return parts[len(parts)-1]
	})
	setup.FailOnError("download golangci-lint", err)
}
