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

	t.Run("Blink()", func(t *testing.T) {
		stoneLine := BuildStones([]int{0, 1, 10, 99, 999})
		want := []int{1, 2024, 1, 0, 9, 9, 2021976}

		stoneLine.Blink()

		got := stoneLine.GetNumbers()

		utils.CheckSlicesHaveSameElements(got, want, t)
	})

	t.Run("BlinkTimes(6)", func(t *testing.T) {
		stoneLine := BuildStones([]int{125, 17})
		want := []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}

		stoneLine.BlinkTimes(6)

		got := stoneLine.GetNumbers()

		utils.CheckSlicesHaveSameElements(got, want, t)
	})

	type testInput struct {
		name   string
		ints   []int
		blinks int
		want   int
	}

	inputs := []testInput{
		{"SimulateBlinkTimes(1) == 3", []int{125, 17}, 1, 3},
		{"SimulateBlinkTimes(2) == 4", []int{125, 17}, 2, 4},
		{"SimulateBlinkTimes(3) == 5", []int{125, 17}, 3, 5},
		{"SimulateBlinkTimes(4) == 9", []int{125, 17}, 4, 9},
		{"SimulateBlinkTimes(5) == 13", []int{125, 17}, 5, 13},
		{"SimulateBlinkTimes(6) == 22", []int{125, 17}, 6, 22},
		{"SimulateBlinkTimes(7) == 31", []int{125, 17}, 7, 31},
		{"SimulateBlinkTimes(8) == 42", []int{125, 17}, 8, 42},
		{"SimulateBlinkTimes(25) == 55312", []int{125, 17}, 25, 55312},
	}

	t.Run("SimulateBlinkTimes(4) == 9", func(t *testing.T) {
		stoneLine := BuildStones([]int{125, 17})
		got := stoneLine.SimulateBlinkTimes(4)

		utils.CheckEqual(got, 9, t)
	})

	for _, tt := range inputs {
		t.Run(tt.name, func(t *testing.T) {
			stoneLine := BuildStones([]int{125, 17})
			got := stoneLine.SimulateBlinkTimes(tt.blinks)

			utils.CheckEqual(got, tt.want, t)
		})
	}
}

func TestPart1(t *testing.T) {
	want := 55312

	input, _ := GetInput("input_test.txt")
	got := Part1(input)

	utils.CheckEqual(got, want, t)
}
