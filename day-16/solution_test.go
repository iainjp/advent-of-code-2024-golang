package main

import (
	"math"
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
	input, _ := GetInput("input_example.txt")

	want := 7036
	got := Part1(input)

	utils.CheckEqual(got-math.MaxInt32, want, t)

}
