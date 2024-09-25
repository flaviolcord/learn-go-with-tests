package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 3, 5, 2}
		got := Sum(numbers)
		expected := 11

		if got != expected {
			t.Errorf("got: %d, expected: %d, numbers: %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}

	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(got, want)
	})

	t.Run("evaluate sum safety empty numbers list", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2}, []int{0, 9})
		want := []int{0, 2, 9}

		checkSums(got, want)
	})

}
