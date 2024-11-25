package dependency_injection

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	// given
	buffer := bytes.Buffer{}

	// when
	Greet(&buffer, "Chris")

	// then
	assert.Equal(t, "Hello, Chris", buffer.String())
}
