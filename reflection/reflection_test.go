package reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"struct with 2 string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non sting field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"struct with nested fields",
			Person{
				Name:    "Chris",
				Profile: Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointer to struct",
			&Person{
				Name:    "Chris",
				Profile: Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			// when
			var calls []string
			walk(test.input, func(input string) {
				calls = append(calls, input)
			})
			// then
			assert.Equal(t, test.expectedCalls, calls)
		})
	}

}
