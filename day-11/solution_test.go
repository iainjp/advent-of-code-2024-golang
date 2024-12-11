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

func TestBuildStones(t *testing.T) {
	input := []int{1, 2, 3}
	want := []int{1, 2, 3}

	got := BuildStones(input)
	gotNumbers := got.GetNumbers()

	utils.CheckSlicesHaveSameElements(gotNumbers, want, t)
}

func TestStoneLine(t *testing.T) {

	t.Run("GetNumbers() returns numbers", func(t *testing.T) {
		stone1 := Stone{number: 1}
		stone2 := Stone{number: 2}
		stone1.next = &stone2
		stone2.prev = &stone1
		stone3 := Stone{number: 3}
		stone2.next = &stone3
		stone3.prev = &stone2

		stoneLine := StoneLine{head: &stone1}

		want := []int{1, 2, 3}
		got := stoneLine.GetNumbers()

		utils.CheckSlicesHaveSameElements(got, want, t)
	})

	t.Run("ToSlice() returns slice", func(t *testing.T) {
		stone1 := Stone{number: 1}
		stone2 := Stone{number: 2}
		stone1.next = &stone2
		stone2.prev = &stone1
		stone3 := Stone{number: 3}
		stone2.next = &stone3
		stone3.prev = &stone2

		stoneLine := StoneLine{head: &stone1}

		want := []Stone{stone1, stone2, stone3}
		got := stoneLine.ToSlice()

		utils.CheckSlicesHaveSameElements(got, want, t)
	})

}
