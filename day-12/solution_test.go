package main

import (
	"maps"
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	input, _ := GetInput("input_minimal.txt")

	t.Run("Get A's", func(t *testing.T) {
		want := []Coord{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		}
		got := input.plotMap.GetCoords("A")
		utils.CheckSlicesHaveSameElements(got, want, t)
	})

	t.Run("Get B's", func(t *testing.T) {
		want := []Coord{
			{0, 1},
			{0, 2},
			{1, 1},
			{1, 2},
		}
		got := input.plotMap.GetCoords("B")
		utils.CheckSlicesHaveSameElements(got, want, t)
	})
}

func TestPlotMap(t *testing.T) {

	t.Run("GetPlots()", func(t *testing.T) {
		input, _ := GetInput("input_minimal.txt")

		want1 := input.plotMap.Get(Coord{0, 0})
		want2 := input.plotMap.Get(Coord{1, 0})
		want3 := input.plotMap.Get(Coord{2, 0})
		want4 := input.plotMap.Get(Coord{3, 0})

		want := []*Plot{
			want1, want2, want3, want4,
		}
		got := input.plotMap.GetPlots("A")

		utils.CheckEqual(len(got), 4, t)
		utils.CheckSlicesHaveSameElements(got, want, t)
	})
}

func TestGetRegions(t *testing.T) {
	input, _ := GetInput("input_minimal.txt")

	want := 5
	got := GetRegions(input.plotMap)

	utils.CheckEqual(len(got), want, t)
}

func TestRegion(t *testing.T) {
	input, _ := GetInput("input_minimal.txt")

	got := GetRegions(input.plotMap)

	crop := "A"
	filter := func(r Region) bool {
		// this feels dirty
		for k := range maps.Keys(r.plots.data) {
			return k.crop == crop
		}
		return false
	}
	regionA := utils.Filter(got, filter)[0]
	t.Run("Area()", func(t *testing.T) {
		utils.CheckEqual(len(regionA.plots.data), 4, t)
		utils.CheckEqual(regionA.Area(), 4, t)
	})

	t.Run("Perimeter()", func(t *testing.T) {
		utils.CheckEqual(regionA.Perimeter(), 10, t)
	})

	t.Run("Sides()", func(t *testing.T) {
		utils.CheckEqual(regionA.Sides(), 4, t)
	})

}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_minimal.txt")

	want := 140
	got := Part1(input)

	utils.CheckEqual(got, want, t)
}

func TestPart2(t *testing.T) {
	t.Run("input_minimal.txt", func(t *testing.T) {

		input, _ := GetInput("input_minimal.txt")

		want := 80
		got := Part2(input)

		utils.CheckEqual(got, want, t)
	})

	t.Run("input_example.txt", func(t *testing.T) {
		input, _ := GetInput("input_example.txt")

		want := 368
		got := Part2(input)

		utils.CheckEqual(got, want, t)
	})
}
