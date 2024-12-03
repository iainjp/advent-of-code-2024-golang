package main

import (
	"reflect"
	"testing"
)

func TestGetMatches(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	want := []Match{
		{2, 4},
		{5, 5},
		{11, 8},
		{8, 5},
	}

	got := GetMatches(input)

	CheckEqual(got, want, t)
}

func TestEvaluateMatches(t *testing.T) {
	input := []Match{
		{2, 4},
		{5, 5},
		{11, 8},
		{8, 5},
	}

	want := 161

	got := EvaluateMatches(input)

	CheckEqual(got, want, t)
}

func CheckEqual[K any](got, want K, t testing.TB) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
