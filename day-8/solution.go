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
	symbol     string // symbol at point
	isAntinode bool
}

func (p *Point) IsAntenna() bool {
	return p.symbol != "."
}

type PointMap = map[Coord]Point

func GetAntennas(pm PointMap) []Coord {
	var antennas []Coord
	for coord, point := range pm {
		if point.IsAntenna() {
			antennas = append(antennas, coord)
		}
	}
	return antennas
}

func SetAntinode(pm PointMap, coord Coord) {
	point := pm[coord]
	point.isAntinode = true
	pm[coord] = point
}

// TODO write test?
func GetUniqueAntinodes(pm PointMap) []Coord {
	var coords []Coord
	for k, v := range pm {
		if v.isAntinode {
			coords = append(coords, k)
		}
	}

	return coords
}

type Input struct {
	pointMap PointMap
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

	return &Input{pointMap}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

type Pair[T any] struct {
	first, second T
}

func GetAllUniquePairs(coords []Coord) []Pair[Coord] {
	var antennaPairs []Pair[Coord]

	for i, a := range coords {
		leftSlice := coords[:i]
		for _, a2 := range leftSlice {
			antennaPairs = append(antennaPairs, Pair[Coord]{a, a2})
		}
		rightSlice := coords[i+1:]
		for _, a2 := range rightSlice {
			antennaPairs = append(antennaPairs, Pair[Coord]{a, a2})
		}
	}

	return antennaPairs
}

func GetAntinodes(pair Pair[Coord]) Pair[Coord] {
	firstXDiff := pair.first.x - pair.second.x
	firstYDiff := pair.first.y - pair.second.y

	secondXDiff := pair.second.x - pair.first.x
	secondYDiff := pair.second.y - pair.first.y

	first := Coord{
		x: pair.first.x + firstXDiff,
		y: pair.first.y + firstYDiff,
	}

	second := Coord{
		x: pair.second.x + secondXDiff,
		y: pair.second.y + secondYDiff,
	}

	return Pair[Coord]{first, second}
}

func Part1(input *Input) int {
	// TODO
	// 1: Get all antennas
	// 2: Get all unique pairs
	// 3: For each pair, set the antinodes for both

	antennas := GetAntennas(input.pointMap)
	aPairs := GetAllUniquePairs(antennas)

	for _, pair := range aPairs {
		antinodes := GetAntinodes(pair)

		SetAntinode(input.pointMap, antinodes.first)
		SetAntinode(input.pointMap, antinodes.second)
	}

	antinodes := GetUniqueAntinodes(input.pointMap)
	return len(antinodes)
}
