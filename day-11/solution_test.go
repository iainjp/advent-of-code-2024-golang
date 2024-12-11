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

	t.Run("len(BlinkTimes(1)) == SimulateBlinkTimes(1)", func(t *testing.T) {
		stoneLine := BuildStones([]int{125, 17})
		want := 3
		got := stoneLine.SimulateBlinkTimes(1)

		utils.CheckEqual(got, want, t)
	})

	t.Run("len(BlinkTimes(2)) == SimulateBlinkTimes(2)", func(t *testing.T) {
		stoneLine := BuildStones([]int{125, 17})
		want := 4
		got := stoneLine.SimulateBlinkTimes(2)

		utils.CheckEqual(got, want, t)
	})

	t.Run("len(BlinkTimes(6)) == SimulateBlinkTimes(6)", func(t *testing.T) {
		stoneLine := BuildStones([]int{125, 17})
		want := 22
		got := stoneLine.SimulateBlinkTimes(6)

		utils.CheckEqual(got, want, t)
	})

	t.Run("len(BlinkTimes(25)) == SimulateBlinkTimes(25)", func(t *testing.T) {
		stoneLine := BuildStones([]int{4, 4841539, 66, 5279, 49207, 134, 609568, 0})
		want := 212655
		got := stoneLine.SimulateBlinkTimes(25)

		utils.CheckEqual(got, want, t)
	})
}

func TestPart1(t *testing.T) {
	want := 55312

	input, _ := GetInput("input_test.txt")
	got := Part1(input)

	utils.CheckEqual(got, want, t)
}
