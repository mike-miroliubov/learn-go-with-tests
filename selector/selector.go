package selector

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, time.Duration(10*time.Second))
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (string, error) {
	select { /// waits on multiple channels
	case <-call(url1):
		return url1, nil
	case <-call(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s for %s and %s", timeout, url1, url2)
	}
}

func measureTime(call func()) time.Duration {
	start1 := time.Now()
	call()
	return time.Since(start1)
}

// return channel of any type, interface{} is an Any/Object type in Go
func call(url string) chan any {
	ch := make(chan any)
	go func() {
		http.Get(url)
		close(ch) // sends a signal into the channel
	}()
	return ch
}
