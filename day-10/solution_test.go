package main

import (
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
		score := make(map[*Node]int)

		walkResult := input.graph.Walk()

		for _, r := range walkResult {
			s := score[r.head] + 1
			score[r.head] = s
		}

		var got []int
		for _, v := range score {
			got = append(got, v)
		}

		want := []int{5, 6, 5, 3, 1, 3, 5, 3, 5}

		utils.CheckSlicesHaveSameElements(got, want, t)
	})

	t.Run("WalkNonUnique(), check scores", func(t *testing.T) {
		input, _ := GetInput("input_test.txt")

		walkResult := input.graph.WalkNonUnique()

		var got []int
		for _, finishes := range walkResult {
			got = append(got, finishes)
		}

		want := []int{20, 24, 10, 4, 1, 4, 5, 8, 5}

		utils.CheckSlicesHaveSameElements(got, want, t)
	})
}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_test.txt")
	want := 36

	got := Part1(input)

	utils.CheckEqual(got, want, t)
}

func TestPart2(t *testing.T) {
	input, _ := GetInput("input_test.txt")
	want := 81

	got := Part2(input)

	utils.CheckEqual(got, want, t)
}
