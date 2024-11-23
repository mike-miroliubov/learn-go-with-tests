package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	// when
	sum := Add(2, 2)

	// then
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 4)
	fmt.Println(sum)

	// Output: 5
}
