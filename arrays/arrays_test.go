package arrays

import (
	"reflect"
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

func TestMultiArraySum(t *testing.T) {
	// when
	result := MultiArraySum([]int{1, 3}, []int{3, 5})

	// then
	assert.Equal(t, []int{4, 8}, result)

	if !reflect.DeepEqual([]int{4, 8}, result) {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestArrayTailSum(t *testing.T) {
	t.Run("Sum tails", func(t *testing.T) {
		// when
		result := ArrayTailSum([]int{1, 2, 3}, []int{0, 9, 1})

		// then
		assert.Equal(t, []int{5, 10}, result)
	})

	t.Run("sum tails of empty slices", func(t *testing.T) {
		// when
		result := ArrayTailSum([]int{}, []int{})

		// then
		assert.Equal(t, []int{0, 0}, result)
	})
}
