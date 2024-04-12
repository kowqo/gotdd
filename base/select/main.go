package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, time.Second*10)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {

	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("lonng time")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
