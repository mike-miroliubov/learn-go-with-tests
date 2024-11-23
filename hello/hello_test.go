package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello when no name is provided", func(t *testing.T) {
		// when
		result := Hello("", "")

		// then
		expected := "Hello, World"
		assertEquals(t, result, expected)
	})

	t.Run("Say hello to a user by Name", func(t *testing.T) {
		// when
		result := Hello("TestUser", "")

		// then
		expected := "Hello, TestUser"
		assertEquals(t, result, expected)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {
		// when
		result := Hello("TestUser", "Spanish")

		// then
		expected := "Hola, TestUser"
		assertEquals(t, result, expected)
	})
}

func assertEquals(t *testing.T, result string, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
