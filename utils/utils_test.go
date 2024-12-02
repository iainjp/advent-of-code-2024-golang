package utils

import "testing"

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

func CheckEquals[K comparable](got, want K, t testing.TB) {
	t.Helper()

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}
