package utils

import (
	"reflect"
	"slices"
	"testing"
)

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

func CheckSlicesHaveSameElements[S ~[]E, E comparable](got S, want S, t testing.TB) {
	t.Helper()
	CheckEqual(len(got), len(want), t)
	for _, e := range want {
		if !slices.Contains(got, e) {
			t.Fatalf("<got> does not contain required element from <want>: %v", e)
		}
	}
}
