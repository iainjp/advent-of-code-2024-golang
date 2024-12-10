package utils

import (
	"reflect"
	"slices"
	"testing"
)

func TestAbs(t *testing.T) {
	t.Run("handles positive numbers", func(t *testing.T) {
		CheckEquals(Abs(7), 7, t)
	})

	t.Run("handles negative numbers", func(t *testing.T) {
		CheckEquals(Abs(-6), 6, t)
	})
}

func TestAll(t *testing.T) {
	t.Run("all predicate responses true, returns true", func(t *testing.T) {
		calls := 0

		var fn = func(t int) bool {
			calls += 1
			return true
		}

		inputs := []int{1, 2, 3}

		result := All(inputs, fn)

		CheckEquals(calls, 3, t)
		CheckEquals(result, true, t)
	})

	t.Run("not all predicate responses true, returns false", func(t *testing.T) {
		calls := 0

		var fn = func(t int) bool {
			calls += 1
			return t == 0 || calls == 2
		}

		inputs := []int{0, 2, 4, 4}

		result := All(inputs, fn)

		CheckEquals(result, false, t)
		CheckEquals(calls, 3, t)
	})

}

func TestAny(t *testing.T) {
	t.Run("no predicate returns true, return false", func(t *testing.T) {
		calls := 0

		var fn = func(t int) bool {
			calls += 1
			return true
		}

		inputs := []int{1, 2, 3}

		result := Any(inputs, fn)

		CheckEquals(result, true, t)
		CheckEquals(calls, 1, t)
	})

	t.Run("no predicate responses true, returns false", func(t *testing.T) {
		calls := 0

		var fn = func(t int) bool {
			calls += 1
			return calls%2 == 0
		}

		inputs := []int{1, 2, 3}

		result := All(inputs, fn)

		CheckEquals(result, false, t)
		CheckEquals(calls, 1, t)
	})
}

func TestFilter(t *testing.T) {
	calls := 0
	input := []int{1, 2, 3}

	var fn = func(t int) bool {
		calls++
		return t%2 == 0
	}

	want := []int{2}

	got := Filter(input, fn)

	CheckEquals(got, want, t)
	CheckEquals(calls, 3, t)
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	fn := func(i int) int {
		return i * 3
	}

	want := []int{
		3, 6, 9,
	}

	got := Map(input, fn)

	CheckEquals(got, want, t)
}

func TestCountOccurences(t *testing.T) {
	input := []bool{
		true,
		false,
		true,
	}

	want := map[bool]int{
		true:  2,
		false: 1,
	}

	got := CountOccurences(input)

	CheckEquals(got, want, t)
}

func TestIterToSlice(t *testing.T) {
	input := slices.Backward([]int{1, 2, 3})

	got := IterToSlice(input)

	CheckEquals(got[0], 3, t)
	CheckEquals(got[1], 2, t)
	CheckEquals(got[2], 1, t)
}

func CheckEquals[T any](got, want T, t testing.TB) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
