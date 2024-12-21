package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"iain.fyi/aoc2024/structure"
)

var ErrInputFile = errors.New("cannot open input file")

type Operation struct {
	OpCode  int
	Operand int
}

type State struct {
	A, B, C int
	Output  structure.List[int]
}

type Debugger struct {
	State            State
	Program          []Operation
	InstructionIndex int
}

type Input struct {
	debugger *Debugger
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

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
		debugger: &Debugger{
			State: State{
				A:      a,
				B:      b,
				C:      c,
				Output: structure.NewList[int](),
			},
			Program:          program,
			InstructionIndex: 0,
		},
	}

	return &input, nil
}

// return output of program
func Part1(input *Input) int {
	return 0
}
