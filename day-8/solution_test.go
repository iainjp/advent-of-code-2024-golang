package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	got, _ := GetInput("input_minimal.txt")

	want := &Input{
		pointMap: PointMap{
			{0, 0}: {symbol: "."},
			{1, 0}: {symbol: "0"},
			{2, 0}: {symbol: "."},
			{0, 1}: {symbol: "A"},
			{1, 1}: {symbol: "."},
			{2, 1}: {symbol: "A"},
			{0, 2}: {symbol: "."},
			{1, 2}: {symbol: "b"},
			{2, 2}: {symbol: "."},
		},
	}

	utils.CheckEqual(got, want, t)
}

func TestPointMap(t *testing.T) {
	pointMap := PointMap{
		{0, 0}: {symbol: "."},
		{1, 0}: {symbol: "0"},
		{2, 0}: {symbol: "."},
		{0, 1}: {symbol: "A"},
		{1, 1}: {symbol: "."},
		{2, 1}: {symbol: "A"},
		{0, 2}: {symbol: "."},
		{1, 2}: {symbol: "b"},
		{2, 2}: {symbol: "."},
	}
	t.Run("GetAntennas returns antenna coords", func(t *testing.T) {
		want := []Coord{
			{1, 0},
			{0, 1},
			{2, 1},
			{1, 2},
		}

		got := GetAntennas(pointMap)

		for _, coord := range want {
			utils.CheckContains(got, coord, t)
		}
	})
}

func TestGetAllUniquePairs(t *testing.T) {
	input := []Coord{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}

	want := []Pair[Coord]{
		{Coord{0, 0}, Coord{0, 1}},
		{Coord{0, 0}, Coord{1, 0}},
		{Coord{0, 0}, Coord{1, 1}},
		{Coord{0, 1}, Coord{0, 0}},
		{Coord{0, 1}, Coord{1, 0}},
		{Coord{0, 1}, Coord{1, 1}},
		{Coord{1, 0}, Coord{0, 0}},
		{Coord{1, 0}, Coord{0, 1}},
		{Coord{1, 0}, Coord{1, 1}},
		{Coord{1, 1}, Coord{0, 0}},
		{Coord{1, 1}, Coord{0, 1}},
		{Coord{1, 1}, Coord{1, 0}},
	}

	got := GetAllUniquePairs(input)

	utils.CheckEqual(len(got), len(want), t)
	for _, pair := range want {
		utils.CheckContains(got, pair, t)
	}
}

func TestGetAntinodes(t *testing.T) {
	t.Run("simple pair", func(t *testing.T) {
		input := Pair[Coord]{
			Coord{4, 4},
			Coord{6, 6},
		}

		want := Pair[Coord]{
			Coord{2, 2},
			Coord{8, 8},
		}

		got := GetAntinodes(input)

		utils.CheckEqual(got, want, t)
	})

	t.Run("other direction pair", func(t *testing.T) {
		input := Pair[Coord]{
			Coord{4, 6},
			Coord{6, 4},
		}

		want := Pair[Coord]{
			Coord{2, 8},
			Coord{8, 2},
		}

		got := GetAntinodes(input)

		utils.CheckEqual(got, want, t)
	})
}
