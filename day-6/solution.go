package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

type Point struct {
	x, y int
}

type Direction int

const (
	None  Direction = iota
	Up              = iota
	Down            = iota
	Left            = iota
	Right           = iota
)

type Guard struct {
	position  Point
	direction Direction
	path      []Point
}

type Map map[Point]string

type Input struct {
	pointMap Map
	guard    Guard
}

const (
	RULES_SEPARATOR = "|"
	PAGE_SEPARATOR  = ","
)

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var guard Guard
	pointMap := make(Map, 100)

	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")

		for x, c := range chars {
			point := Point{x, y}
			pointMap[point] = c

			possibleGuard := MakeGuard(c, point)
			if possibleGuard != nil {
				guard = *possibleGuard
			}
		}
	}

	return &Input{pointMap, guard}, nil
}

func MakeGuard(c string, p Point) *Guard {
	directions := map[string]Direction{
		"^": Up,
		">": Right,
		"<": Left,
		"v": Down,
	}

	if directions[c] == None {
		return nil
	}

	return &Guard{
		position:  p,
		direction: directions[c],
		path:      make([]Point, 100),
	}
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func Part1(input *Input) int {
	// TODO
	// 1. Parse map and place guard
	// 2. Walk guard using rules, until you hit boundary
	// 3. Record the steps in full
	// 4. Count steps, donezo

	return 0

}
