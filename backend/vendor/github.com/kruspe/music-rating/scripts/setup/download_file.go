package setup

import (
	"fmt"
	"io"
	"net/http"
)

func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error on downloading file, expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
