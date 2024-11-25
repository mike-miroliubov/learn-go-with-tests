package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraySum(t *testing.T) {
	t.Run("Sum an array of numbers", func(t *testing.T) {
		// given
		array := []int{1, 2, 3}

		// when
		result := ArraySum(array)

		// then
		assert.Equal(t, 6, result)
	})
}
