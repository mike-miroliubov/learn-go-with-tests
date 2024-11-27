package selector

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Should find the faster one", func(t *testing.T) {
		// given
		// HandlerFunc is a function type - an alias over a function
		slowServer := createServerWithDelay(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := createServerWithDelay(5 * time.Millisecond)
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		// when
		result, err := Racer(slowUrl, fastUrl)

		// then
		assert.Equal(t, fastUrl, result)
		assert.NoError(t, err)
	})

	t.Run("Should fail with a timeout", func(t *testing.T) {
		// given
		// HandlerFunc is a function type - an alias over a function
		slowServer := createServerWithDelay(11 * time.Millisecond)
		defer slowServer.Close()
		fastServer := createServerWithDelay(12 * time.Millisecond)
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		// when
		_, err := ConfigurableRacer(slowUrl, fastUrl, 10*time.Millisecond)

		// then
		assert.Error(t, err)
	})
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}
