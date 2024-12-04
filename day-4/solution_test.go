package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	testInputFilename := "input_test.txt"

	wantedGrid := map[Coord]string{
		Coord{0, 0}: "M",
		Coord{1, 0}: "M",
		Coord{2, 0}: "M",
		Coord{3, 0}: "S",
		Coord{0, 1}: "M",
		Coord{1, 1}: "S",
		Coord{2, 1}: "A",
		Coord{3, 1}: "M",
		Coord{0, 2}: "A",
		Coord{1, 2}: "M",
		Coord{2, 2}: "X",
		Coord{3, 2}: "S",
	}

	want := &Input{
		maxX: 3,
		maxY: 2,
		grid: wantedGrid,
	}
	got, _ := GetInput(testInputFilename)

	utils.CheckEqual(got, want, t)
}
