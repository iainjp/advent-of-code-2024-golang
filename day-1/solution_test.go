package main

import "testing"

func TestSumDistances(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	want := 11

	got := SumDistances(left, right)

	if got != want {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestSimilarityScore(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	want := 31

	got := SimilarityScore(left, right)

	if got != want {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
