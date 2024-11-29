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
		{
			"arrays",
			[2]Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"maps",
			map[string]string{
				"City":    "London",
				"Country": "England",
			},
			[]string{"London", "England"},
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

	//t.Run("channels", func(t *testing.T) {
	//	// given
	//	ch := make(chan Profile)
	//
	//	go func() {
	//		ch <- Profile{Age: 33, City: "Berlin"}
	//		ch <- Profile{Age: 34, City: "Katowice"}
	//		close(ch)
	//	}()
	//
	//	// when
	//	var calls []string
	//	walk(ch, func(input string) {
	//		calls = append(calls, input)
	//	})
	//
	//	// then
	//	assert.Equal(t, []string{"Berlin", "Katowice"}, calls)
	//})

}
