package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world!")
}

func abs(v int) int {
	return max(v, -v)
}

func SumDistances(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for index, leftNum := range left {
		distance := abs(leftNum - right[index])
		total += distance
	}

	return total
}
