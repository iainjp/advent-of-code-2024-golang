package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	got, _ := GetInput("input_minimal.txt")

	want := &Input{
		pointMap: &PointMap{
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
