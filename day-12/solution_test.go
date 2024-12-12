package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	input, _ := GetInput("input_minimal.txt")

	t.Run("Get A's", func(t *testing.T) {
		want := []Coord{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		}
		got := input.plotMap.GetCoords("A")
		utils.CheckSlicesHaveSameElements(got, want, t)
	})

	t.Run("Get B's", func(t *testing.T) {
		want := []Coord{
			{0, 1},
			{0, 2},
			{1, 1},
			{1, 2},
		}
		got := input.plotMap.GetCoords("B")
		utils.CheckSlicesHaveSameElements(got, want, t)
	})
}
