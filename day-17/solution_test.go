package main

import (
	"fmt"
	"testing"

	"iain.fyi/aoc2024/structure"
	"iain.fyi/aoc2024/utils"
)

type StateBuilder struct {
	A, B, C          int
	Output           structure.List[int]
	InstructionIndex int
}

func NewStateBuilder() *StateBuilder {

	sb := StateBuilder{
		A:                0,
		B:                0,
		C:                0,
		Output:           structure.NewList[int](),
		InstructionIndex: 0,
	}
	return &sb
}

func (sb *StateBuilder) SetA(a int) *StateBuilder {
	sb.A = a
	return sb
}

func (sb *StateBuilder) SetB(b int) *StateBuilder {
	sb.B = b
	return sb
}

func (sb *StateBuilder) SetC(c int) *StateBuilder {
	sb.C = c
	return sb
}

func (sb *StateBuilder) SetOutput(ints []int) *StateBuilder {
	for _, i := range ints {
		sb.Output.Add(i)
	}
	return sb
}

func (sb *StateBuilder) SetInstructionIndex(i int) *StateBuilder {
	sb.InstructionIndex = i
	return sb
}

func (sb *StateBuilder) Build() State {
	return State{
		A:                sb.A,
		B:                sb.B,
		C:                sb.C,
		Output:           sb.Output,
		InstructionIndex: sb.InstructionIndex,
	}
}

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
	wantState := NewStateBuilder().SetA(729).Build()
	want := Input{
		debugger: Debugger{
			State: wantState,
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

func TestCombo(t *testing.T) {
	state := NewStateBuilder().SetA(10).SetB(20).SetC(30).Build()

	type Case struct {
		input int
		want  int
	}

	cases := []Case{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 10},
		{5, 20},
		{6, 30},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Given %v, returns %v", c.input, c.want), func(t *testing.T) {
			got := combo(c.input, state)
			utils.CheckEqual(got, c.want, t)
		})
	}
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
		state.A = 12
		op := Operation{OpCode: 2, Operand: 4}

		want := newState()
		want.A = 12
		want.B = 4
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
		state.A = 12
		op := Operation{OpCode: 5, Operand: 4}

		want := newState()
		want.A = 12
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

func TestDebugger(t *testing.T) {
	t.Run("C==9, program 2,6 -> B==1", func(t *testing.T) {
		debugger := Debugger{
			State:   NewStateBuilder().SetC(9).Build(),
			Program: []Operation{{2, 6}},
		}

		debugger.Run()

		utils.CheckEqual(debugger.State.B, 1, t)
	})

	t.Run("A==10, program 5,0,5,1,5,4 -> output==0,1,2", func(t *testing.T) {
		debugger := Debugger{
			State:   NewStateBuilder().SetA(10).Build(),
			Program: []Operation{{5, 0}, {5, 1}, {5, 4}},
		}

		debugger.Run()

		utils.CheckEqual(debugger.State.Output.AsSlice(), []int{0, 1, 2}, t)
	})

	t.Run("A==2024, program 0,1,5,4,3,0 -> output==4,2,5,6,7,7,7,7,3,1,0, A==0", func(t *testing.T) {
		debugger := Debugger{
			State:   NewStateBuilder().SetA(2024).Build(),
			Program: []Operation{{0, 1}, {5, 4}, {3, 0}},
		}

		debugger.Run()

		utils.CheckEqual(debugger.State.Output.AsSlice(), []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}, t)
		utils.CheckEqual(debugger.State.A, 0, t)
	})

	t.Run("B==29, program 1,7 -> B==26", func(t *testing.T) {
		debugger := Debugger{
			State:   NewStateBuilder().SetB(29).Build(),
			Program: []Operation{{1, 7}},
		}

		debugger.Run()

		utils.CheckEqual(debugger.State.B, 26, t)
	})

	t.Run("B==2024, C==43690, program 4,0 -> B==44354", func(t *testing.T) {
		debugger := Debugger{
			State:   NewStateBuilder().SetB(2024).SetC(43690).Build(),
			Program: []Operation{{4, 0}},
		}

		debugger.Run()

		utils.CheckEqual(debugger.State.B, 44354, t)
	})

}

func TestTruncatedDivision(t *testing.T) {
	type Case struct {
		num   int
		denom int
		want  int
	}

	cases := []Case{
		{4, 2, 2},
		{6, 4, 1},
		{8, 5, 1},
		{8, 6, 1},
		{8, 7, 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v / %v == %v", c.num, c.denom, c.want), func(t *testing.T) {
			got := c.num / c.denom
			utils.CheckEqual(got, c.want, t)
		})
	}
}

func TestPart1(t *testing.T) {
	t.Run("input_example.txt", func(t *testing.T) {
		input, _ := GetInput("input_example.txt")

		got := Part1(input)
		want := "4,6,3,5,6,3,5,2,1,0"

		utils.CheckEqual(got, want, t)
	})

	t.Run("input_example2.txt", func(t *testing.T) {
		input, _ := GetInput("input_example2.txt")

		got := Part1(input)
		want := "3,1,5,3,7,4,2,7,5"

		utils.CheckEqual(got, want, t)
	})
}

func TestPart2(t *testing.T) {
	input, _ := GetInput("input_example_part2.txt")

	got := Part2(input)
	want := 117440

	utils.CheckEqual(got, want, t)
}
