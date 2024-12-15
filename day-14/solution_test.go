package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	want0 := Robot{
		position: &Position{
			x: 0, y: 4,
		},
		velocity: &Velocity{
			x: 3, y: -3,
		},
	}
	want2 := Robot{
		position: &Position{
			x: 10, y: 3,
		},
		velocity: &Velocity{
			x: -1, y: 2,
		},
	}
	want11 := Robot{
		position: &Position{
			x: 9, y: 5,
		},
		velocity: &Velocity{
			x: -3, y: -3,
		},
	}

	got, _ := GetInput("input_example.txt")

	utils.CheckEqual(*got.robots[0], want0, t)
	utils.CheckEqual(*got.robots[2], want2, t)
	utils.CheckEqual(*got.robots[11], want11, t)
}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_example.txt")
	Part1(input)
}
