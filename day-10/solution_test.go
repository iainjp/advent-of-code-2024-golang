package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	got, _ := GetInput("input_test.txt")

	utils.CheckEqual(len(got.graph.trailheads), 9, t)
}
