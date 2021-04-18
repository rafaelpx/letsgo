package greetings

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Rafa")
		expected := "Hello, Rafa"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		expected := "Hello, World"

		assertCorrectMessage(t, got, expected)
	})
}
