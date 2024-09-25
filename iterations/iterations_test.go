package main

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Testing repeat function", func(t *testing.T) {
		got := Repeat("a")
		expected := "aaaa"

		if got != expected {
			t.Errorf("got: %q, expected: %q", got, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
