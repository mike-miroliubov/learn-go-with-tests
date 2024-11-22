package hello

import "testing"

func TestHello(t *testing.T) {
	// when
	result := Hello("TestUser")

	// then
	expected := "Hello World, TestUser"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
