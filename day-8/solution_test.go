package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	got, _ := GetInput("input_minimal.txt")

	want := &Input{
		pointMap: PointMap{
			m: map[Coord]Point{
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
		},
	}

	utils.CheckEqual(got, want, t)
}

func TestPointMap(t *testing.T) {
	pointMap := PointMap{
		m: map[Coord]Point{
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
	t.Run("GetAntennas returns antenna coords", func(t *testing.T) {
		want := []Coord{
			{1, 0},
			{0, 1},
			{2, 1},
			{1, 2},
		}

		got := pointMap.GetAntennas()

		for _, coord := range want {
			utils.CheckContains(got, coord, t)
		}
	})
}

func TestPoint(t *testing.T) {
	antenna := Point{symbol: "A"}
	notAntenna := Point{symbol: "."}

	t.Run("isAntenna returns true for antenna", func(t *testing.T) {
		want := true
		got := antenna.IsAntenna()

		utils.CheckEqual(got, want, t)
	})

	t.Run("isAntenna returns false for non-antenna", func(t *testing.T) {
		want := false
		got := notAntenna.IsAntenna()

		utils.CheckEqual(got, want, t)
	})

	t.Run("SameFrequency returns false for not same frequency", func(t *testing.T) {
		otherAntenna := Point{symbol: "a"}
		want := false

		got := antenna.SameFrequency(&otherAntenna)

		utils.CheckEqual(got, want, t)
	})

	t.Run("SameFrequency returns true for same frequency", func(t *testing.T) {
		otherAntenna := Point{symbol: "A"}
		want := true

		got := antenna.SameFrequency(&otherAntenna)

		utils.CheckEqual(got, want, t)
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

		want := []Coord{
			{2, 2},
			{8, 8},
		}

		got := GetAntinodes(input)

		for _, want := range want {
			utils.CheckContains(got, want, t)
		}

	})

	t.Run("other direction pair", func(t *testing.T) {
		input := Pair[Coord]{
			Coord{4, 6},
			Coord{6, 4},
		}

		want := []Coord{
			{2, 8},
			{8, 2},
		}

		got := GetAntinodes(input)

		for _, want := range want {
			utils.CheckContains(got, want, t)
		}
	})
}

func TestGetAntinodesWithResonantHarmonics(t *testing.T) {

}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_test.txt")

	want := 14
	got := Part1(input)

	utils.CheckEqual(got, want, t)
}

func TestPart2(t *testing.T) {
	input, _ := GetInput("input_test.txt")

	want := 34
	got := Part2(input)

	utils.CheckEqual(got, want, t)
}
