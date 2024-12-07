package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

type Equation struct {
	targetTotal int
	operands    []int
}

type Input struct {
	equations []Equation
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var equations []Equation
	for scanner.Scan() {
		//handle string
		line := scanner.Text()

		split := strings.Split(line, ":")
		targetTotal, _ := strconv.Atoi(split[0])

		var operands []int
		for _, o := range strings.Split(strings.TrimSpace(split[1]), " ") {
			oint, _ := strconv.Atoi(o)
			operands = append(operands, oint)
		}

		equations = append(equations, Equation{targetTotal, operands})
	}

	return &Input{equations: equations}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func Part1(input *Input) int {
	// 1. Parse map and place guard
	return 0
}
