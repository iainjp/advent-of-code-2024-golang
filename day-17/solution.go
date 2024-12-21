package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"iain.fyi/aoc2024/structure"
)

var ErrInputFile = errors.New("cannot open input file")

// non-negative modulo
func mod(a, b int) int {
	return (a%b + b) % b
}

type State struct {
	A, B, C          int
	Output           structure.List[int]
	InstructionIndex int
}

type Operation struct {
	OpCode  int
	Operand int
}

func combo(operand int, state State) int {
	if operand < 4 {
		return operand
	}

	var comb int
	switch operand {
	case 4:
		comb = state.A
	case 5:
		comb = state.B
	case 6:
		comb = state.C
	}
	return comb
}

func (o *Operation) adv(state *State) {
	num := state.A
	denom := int(math.Pow(2, float64(combo(o.Operand, *state))))

	state.A = num / denom
	state.InstructionIndex += 1
}

func (o *Operation) bxl(state *State) {
	newB := state.B ^ o.Operand
	state.B = newB
	state.InstructionIndex += 1
}

func (o *Operation) bst(state *State) {
	mod := mod(combo(o.Operand, *state), 8)
	state.B = mod
	state.InstructionIndex += 1
}

func (o *Operation) jnz(state *State) {
	if state.A == 0 {
		state.InstructionIndex += 1
		return
	}
	// divide by 2, as we're working on []Operation, not raw []ints
	state.InstructionIndex = o.Operand / 2
}

func (o *Operation) bxc(state *State) {
	state.B = state.B ^ state.C
	state.InstructionIndex += 1
}

func (o *Operation) out(state *State) {
	val := mod(combo(o.Operand, *state), 8)

	// split and add each digit to output
	valAsString := strconv.Itoa(val)
	for _, s := range strings.Split(valAsString, "") {
		i, _ := strconv.Atoi(s)
		state.Output.Add(i)
	}

	state.InstructionIndex += 1
}

func (o *Operation) bdv(state *State) {
	num := state.A
	denom := int(math.Pow(2, float64(combo(o.Operand, *state))))

	state.B = num / denom
	state.InstructionIndex += 1
}

func (o *Operation) cdv(state *State) {
	num := state.A
	denom := int(math.Pow(2, float64(combo(o.Operand, *state))))

	state.C = num / denom
	state.InstructionIndex += 1
}

func (o *Operation) Execute(state *State) {
	switch o.OpCode {
	case 0:
		o.adv(state)
	case 1:
		o.bxl(state)
	case 2:
		o.bst(state)
	case 3:
		o.jnz(state)
	case 4:
		o.bxc(state)
	case 5:
		o.out(state)
	case 6:
		o.bdv(state)
	case 7:
		o.cdv(state)
	}
}

type Debugger struct {
	State   State
	Program []Operation
}

func (d *Debugger) Run() {
	for d.State.InstructionIndex < len(d.Program) {
		d.Program[d.State.InstructionIndex].Execute(&d.State)
	}
}

func (d *Debugger) RawProgram() []int {
	var raw []int
	for _, o := range d.Program {
		raw = append(raw, o.OpCode)
		raw = append(raw, o.Operand)
	}
	return raw
}

func (d *Debugger) ResetIndex() {
	d.State.InstructionIndex = 0
}

type Input struct {
	debugger Debugger
}

func (i *Input) Clone() *Input {
	state := State{
		A:                i.debugger.State.A,
		B:                i.debugger.State.B,
		C:                i.debugger.State.C,
		InstructionIndex: i.debugger.State.InstructionIndex,
		Output:           structure.NewList[int](),
	}

	for _, i := range i.debugger.State.Output.AsSlice() {
		state.Output.Add(i)
	}

	debugger := Debugger{
		State:   state,
		Program: i.debugger.Program,
	}

	return &Input{
		debugger: debugger,
	}
}

func main() {

	input, _ := GetInput("input.txt")
	original := input.Clone()

	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

	p2Result := Part2(original)
	fmt.Printf("Part 2: got %v\n", p2Result)

}

func ParseRegisterValue(line string) int {
	strValue := strings.TrimSpace(strings.Split(line, ":")[1])
	value, _ := strconv.Atoi(strValue)
	return value
}

func ParseProgram(line string) []Operation {
	strValues := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), ",")

	var operations []Operation
	for pair := range slices.Chunk(strValues, 2) {
		opcode, _ := strconv.Atoi(pair[0])
		operand, _ := strconv.Atoi(pair[1])
		operations = append(operations, Operation{OpCode: opcode, Operand: operand})
	}

	return operations
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	aLine := scanner.Text()
	a := ParseRegisterValue(aLine)

	scanner.Scan()
	bLine := scanner.Text()
	b := ParseRegisterValue(bLine)

	scanner.Scan()
	cLine := scanner.Text()
	c := ParseRegisterValue(cLine)

	scanner.Scan()
	scanner.Scan() // twice to skip empty line
	pLine := scanner.Text()

	program := ParseProgram(pLine)

	input := Input{
		debugger: Debugger{
			State: State{
				A:                a,
				B:                b,
				C:                c,
				Output:           structure.NewList[int](),
				InstructionIndex: 0,
			},
			Program: program,
		},
	}

	return &input, nil
}

// return output of program
func Part1(input *Input) string {
	input.debugger.Run()
	var parts []string
	for _, i := range input.debugger.State.Output.AsSlice() {
		parts = append(parts, strconv.Itoa(i))
	}
	return strings.Join(parts, ",")
}

func equal(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx := range len(s1) {
		if s1[idx] != s2[idx] {
			return false
		}
	}

	return true
}

// find register A value that makes program output itself
func Part2(input *Input) int {
	originalB := input.debugger.State.B
	originalC := input.debugger.State.C

	program := input.debugger.RawProgram()
	fmt.Printf("Looking for %v ...", program)

	aValue := 0
	input.debugger.State.A = aValue
	input.debugger.Run()

	output := input.debugger.State.Output.AsSlice()

	for !equal(output, program) {
		if mod(aValue, 1000000) == 0 {
			fmt.Printf("Trying A==%v \n", aValue)
		}

		// reset everything to original
		// increment A
		aValue += 1
		input.debugger.State.A = aValue
		input.debugger.State.B = originalB
		input.debugger.State.C = originalC
		input.debugger.State.InstructionIndex = 0
		input.debugger.State.Output = structure.NewList[int]()

		input.debugger.Run()
		output = input.debugger.State.Output.AsSlice()
	}

	fmt.Printf("Quiting with A=%v \n", aValue)
	return aValue
}
