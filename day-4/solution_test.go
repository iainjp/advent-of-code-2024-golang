package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	testInputFilename := "input_test.txt"

	wantedGrid := GridMap{
		{0, 0}: "M",
		{1, 0}: "M",
		{2, 0}: "M",
		{3, 0}: "S",
		{0, 1}: "M",
		{1, 1}: "S",
		{2, 1}: "A",
		{3, 1}: "M",
		{0, 2}: "A",
		{1, 2}: "M",
		{2, 2}: "X",
		{3, 2}: "S",
	}

	want := &Input{
		maxX: 3,
		maxY: 2,
		grid: wantedGrid,
	}
	got, _ := GetInput(testInputFilename)

	utils.CheckEqual(got, want, t)
}

func TestGetCoordsOfLetter(t *testing.T) {
	grid := GridMap{
		{0, 0}: "M",
		{1, 0}: "M",
		{2, 0}: "M",
		{3, 0}: "S",
		{0, 1}: "M",
		{1, 1}: "S",
		{2, 1}: "A",
		{3, 1}: "M",
		{0, 2}: "A",
		{1, 2}: "M",
		{2, 2}: "X",
		{3, 2}: "S",
	}

	want := []Point{
		{2, 1},
		{0, 2},
	}

	got := GetCoordsOfLetter(grid, "A")

	utils.CheckEqual(got, want, t)
}

func TestGetOrderedKeys(t *testing.T) {
	grid := GridMap{
		{0, 0}: "M",
		{1, 0}: "M",
		{2, 0}: "M",
		{3, 0}: "S",
		{0, 1}: "M",
		{1, 1}: "S",
		{2, 1}: "A",
		{3, 1}: "M",
		{0, 2}: "A",
		{1, 2}: "M",
		{2, 2}: "X",
		{3, 2}: "S",
	}

	want := []Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{3, 1},
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
	}

	got := GetOrderedKeys(grid)

	utils.CheckEqual(got, want, t)
}
