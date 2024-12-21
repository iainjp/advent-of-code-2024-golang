package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestParseRegister(t *testing.T) {
	input := "Register A: 729"

	want := 729
	got := ParseRegisterValue(input)

	utils.CheckEqual(got, want, t)
}

func TestParseProgram(t *testing.T) {
	input := "Program: 0,1,5,4,3,0"

	want := []Operation{
		{0, 1},
		{5, 4},
		{3, 0},
	}
	got := ParseProgram(input)

	utils.CheckEqual(got, want, t)
}

func TestGetInput(t *testing.T) {
	want := Input{
		debugger: &Debugger{
			A: 729,
			B: 0,
			C: 0,
			Program: []Operation{
				{0, 1},
				{5, 4},
				{3, 0},
			},
		},
	}
	got, _ := GetInput("input_example.txt")

	utils.CheckEqual(*got, want, t)
}
