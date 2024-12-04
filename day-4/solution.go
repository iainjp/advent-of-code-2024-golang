package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

type Point struct {
	x, y int
}

type Input struct {
	maxX int
	maxY int
	grid GridMap
}

type GridMap = map[Point]string

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var maxX int
	var y = 0
	grid := make(GridMap, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		maxX = len(line)

		for x, c := range strings.Split(line, "") {
			coord := Point{x, y}
			grid[coord] = c
		}
		y++
	}

	maxY := y
	fmt.Printf("got maxX=<%v>, maxY=<%v>", maxX, maxY)
	return &Input{maxX: maxX, maxY: maxY, grid: grid}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)
}

func Part1(input *Input) int {
	// TODO
	// 1. get starting points of X
	// 2. get combinations of 4 letters in row (up,down,left,right,diagonals) from starting point
	// 3. count those that spell "XMAS"
	// 4. donezo.

	startingPoints := GetCoordsOfLetter(input.grid, "X")
	counter := 0

	for _, p := range startingPoints {
		sequencesToConsider := GetSeqOfPointsSurrounding(p, input.maxX, input.maxY)

		for _, ps := range sequencesToConsider {
			if GetPointsAsString(input.grid, ps...) == "XMAS" {
				// fmt.Printf("Found XMAS at: %v\n", ps)
				counter++
			}
		}
	}

	return counter
}

func GetPointsAsString(grid GridMap, ps ...Point) string {
	var result []string

	for _, p := range ps {
		result = append(result, grid[p])
	}

	return strings.Join(result, "")
}

func GetSeqOfPointsSurrounding(p Point, maxX, maxY int) [][]Point {
	var results [][]Point
	spaceLeft := p.x >= 3
	spaceRight := p.x+3 <= maxX
	spaceUp := p.y >= 3
	spaceDown := p.y+3 <= maxY
	spaceUL := spaceUp && spaceLeft
	spaceDL := spaceDown && spaceLeft
	spaceUR := spaceUp && spaceRight
	spaceDR := spaceDown && spaceRight

	if spaceLeft {
		seq := []Point{p, Point{p.x - 1, p.y}, Point{p.x - 2, p.y}, Point{p.x - 3, p.y}}
		results = append(results, seq)
	}

	if spaceRight {
		seq := []Point{p, Point{p.x + 1, p.y}, Point{p.x + 2, p.y}, Point{p.x + 3, p.y}}
		results = append(results, seq)
	}

	if spaceUp {
		seq := []Point{p, Point{p.x, p.y - 1}, Point{p.x, p.y - 2}, Point{p.x, p.y - 3}}
		results = append(results, seq)
	}

	if spaceDown {
		seq := []Point{p, Point{p.x, p.y + 1}, Point{p.x, p.y + 2}, Point{p.x, p.y + 3}}
		results = append(results, seq)
	}

	if spaceUL {
		seq := []Point{p, Point{p.x - 1, p.y - 1}, Point{p.x - 2, p.y - 2}, Point{p.x - 3, p.y - 3}}
		results = append(results, seq)
	}

	if spaceDL {
		seq := []Point{p, Point{p.x - 1, p.y + 1}, Point{p.x - 2, p.y + 2}, Point{p.x - 3, p.y + 3}}
		results = append(results, seq)
	}

	if spaceUR {
		seq := []Point{p, Point{p.x + 1, p.y - 1}, Point{p.x + 2, p.y - 2}, Point{p.x + 3, p.y - 3}}
		results = append(results, seq)
	}

	if spaceDR {
		seq := []Point{p, Point{p.x + 1, p.y + 1}, Point{p.x + 2, p.y + 2}, Point{p.x + 3, p.y + 3}}
		results = append(results, seq)
	}

	return results
}

func GetCoordsOfLetter(grid GridMap, letter string) []Point {
	var result []Point

	for p, v := range grid {
		if v == letter {
			result = append(result, p)
		}
	}

	return result
}

func pointCmp(l, r Point) int {
	if l.y < r.y {
		return -1
	}
	if l.y == r.y {
		if l.x < r.x {
			return -1
		}
	}
	return 1
}

// get keys ordered by y, x - like scanning along a big line
func GetOrderedKeys(dict GridMap) []Point {
	keys := make([]Point, len(dict))
	i := 0
	for k := range dict {
		keys[i] = k
		i++
	}

	slices.SortFunc(keys, pointCmp)

	return keys
}
