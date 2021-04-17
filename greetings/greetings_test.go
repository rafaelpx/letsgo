package greetings

import "testing"

func TestHello(t *testing.T) {
	t.Run("greets GoCon", func(t *testing.T) {
		got := Hello("GoCon")
		expected := "Hello GoCon!"

		if got != expected {
			t.Errorf("Got: %s, Expected: %s", got, expected)
		}
	})
}
