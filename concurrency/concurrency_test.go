package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	// given
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	// when
	result := CheckWebsites(mockWebsiteChecker, websites)

	// then
	expected := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	assert.Equal(t, expected, result)
}

func slowWebsitesChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 20)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsitesChecker, urls)
	}
}
