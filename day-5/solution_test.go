package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	fileName := "input_minimal_test.txt"

	wantedRules := []OrderingRule{
		{47, 53},
		{97, 13},
		{97, 61},
	}

	wantedUpdates := []Update{
		{pages: []int{75, 47, 61, 53, 29}},
		{pages: []int{97, 61, 53, 29, 13}},
		{pages: []int{75, 29, 13}},
	}

	want := &Input{
		rules:   wantedRules,
		updates: wantedUpdates,
	}

	got, _ := GetInput(fileName)

	utils.CheckEqual(got, want, t)
}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_test.txt")

	want := 143
	got := Part1(input)

	utils.CheckEqual(got, want, t)
}

func TestPart2(t *testing.T) {
	input, _ := GetInput("input_test.txt")

	want := 123
	got := Part2(input)

	utils.CheckEqual(got, want, t)
}
