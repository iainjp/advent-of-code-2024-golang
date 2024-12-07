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

func TestCanBeSolved(t *testing.T) {

	t.Run("can be solved", func(t *testing.T) {
		eq := Equation{targetTotal: 190, operands: []int{19, 10}}
		want := true

		got := CanBeSolved(eq)

		utils.CheckEqual(got, want, t)
	})

	t.Run("can't be solved", func(t *testing.T) {
		eq := Equation{targetTotal: 21037, operands: []int{9, 7, 18, 13}}
		want := false

		got := CanBeSolved(eq)

		utils.CheckEqual(got, want, t)
	})
}
