package main

import (
	"reflect"
	"testing"
)

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

func TestGetInput(t *testing.T) {
	t.Run("gets input from file", func(t *testing.T) {
		want := &Input{
			left:  []int{1, 2, 3, 5},
			right: []int{9, 8, 7, 5},
		}

		got, err := GetInput("input_test.txt")

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})

	t.Run("returns err on bad path", func(t *testing.T) {
		_, err := GetInput("missing_file.txt")

		if err != ErrInputFile {
			t.Fatal("didn't get an error but wanted one")
		}
	})
}
