package util

import (
	"net/http"
	"time"
)

func NewHttpClient(timeout time.Duration) *http.Client {
	// set default value
	if timeout == 0 {
		timeout = time.Minute
	}
	client := &http.Client{
		Timeout: timeout,
	}
	return client
}
