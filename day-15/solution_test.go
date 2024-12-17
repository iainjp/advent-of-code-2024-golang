package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	got, _ := GetInput("input_example.txt")

	utils.CheckEqual(len(got.grid), 100, t)
	utils.CheckEqual(len(*got.moves), 700, t)
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
