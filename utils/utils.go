package utils

import (
	"iter"
	"reflect"
	"slices"
	"testing"
)

func Abs(v int) int {
	return max(v, -v)
}

func All[T any](ts []T, predicate func(T) bool) bool {
	for _, t := range ts {
		if !predicate(t) {
			return false
		}
	}
	return true
}

func Any[T any](ts []T, predicate func(T) bool) bool {
	for _, t := range ts {
		if predicate(t) {
			return true
		}
	}
	return false
}

func Map[T any, R any](ts []T, fn func(T) R) []R {
	var results []R
	for _, t := range ts {
		result := fn(t)
		results = append(results, result)
	}
	return results
}

func Filter[T any](ts []T, fn func(T) bool) []T {
	var results []T
	for _, t := range ts {
		if fn(t) {
			results = append(results, t)
		}
	}
	return results
}

func CountOccurences[T comparable](ts []T) map[T]int {
	counts := make(map[T]int)
	for _, num := range ts {
		counts[num] = counts[num] + 1
	}
	return counts
}

func CheckEqual[K any](got, want K, t testing.TB) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func CheckContains[K comparable, E ~[]K](slice E, item K, t testing.TB) {
	t.Helper()
	if !slices.Contains(slice, item) {
		t.Fatalf("%v does not contain %v", slice, item)
	}
}

func CheckNotNil[K any](got *K, t testing.TB) {
	t.Helper()
	if got == nil {
		t.Fatal("expected not nil, got nil")
	}
}

func CheckSlicesHaveSameElements[S ~[]E, E comparable](s1 S, s2 S, t testing.TB) {
	t.Helper()
	CheckEqual(len(s1), len(s2), t)
	for _, e := range s1 {
		if !slices.Contains(s2, e) {
			t.Fatalf("slice does not contain required element %v", e)
		}
	}
}

func IterToSlice[E any, B any](it iter.Seq2[B, E]) []E {
	var es []E
	for _, ei := range it {
		es = append(es, ei)
	}
	return es
}
