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
	input, _ := GetInput("input.txt")
	Part1(input.memory)
	Part2(input.memory)
}

func Part1(input string) int {
	matches := GetMatches(input)
	result := EvaluateMatches(matches)

	fmt.Printf("Part 1: got <%v>\n", result)
	return result
}

func Part2(input string) int {
	cleanInput := RemoveAfterDontUntilDoOrEnd(input)
	matches := GetMatches(cleanInput)
	result := EvaluateMatches(matches)

	fmt.Printf("Part 2: got <%v>\n", result)
	return result
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

// remove `dont().*` up until do(), returned as new string
func RemoveAfterDontUntilDoOrEnd(input string) string {
	pattern := `(don't\(\)).+?(do\(\))`
	r, _ := regexp.Compile(pattern)

	outcomeSlice := strings.Split(input, "")
	var outcomeString string
	inputBytes := []byte(input)

	results := r.FindSubmatchIndex(inputBytes)

	if results == nil {
		return input
	}
	for results != nil {
		dontStartIndex := 2
		doEndIndex := 5

		newSlice := make([]string, 0)

		newSlice = append(newSlice, outcomeSlice[:results[dontStartIndex]]...)
		newSlice = append(newSlice, outcomeSlice[results[doEndIndex]:]...)

		outcomeSlice = newSlice
		outcomeString = strings.Join(outcomeSlice, "")
		results = r.FindSubmatchIndex([]byte(outcomeString))
	}

	// check for outstanding `don't().*`
	lastR, _ := regexp.Compile(`(don't\(\).*)`)
	results = lastR.FindSubmatchIndex([]byte(outcomeString))

	if results == nil {
		return outcomeString
	}

	dontStartIndex := 2
	newSlice := make([]string, 0)
	newSlice = append(newSlice, outcomeSlice[:results[dontStartIndex]]...)

	outcomeSlice = newSlice
	outcomeString = strings.Join(outcomeSlice, "")

	return outcomeString
}

type Input struct {
	memory string
}

var ErrInputFile = errors.New("cannot open input file")
