package main

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestGetInput(t *testing.T) {
	want := Input{
		equations: []Equation{
			{targetTotal: 190, operands: []uint64{10, 19}},
			{targetTotal: 3267, operands: []uint64{81, 40, 27}},
			{targetTotal: 83, operands: []uint64{17, 5}},
		},
	}

	got, _ := GetInput("input_minimal.txt")

	utils.CheckEqual(got, &want, t)
}

func TestCanBeSolved(t *testing.T) {

	t.Run("can be solved", func(t *testing.T) {
		eq := Equation{targetTotal: 190, operands: []uint64{19, 10}}
		want := true

		got := CanBeSolved(eq)

		utils.CheckEqual(got, want, t)
	})

	t.Run("can't be solved", func(t *testing.T) {
		eq := Equation{targetTotal: 21037, operands: []uint64{9, 7, 18, 13}}
		want := false

		got := CanBeSolved(eq)

		utils.CheckEqual(got, want, t)
	})
}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_test.txt")
	want := uint64(3749)

	got := Part1(input)

	utils.CheckEqual(got, want, t)
}

func TestTree(t *testing.T) {
	t.Run("first insert adds root", func(t *testing.T) {
		tree := Tree{}
		tree.Insert(64)

		got := tree.root
		want := &Node{value: uint64(64)}

		utils.CheckEqual(got, want, t)
	})

	t.Run("second insert adds and mults", func(t *testing.T) {
		tree := Tree{root: &Node{value: uint64(64)}}
		tree.Insert(2)

		gotMult := tree.root.mult
		gotAdd := tree.root.add

		wantMult := &Node{value: uint64(128)}
		wantAdd := &Node{value: uint64(66)}

		utils.CheckEqual(gotMult, wantMult, t)
		utils.CheckEqual(gotAdd, wantAdd, t)
	})

	t.Run("third insert adds and mults _again_", func(t *testing.T) {
		tree := Tree{root: &Node{value: uint64(64)}}
		tree.Insert(2)
		tree.Insert(3)

		gotMult1 := tree.root.mult.mult
		gotMult2 := tree.root.add.mult
		gotAdd1 := tree.root.mult.add
		gotAdd2 := tree.root.add.add

		// 64 x 2 x 3
		wantMult1 := &Node{value: uint64(384)}
		// 64 + 2 x 3
		wantMult2 := &Node{value: uint64(198)}
		// 64 x 2 + 3
		wantAdd1 := &Node{value: uint64(131)}
		// 64 + 2 + 3
		wantAdd2 := &Node{value: uint64(69)}

		utils.CheckEqual(gotMult1, wantMult1, t)
		utils.CheckEqual(gotMult2, wantMult2, t)
		utils.CheckEqual(gotAdd1, wantAdd1, t)
		utils.CheckEqual(gotAdd2, wantAdd2, t)
	})

	t.Run("get leaf nodes", func(t *testing.T) {
		tree := Tree{}
		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)

		wantedValues := []uint64{24, 10, 20, 9}
		var gotValues []uint64
		leafNodes := tree.GetLeafNodes()

		for _, n := range leafNodes {
			gotValues = append(gotValues, uint64(n.value))
		}

		utils.CheckEqual(gotValues, wantedValues, t)
	})
}
