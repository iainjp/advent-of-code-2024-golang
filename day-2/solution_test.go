package main

import (
	"reflect"
	"testing"
)

func TestGetInput(t *testing.T) {
	want := &Input{reports: [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}}

	got, err := GetInput("input_test.txt")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	CheckEqual(got, want, t)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestIsIncreasingOrDecreasing(t *testing.T) {
	t.Run("is increasing returns true", func(t *testing.T) {
		levels := []int{1, 2, 3}
		want := true

		got := IsIncreasingOrDecreasing(levels)

		CheckEqual(got, want, t)
	})

	t.Run("is decreasing returns true", func(t *testing.T) {
		levels := []int{3, 2, 1}
		want := true

		got := IsIncreasingOrDecreasing(levels)

		CheckEqual(got, want, t)
	})

	t.Run("neither increasing or decreasing returns false", func(t *testing.T) {
		levels := []int{5, 6, 1}
		want := false

		got := IsIncreasingOrDecreasing(levels)

		CheckEqual(got, want, t)
	})
}

func TestReportSafetyCheck(t *testing.T) {
	input := &Input{reports: [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}}

	want := []bool{
		true,
		false,
		false,
		false,
		false,
		true,
	}

	got := ReportSafetyCheck(input.reports)

	CheckEqual(got, want, t)
}

func TestReportSafetyCheckWithTolerance(t *testing.T) {
	input := &Input{reports: [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}}

	want := []bool{
		true,
		false,
		false,
		true,
		true,
		true,
	}

	got := ReportSafetyCheckWithTolerance(input.reports)

	CheckEqual(got, want, t)
}

func TestGetDistances(t *testing.T) {
	input := []int{1, 2, 4}
	want := []int{1, 2}

	got := GetDistances(input)

	CheckEqual(got, want, t)
}

func TestGetPermutationsWithOneMissing(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := [][]int{
		{2, 3, 4},
		{1, 3, 4},
		{1, 2, 4},
		{1, 2, 3},
	}

	got := GetPermutationsWithOneMissing(input)

	CheckEqual(got, want, t)
}

func CheckEqual[K any](got, want K, t testing.TB) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
