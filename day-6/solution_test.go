package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestMakeGuard(t *testing.T) {
	point := Point{1, 1}
	type test struct {
		input string
		want  Direction
	}

	tests := []test{
		{input: "^", want: Up},
		{input: ">", want: Right},
		{input: "v", want: Down},
		{input: "<", want: Left},
	}

	for _, tc := range tests {
		t.Run("creates guard in right direction", func(t *testing.T) {
			got := MakeGuard(tc.input, point)

			utils.CheckNotNil(got, t)
			utils.CheckEqual(got.direction, tc.want, t)
			utils.CheckEqual(got.position, point, t)
		})
	}

	t.Run("returns nil for non-guard position", func(t *testing.T) {
		var want *Guard

		got := MakeGuard(".", point)
		utils.CheckEqual(got, want, t)
	})
}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_test.txt")
	want := 41

	got := Part1(input)

	utils.CheckEqual(got, want, t)
}

func TestPart2(t *testing.T) {
	input, _ := GetInput("input_test.txt")
	want := 6

	got := Part2(input)

	utils.CheckEqual(got, want, t)
}

func TestAllMapOptions(t *testing.T) {
	m := Map{
		{0, 0}: ".",
		{0, 1}: ".",
		{0, 2}: ".",
		{1, 0}: "^",
		{1, 1}: ".",
		{1, 2}: "#",
		{2, 0}: ".",
		{2, 1}: ".",
		{2, 2}: ".",
	}

	guard := Guard{
		position:  Point{0, 2},
		direction: Up,
		path: []Point{
			{0, 1},
			{0, 2},
			{2, 1},
			{1, 0},
			{1, 2},
		},
		seen: make(map[Point]int),
	}

	want := 3
	got := AllMapOptions(m, guard)

	utils.CheckEqual(len(got), want, t)
}
