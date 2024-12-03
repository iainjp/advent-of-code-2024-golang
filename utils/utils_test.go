package utils

import (
	"reflect"
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
			return calls%2 == 0
		}

		inputs := []int{1, 2, 3}

		result := All(inputs, fn)

		CheckEquals(result, false, t)
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

		CheckEquals(calls, 1, t)
		CheckEquals(result, true, t)
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
	})
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

func CheckEquals[T any](got, want T, t testing.TB) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
