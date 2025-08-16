package select_go

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(1 * time.Second):
		return "", fmt.Errorf("timed out, tired of waiting")
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})
	go func() {
		http.Get(url)
		close(channel)
	}()

	return channel
}
