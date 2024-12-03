package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fullMemory := strings.Join(lines, "")

	return &Input{memory: fullMemory}, nil
}

func main() {
	part1()
}

func part1() {
	input, _ := GetInput("input.txt")
	matches := GetMatches(input.memory)
	result := EvaluateMatches(matches)

	fmt.Printf("Part 1: got <%v>\n", result)
}

type Match struct {
	left, right int
}

func EvaluateMatches(ms []Match) int {
	result := 0
	for _, m := range ms {
		res := m.left * m.right
		result += res
	}
	return result
}

func GetMatches(input string) []Match {
	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	var matches []Match
	results := r.FindAllStringSubmatch(input, -1)
	if results == nil {
		return matches
	}

	for _, result := range results {
		left, _ := strconv.Atoi(result[1])
		right, _ := strconv.Atoi(result[2])
		match := Match{
			left:  left,
			right: right,
		}
		matches = append(matches, match)
	}

	return matches
}

type Input struct {
	memory string
}

var ErrInputFile = errors.New("cannot open input file")
