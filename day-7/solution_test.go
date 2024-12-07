package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	want := Input{
		equations: []Equation{
			{targetTotal: 190, operands: []int{10, 19}},
			{targetTotal: 3267, operands: []int{81, 40, 27}},
			{targetTotal: 83, operands: []int{17, 5}},
		},
	}

	got, _ := GetInput("input_minimal.txt")

	utils.CheckEqual(got, &want, t)
}
