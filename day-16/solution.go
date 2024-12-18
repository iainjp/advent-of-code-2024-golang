package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

const (
	WALL  = "#"
	EMPTY = "."
	START = "S"
	END   = "E"
)

const (
	NORTH = iota
	EAST  = iota
	SOUTH = iota
	WEST  = iota
)

type Tile struct {
	symbol string
	// indexed by directional const
	neighbours []*Tile
}

type Maze struct {
	start *Tile
}

type Input struct {
	maze *Maze
}

type Coord struct {
	x, y int
}

// in same order as directional const definition
func (c *Coord) Adjacent() []Coord {
	return []Coord{
		{c.x, c.y - 1},
		{c.x + 1, c.y},
		{c.x, c.y + 1},
		{c.x - 1, c.y},
	}
}

func main() {
	input, _ := GetInput("input_example.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	register := make(map[Coord]*Tile)
	y := 0
	maxX := 0

	// build initial Tiles
	for scanner.Scan() {
		line := scanner.Text()
		x := 0
		for _, c := range strings.Split(line, "") {
			// create coord and tile
			coord := Coord{x: x, y: y}
			tile := Tile{
				symbol:     c,
				neighbours: make([]*Tile, 4),
			}
			register[coord] = &tile
			x++
			if x > maxX {
				maxX = x
			}
		}
		y++
	}

	var start *Tile

	// create graph
	for k, v := range register {
		adj := k.Adjacent()
		for i, n := range adj {
			t, ok := register[n]
			if !ok || t.symbol == WALL {
				// set to nil if not in register (e.g. out-of-bounds) OR wall
				v.neighbours[i] = nil
				continue
			}

			if t.symbol == START {
				start = t
			}

			v.neighbours[i] = t
		}
	}

	return &Input{
		maze: &Maze{
			start: start,
		},
	}, nil
}

func Part1(input *Input) int {
	return len(input.maze.start.neighbours)
}
