package reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalk(t *testing.T) {
	// given
	cases := []struct {
		name          string
		input         interface{}
		expectedCalls []string
	}{
		{
			"struct with 1 string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		// when
		var calls []string
		walk(test.input, func(input string) {
			calls = append(calls, input)
		})
		// then
		assert.Equal(t, test.expectedCalls, calls)
	}

}
