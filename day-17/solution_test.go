package main

import (
	"testing"

	"iain.fyi/aoc2024/structure"
	"iain.fyi/aoc2024/utils"
)

func TestParseRegister(t *testing.T) {
	input := "Register A: 729"

	want := 729
	got := ParseRegisterValue(input)

	utils.CheckEqual(got, want, t)
}

func TestParseProgram(t *testing.T) {
	input := "Program: 0,1,5,4,3,0"

	want := []Operation{
		{0, 1},
		{5, 4},
		{3, 0},
	}
	got := ParseProgram(input)

	utils.CheckEqual(got, want, t)
}

func TestGetInput(t *testing.T) {
	want := Input{
		debugger: &Debugger{
			State: State{
				A:                729,
				B:                0,
				C:                0,
				Output:           structure.NewList[int](),
				InstructionIndex: 0,
			},
			Program: []Operation{
				{0, 1},
				{5, 4},
				{3, 0},
			},
		},
	}
	got, _ := GetInput("input_example.txt")

	utils.CheckEqual(*got, want, t)
}

func TestOperation(t *testing.T) {

	var newState = func() State {
		return State{
			A:                0,
			B:                0,
			C:                0,
			Output:           structure.List[int]{},
			InstructionIndex: 0,
		}
	}

	t.Run("adv()", func(t *testing.T) {
		state := newState()
		state.A = 80
		op := Operation{OpCode: 0, Operand: 3}

		want := newState()
		want.A = 10
		want.InstructionIndex = 1

		op.adv(&state)

		utils.CheckEqual(state, want, t)
	})

	t.Run("bxl()", func(t *testing.T) {
		state := newState()
		state.B = 2
		op := Operation{OpCode: 1, Operand: 4}

		want := newState()
		want.B = 6
		want.InstructionIndex = 1

		op.bxl(&state)

		utils.CheckEqual(state, want, t)
	})

	t.Run("bst()", func(t *testing.T) {
		state := newState()
		op := Operation{OpCode: 2, Operand: 19}

		want := newState()
		want.B = 3
		want.InstructionIndex = 1

		op.bst(&state)

		utils.CheckEqual(state, want, t)
	})

	t.Run("jnz() do nothing on A==0", func(t *testing.T) {
		state := newState()
		state.A = 0
		op := Operation{OpCode: 3, Operand: 10}

		want := newState()
		want.InstructionIndex = 1

		op.jnz(&state)

		utils.CheckEqual(state, want, t)
	})

	t.Run("jnz() jump on on A!=0", func(t *testing.T) {
		state := newState()
		state.A = 3
		op := Operation{OpCode: 3, Operand: 4}

		want := newState()
		want.A = 3
		want.InstructionIndex = 2

		op.jnz(&state)

		utils.CheckEqual(state, want, t)
	})

	t.Run("bxc()", func(t *testing.T) {
		state := newState()
		state.B = 3
		state.C = 4
		op := Operation{OpCode: 4, Operand: 4}

		want := newState()
		want.B = 7
		want.C = 4
		want.InstructionIndex = 1

		op.bxc(&state)

		utils.CheckEqual(state, want, t)
	})

	t.Run("out()", func(t *testing.T) {
		state := newState()
		op := Operation{OpCode: 5, Operand: 12}

		want := newState()
		wantOutput := structure.NewList[int]()
		wantOutput.Add(4)
		want.Output = wantOutput
		want.InstructionIndex = 1

		op.out(&state)

		utils.CheckEqual(state, want, t)
	})

	t.Run("bdv()", func(t *testing.T) {
		state := newState()
		state.A = 80
		op := Operation{OpCode: 0, Operand: 3}

		want := newState()
		want.A = 80
		want.B = 10
		want.InstructionIndex = 1

		op.bdv(&state)

		utils.CheckEqual(state, want, t)
	})

	t.Run("cdv()", func(t *testing.T) {
		state := newState()
		state.A = 80
		op := Operation{OpCode: 0, Operand: 3}

		want := newState()
		want.A = 80
		want.C = 10
		want.InstructionIndex = 1

		op.cdv(&state)

		utils.CheckEqual(state, want, t)
	})
}

func TestPart1(t *testing.T) {
	input, _ := GetInput("input_example.txt")

	got := Part1(input)
	want := 4635635210

	utils.CheckEqual(got, want, t)
}
