package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	expected := 3

	if got != expected {
		t.Errorf("got: %d, expected: %d", got, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
