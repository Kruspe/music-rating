package setup

import (
	"errors"
	"net/http"
	"time"
)

const startupSleepBetweenRetries = 100 * time.Millisecond
const startupRetries = 200

func AwaitHttpServiceStartup(endpoint string) error {
	var retries = 0
	for {
		resp, err := http.Get(endpoint)
		if err != nil && retries > startupRetries {
			return errors.New("startup failed")
		} else if resp != nil {
			return nil
		} else {
			time.Sleep(startupSleepBetweenRetries)
			retries++
		}
	}
}
