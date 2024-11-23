package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello when no name is provided", func(t *testing.T) {
		// when
		result := Hello("")

		// then
		expected := "Hello, World"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("Say hello to a user by Name", func(t *testing.T) {
		// when
		result := Hello("TestUser")

		// then
		expected := "Hello, TestUser"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}
