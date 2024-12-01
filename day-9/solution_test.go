package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	wantedSeries := []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}
	want := &Input{diskmap: Diskmap{series: wantedSeries}}

	got, _ := GetInput("input_test.txt")

	utils.CheckEqual(got, want, t)
}

func TestDiskMap(t *testing.T) {
	blocks := BlocksOnDisk{
		blocks: []Block{
			FileBlock{id: 0},
			Space{},
			Space{},
			FileBlock{id: 1},
			FileBlock{id: 1},
			FileBlock{id: 1},
			Space{},
			Space{},
			Space{},
			Space{},
			FileBlock{id: 2},
			FileBlock{id: 2},
			FileBlock{id: 2},
			FileBlock{id: 2},
			FileBlock{id: 2},
		},
	}
	t.Run(".ToBlocks() returns blocks", func(t *testing.T) {
		diskMap := Diskmap{
			series: []int{1, 2, 3, 4, 5},
		}

		got := diskMap.ToBlocks()
		want := blocks

		utils.CheckEqual(len(got.blocks), len(want.blocks), t)
		for _, w := range want.blocks {
			utils.CheckContains(got.blocks, w, t)
		}
	})

	t.Run("Print() prints block symbols", func(t *testing.T) {
		want := "0..111....22222"
		got := blocks.Print()

		utils.CheckEqual(got, want, t)
	})

	t.Run("ToBlocks() then Print() more complex", func(t *testing.T) {
		diskmap := Diskmap{
			series: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
		}

		want := "00...111...2...333.44.5555.6666.777.888899"
		got := diskmap.ToBlocks().Print()

		utils.CheckEqual(got, want, t)
	})

	t.Run("Compact()", func(t *testing.T) {
		want := BlocksOnDisk{
			blocks: []Block{
				FileBlock{id: 0},
				FileBlock{id: 2},
				FileBlock{id: 2},
				FileBlock{id: 1},
				FileBlock{id: 1},
				FileBlock{id: 1},
				FileBlock{id: 2},
				FileBlock{id: 2},
				FileBlock{id: 2},
				Space{},
				Space{},
				Space{},
				Space{},
				Space{},
				Space{},
			},
		}

		got := blocks.Compact()

		utils.CheckEqual(len(got.blocks), len(want.blocks), t)
		for _, w := range want.blocks {
			utils.CheckContains(got.blocks, w, t)
		}
	})

	t.Run("Compact() and Print() more complex", func(t *testing.T) {
		diskmap := Diskmap{
			series: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
		}

		want := "0099811188827773336446555566.............."
		got := diskmap.ToBlocks().Compact().Print()

		utils.CheckEqual(got, want, t)
	})

	t.Run("BuildIndex() returns ranges by ID", func(t *testing.T) {
		got := blocks.BuildIndex()

		first := blocks.blocks[got[0].start:got[0].end]
		utils.CheckEqual(len(first), 1, t)
		utils.CheckEqual(first[0], blocks.blocks[0], t)

		second := blocks.blocks[got[1].start:got[1].end]
		utils.CheckEqual(len(second), 3, t)
		utils.CheckEqual(second[0], blocks.blocks[3], t)
		utils.CheckEqual(second[1], blocks.blocks[4], t)
		utils.CheckEqual(second[2], blocks.blocks[5], t)

		third := blocks.blocks[got[2].start:got[2].end]
		utils.CheckEqual(len(third), 5, t)
		utils.CheckEqual(third[0], blocks.blocks[10], t)
		utils.CheckEqual(third[1], blocks.blocks[11], t)
		utils.CheckEqual(third[2], blocks.blocks[12], t)
		utils.CheckEqual(third[3], blocks.blocks[13], t)
		utils.CheckEqual(third[4], blocks.blocks[14], t)
	})

	t.Run("OrderedSpaceSpans() returns space spans left-to-right", func(t *testing.T) {
		got := blocks.OrderedSpaceSpans()

		first := blocks.blocks[got[0].start:got[0].end]
		utils.CheckEqual(len(first), 2, t)
		utils.CheckEqual(first[0], blocks.blocks[1], t)
		utils.CheckEqual(first[1], blocks.blocks[2], t)

		second := blocks.blocks[got[1].start:got[1].end]
		utils.CheckEqual(len(second), 4, t)
		utils.CheckEqual(second[0], blocks.blocks[6], t)
		utils.CheckEqual(second[1], blocks.blocks[7], t)
		utils.CheckEqual(second[2], blocks.blocks[8], t)
		utils.CheckEqual(second[3], blocks.blocks[9], t)
	})

	t.Run("CompactContiguousFiles() and Print(), complex", func(t *testing.T) {
		diskmap := Diskmap{
			series: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
		}

		want := "00992111777.44.333....5555.6666.....8888.."
		got := diskmap.ToBlocks().CompactContiguousFiles().Print()

		utils.CheckEqual(got, want, t)
	})

	t.Run("UniqueFileIds()", func(t *testing.T) {
		want := []int{0, 1, 2}
		got := blocks.UniqueFileIds()

		utils.CheckEqual(len(got), len(want), t)
		for _, wantInt := range want {
			utils.CheckContains(got, wantInt, t)
		}
	})
}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_test.txt")

	want := 1928
	got := Part1(input)

	utils.CheckEqual(got, want, t)
}
