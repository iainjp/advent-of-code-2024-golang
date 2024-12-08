package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

type Coord struct {
	x, y int
}

type Point struct {
	symbol    string   // symbol at point
	anitnodes []*Coord // coords of antinodes (if any)
}

type PointMap = map[Coord]Point

type Input struct {
	pointMap *PointMap
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pointMap := make(PointMap)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		for x, c := range split {
			coord := Coord{x, y}
			point := Point{symbol: c}
			pointMap[coord] = point
		}
		y += 1
	}

	return &Input{&pointMap}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func Part1(input *Input) uint64 {
	// TODO
	// 1: Get all antennas
	// 2: For each, calculate the antinodes of all others (oof)

	return 0
}
