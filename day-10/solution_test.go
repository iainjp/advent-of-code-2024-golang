package main

import (
	"maps"
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	t.Run("graph contains 9 trailheads", func(t *testing.T) {
		got, _ := GetInput("input_test.txt")

		utils.CheckEqual(len(got.graph.trailheads), 9, t)
		for _, th := range got.graph.trailheads {
			utils.CheckEqual(th.height, 0, t)
		}
	})

	t.Run("Walk(), check scores", func(t *testing.T) {
		input, _ := GetInput("input_test.txt")

		walkResult := input.graph.Walk()

		want := []int{5, 6, 5, 3, 1, 3, 5, 3, 5}
		gotIter := maps.Values(walkResult)

		var got []int
		for g := range gotIter {
			got = append(got, g)
		}

		utils.CheckEqual(len(got), len(want), t)
	})
}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_test.txt")
	want := 36

	got := Part1(input)

	utils.CheckEqual(got, want, t)
}
