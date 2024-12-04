package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var maxX int
	var y = 0
	grid := make(map[Coord]string, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		maxX = len(line) - 1

		for x, c := range strings.Split(line, "") {
			coord := Coord{x, y}
			grid[coord] = c
		}
		y++
	}

	return &Input{maxX: maxX, maxY: y - 1, grid: grid}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	fmt.Printf("Got input: %v", input)
}

func Part1(input string) int {
	return 0
}

type Coord struct {
	x, y int
}

type Input struct {
	maxX int
	maxY int
	grid map[Coord]string
}
