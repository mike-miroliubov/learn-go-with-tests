package iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	t.Run("Must repeat a char 5 times", func(t *testing.T) {
		// when
		var result = Repeat("a", 5)

		// then
		assert.Equal(t, "aaaaa", result)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
