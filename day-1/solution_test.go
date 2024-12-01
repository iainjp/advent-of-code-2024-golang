package main

import "testing"

func TestSumDistances(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	want := 11

	result := SumDistances(left, right)

	if result != want {
		t.Fatalf("wanted 11, got %v", result)
	}
}
