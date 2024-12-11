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

type Stone struct {
	prev   *Stone
	next   *Stone
	number int
}

type StoneLine struct {
	head *Stone
}

func (sl *StoneLine) ToSlice() []Stone {
	current := sl.head
	stoneSlice := []Stone{*current}

	for current.next != nil {
		current = current.next
		stoneSlice = append(stoneSlice, *current)
	}

	return stoneSlice
}

func (sl *StoneLine) GetNumbers() []int {
	current := sl.head
	resultSlice := []int{current.number}

	for current.next != nil {
		current = current.next
		resultSlice = append(resultSlice, current.number)
	}

	return resultSlice
}

type Input struct {
	stones []int
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	var stones []int
	for _, c := range strings.Split(line, " ") {
		i, _ := strconv.Atoi(c)
		stones = append(stones, i)
	}

	return &Input{stones: stones}, nil
}

// build LinkedList of stones, returning head
func BuildStones(ints []int) *StoneLine {
	var head *Stone
	var prev *Stone
	for i, n := range ints {
		stone := &Stone{
			number: n,
		}
		if i == 0 {
			head = stone
		} else {
			stone.prev = prev
			prev.next = stone
		}
		prev = stone
	}

	return &StoneLine{head: head}
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)
}

func Part1(input *Input) int {
	return 0
}
