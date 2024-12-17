package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	got, _ := GetInput("input_example.txt")

	utils.CheckEqual(len(got.grid.cgMap), 100, t)
	utils.CheckEqual(len(*got.moves), 700, t)
	utils.CheckEqual(got.grid.maxX, 10, t)
	utils.CheckEqual(got.grid.maxY, 10, t)
}

func TestInput(t *testing.T) {
	t.Run("Run() doesn't panic", func(t *testing.T) {
		got, _ := GetInput("input_example.txt")
		got.Run()
	})

	t.Run("PopMove() gets next move", func(t *testing.T) {
		input := Input{
			moves: &[]string{"^", "v", ">"},
		}

		want := []string{"^", "v", ">", ""}

		got := []string{
			input.PopMove(),
			input.PopMove(),
			input.PopMove(),
			input.PopMove(),
		}

		utils.CheckEqual(got[0], want[0], t)
		utils.CheckEqual(got[1], want[1], t)
		utils.CheckEqual(got[2], want[2], t)
		utils.CheckEqual(got[3], want[3], t)
	})
}

func TestCoord(t *testing.T) {
	coord := Coord{x: 4, y: 6}

	t.Run("GPS()", func(t *testing.T) {
		want := 604
		got := coord.GPS()

		utils.CheckEqual(got, want, t)
	})
}

func TestMoves(t *testing.T) {
	t.Run("Left() moves robot and box left", func(t *testing.T) {
		input, _ := GetInput("input_left.txt")

		want := []string{"#", ".", "O", "O", "@", ".", ".", "#"}

		input.grid.Left()
		got := input.grid.Row(1)

		utils.CheckEqual(got, want, t)
	})

	t.Run("Left() moves robot and multiple boxes left", func(t *testing.T) {
		input, _ := GetInput("input_left.txt")

		want := []string{"#", "O", "O", "@", ".", ".", ".", "#"}

		input.grid.Left()
		input.grid.Left()
		got := input.grid.Row(1)

		utils.CheckEqual(got, want, t)
	})

	t.Run("Left(), no space, no change", func(t *testing.T) {
		input, _ := GetInput("input_left.txt")

		want := []string{"#", "O", "O", "@", ".", ".", ".", "#"}

		input.grid.Left()
		input.grid.Left()
		input.grid.Left()
		got := input.grid.Row(1)

		utils.CheckEqual(got, want, t)
	})

	t.Run("Right() moves robot and box right", func(t *testing.T) {
		input, _ := GetInput("input_right.txt")

		want := []string{"#", ".", ".", "@", "O", "O", ".", "#"}

		input.grid.Right()
		got := input.grid.Row(1)

		utils.CheckEqual(got, want, t)
	})

	t.Run("Right() moves robot and multiple boxes right", func(t *testing.T) {
		input, _ := GetInput("input_right.txt")

		want := []string{"#", ".", ".", ".", "@", "O", "O", "#"}

		input.grid.Right()
		input.grid.Right()
		got := input.grid.Row(1)

		utils.CheckEqual(got, want, t)
	})

	t.Run("Right(), no space, no change", func(t *testing.T) {
		input, _ := GetInput("input_right.txt")

		want := []string{"#", ".", ".", ".", "@", "O", "O", "#"}

		input.grid.Right()
		input.grid.Right()
		input.grid.Right()
		got := input.grid.Row(1)

		utils.CheckEqual(got, want, t)
	})
}
