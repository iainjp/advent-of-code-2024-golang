package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	want := &Input{stones: []int{125, 17}}
	got, _ := GetInput("input_test.txt")

	utils.CheckEqual(got, want, t)
}
