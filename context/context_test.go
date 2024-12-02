package context

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	// given
	testData := "hello"
	server := Server(&SpyStore{testData})

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	// when
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// then
	assert.Equal(t, testData, response.Body.String())
}
