package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to the yourself", func(t *testing.T) {
		got := Hello("yourself", "")
		want := "Hello, yourself"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying `Hello, world` when a empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Want: %q, got: %q", want, got)
	}
}
