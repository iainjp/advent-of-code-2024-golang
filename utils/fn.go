package utils

import (
	"iter"
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

func First[T any](ts []*T, fn func(*T) bool) (int, *T) {
	for i, t := range ts {
		if fn(t) {
			return i, t
		}
	}
	return -1, nil
}

func CountOccurences[T comparable](ts []T) map[T]int {
	counts := make(map[T]int)
	for _, num := range ts {
		counts[num] = counts[num] + 1
	}
	return counts
}

func IterSeqToSlice[E any](it iter.Seq[E]) []E {
	var es []E
	for ei := range it {
		es = append(es, ei)
	}
	return es
}

func IterToSlice[E any, B any](it iter.Seq2[B, E]) []E {
	var es []E
	for _, ei := range it {
		es = append(es, ei)
	}
	return es
}
