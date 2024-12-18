package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	input, _ := GetInput("input_example.txt")

	utils.CheckNotNil(input, t)
	utils.CheckNotNil(input.maze.start, t)
	utils.CheckNotNil(input.maze.end, t)
	utils.CheckEqual(len(input.maze.all), 104, t)
}

func TestPart1(t *testing.T) {
	t.Run("input_example.txt", func(t *testing.T) {
		input, _ := GetInput("input_example.txt")
		want := 7036
		got := Part1(input)

		utils.CheckEqual(got, want, t)
	})

	t.Run("input_example2.txt", func(t *testing.T) {
		input, _ := GetInput("input_example2.txt")
		want := 11048
		got := Part1(input)

		utils.CheckEqual(got, want, t)
	})
}
