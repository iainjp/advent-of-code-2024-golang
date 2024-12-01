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

func TestRemoveAfterDontUntilDoOrEnd(t *testing.T) {
	input := "HELLOdon't()WORLDdo()TESTdon't()AGAIN"
	want := "HELLOTEST"

	got := RemoveAfterDontUntilDoOrEnd(input)

	CheckEqual(got, want, t)
}

func TestPart2(t *testing.T) {
	t.Run("based on puzzle input", func(t *testing.T) {
		input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))don't()mul(7,3)"
		want := 48

		got := Part2(input)

		CheckEqual(got, want, t)
	})

	t.Run("my own input", func(t *testing.T) {
		input := "do()xtjh76mul(3,4)gdon't()mul(10,10)do()mul(5,4)don't()mul(10,10)"
		want := 32

		got := Part2(input)

		CheckEqual(got, want, t)
	})
}

func CheckEqual[K any](got, want K, t testing.TB) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
