package mocks

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockSleeper struct {
	Calls int
}

func (s *MockSleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	// given
	buffer := &bytes.Buffer{}
	mockSleeper := MockSleeper{}

	// when
	Countdown(buffer, &mockSleeper)

	// then
	expected := `3
2
1
Go!`
	assert.Equal(t, expected, buffer.String())
	assert.Equal(t, 3, mockSleeper.Calls)
}
