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

func TestRobot(t *testing.T) {
	makeRobot := func() Robot {
		return Robot{
			position: &Position{
				x: 1, y: 1,
			},
			velocity: &Velocity{
				x: -2, y: -1,
			},
		}
	}

	t.Run("Tick() once", func(t *testing.T) {
		want := Position{
			x: 4, y: 0,
		}
		robot := makeRobot()
		robot.Tick(5, 5)
		got := *robot.position

		utils.CheckEqual(got, want, t)
	})

	t.Run("Tick() twice", func(t *testing.T) {
		want := Position{
			x: 2, y: 4,
		}
		robot := makeRobot()
		robot.Tick(5, 5)
		robot.Tick(5, 5)
		got := *robot.position

		utils.CheckEqual(got, want, t)
	})

}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_example.txt")
	input.height = 7
	input.width = 11

	got := Part1(input)

	utils.CheckEqual(got, 12, t)
}
