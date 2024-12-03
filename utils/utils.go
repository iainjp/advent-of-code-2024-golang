package utils

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

func CountOccurences[T comparable](ts []T) map[T]int {
	counts := make(map[T]int)
	for _, num := range ts {
		counts[num] = counts[num] + 1
	}
	return counts
}
