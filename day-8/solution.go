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

func (p *Point) SameFrequency(op *Point) bool {
	return p.IsAntenna() && op.IsAntenna() &&
		p.symbol == op.symbol
}

type PointMap struct {
	m map[Coord]Point
}

func (pm *PointMap) GetAntennas() []Coord {
	var antennas []Coord
	for coord, point := range pm.m {
		if point.IsAntenna() {
			antennas = append(antennas, coord)
		}
	}
	return antennas
}

func (pm *PointMap) GetUniqueAntinodes() []Coord {
	var coords []Coord
	for k, v := range pm.m {
		if v.isAntinode {
			coords = append(coords, k)
		}
	}

	return coords
}

func (pm *PointMap) SetAntinodeIfInBounds(coord Coord) {
	point, ok := pm.m[coord]
	if ok {
		point.isAntinode = true
		pm.m[coord] = point
	}
}

func (pm *PointMap) Get(coord Coord) *Point {
	point, ok := pm.m[coord]
	if ok {
		return &point
	}

	return nil
}

func (pm *PointMap) Put(coord Coord, point Point) {
	pm.m[coord] = point
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

	pointMap := PointMap{
		m: make(map[Coord]Point),
	}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		for x, c := range split {
			coord := Coord{x, y}
			point := Point{symbol: c}
			pointMap.Put(coord, point)
		}
		y += 1
	}

	return &Input{pointMap}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

	p2Result := Part2(input)
	fmt.Printf("Part 2: got %v\n", p2Result)
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

func GetAntinodes(pair Pair[Coord]) []Coord {
	xDiff := pair.first.x - pair.second.x
	yDiff := pair.first.y - pair.second.y

	first := Coord{
		x: pair.first.x + xDiff,
		y: pair.first.y + yDiff,
	}

	second := Coord{
		x: pair.second.x - xDiff,
		y: pair.second.y - yDiff,
	}

	return []Coord{first, second}
}

func GetAntinodesWithResonantHarmonics(pair Pair[Coord], pm PointMap) []Coord {
	var antinodeCoords []Coord

	// Need to also include teh antennas
	antinodeCoords = append(antinodeCoords, pair.first)
	antinodeCoords = append(antinodeCoords, pair.second)

	xDiff := pair.first.x - pair.second.x
	yDiff := pair.first.y - pair.second.y

	positive := Coord{
		x: pair.first.x + xDiff,
		y: pair.first.y + yDiff,
	}
	positivePoint := pm.Get(positive)
	// keep adding coords in positive direction until they out out-of-bounds
	for positivePoint != nil {
		antinodeCoords = append(antinodeCoords, positive)
		positive = Coord{
			x: positive.x + xDiff,
			y: positive.y + yDiff,
		}
		positivePoint = pm.Get(positive)
	}

	negative := Coord{
		x: pair.second.x - xDiff,
		y: pair.second.y - yDiff,
	}
	negativePoint := pm.Get(negative)
	// keep adding coords in negative direction until they out out-of-bounds
	for negativePoint != nil {
		antinodeCoords = append(antinodeCoords, negative)
		negative = Coord{
			x: positive.x + xDiff,
			y: positive.y + yDiff,
		}
		negativePoint = pm.Get(positive)
	}

	return antinodeCoords
}

func Part1(input *Input) int {
	// 1: Get all antennas
	// 2: Get all unique pairs
	// 3: For each pair, if same frequency, get antinodes
	// 4: If antinode in bounds, set antinode flag in map
	// 4: Count unique antinodes

	antennas := input.pointMap.GetAntennas()
	aPairs := GetAllUniquePairs(antennas)

	for _, pair := range aPairs {
		first := input.pointMap.Get(pair.first)
		second := input.pointMap.Get(pair.second)

		if first.SameFrequency(second) {
			antinodes := GetAntinodes(pair)
			for _, an := range antinodes {
				input.pointMap.SetAntinodeIfInBounds(an)
			}
		}
	}

	antinodes := input.pointMap.GetUniqueAntinodes()
	return len(antinodes)
}

func Part2(input *Input) int {
	// 1: Get all antennas
	// 2: Get all unique pairs
	// 3: For each pair, if same frequency, get all resonant antinodes
	// 4: For each antinode in bounds, set antinode flag in map (including on Antenna)
	// 4: Count unique antinodes

	antennas := input.pointMap.GetAntennas()
	aPairs := GetAllUniquePairs(antennas)

	for _, pair := range aPairs {
		first := input.pointMap.Get(pair.first)
		second := input.pointMap.Get(pair.second)

		if first.SameFrequency(second) {
			antinodes := GetAntinodesWithResonantHarmonics(pair, input.pointMap)
			for _, an := range antinodes {
				input.pointMap.SetAntinodeIfInBounds(an)
			}
		}
	}

	antinodes := input.pointMap.GetUniqueAntinodes()
	return len(antinodes)
}
