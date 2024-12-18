package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"iain.fyi/aoc2024/utils"
)

var ErrInputFile = errors.New("cannot open input file")

const (
	WALL  = "#"
	EMPTY = "."
	START = "S"
	END   = "E"
)

const (
	_     = iota
	NORTH = iota // 1
	EAST  = iota // 2
	SOUTH = iota // 3
	WEST  = iota // 4
)

type Tile struct {
	symbol string
	// indexed by directional const
	neighbours []*Tile
	coord      Coord
	cost       int
}

func (a Tile) Equal(o Tile) bool {
	return a.coord == o.coord && a.cost == o.cost
}

type Maze struct {
	start *Tile
	end   *Tile
	all   []*Tile
}

// Returns map of *Tile -> cost to get there
func (m *Maze) Dijkstra() map[*Tile]int {
	direction := EAST
	start := m.start

	distances := make(map[*Tile]int, len(m.all))
	for _, tile := range m.all {
		distances[tile] = math.MaxInt32
	}
	distances[start] = 0

	// set initial costs and distances
	for i, n := range start.neighbours {
		if n == nil {
			continue
		}
		if i == direction {
			n.cost = 1
			distances[n] = 1
		}
		// single turn
		if utils.Abs(direction-i) == 1 || utils.Abs(direction-1) == 3 {
			n.cost = 1001
			distances[n] = 1001
		}

		// 2 turns
		if utils.Abs(direction-i) == 2 {
			n.cost = 2001
			distances[n] = 2001
		}
	}

	var options []*Tile
	for _, n := range start.neighbours {
		if n != nil {
			options = append(options, n)
		}
	}

	type Step struct {
		coord Coord
		cost  int
	}

	tileSeen := make(map[Step]bool, len(m.all))
	// not sure on this condition
	for len(options) > 0 {
		sort.SliceStable(options, func(i, j int) bool {
			return distances[options[i]] < distances[options[j]]
		})

		curr := options[0]
		options = options[1:]

		for i, next := range curr.neighbours {
			if next == nil {
				continue
			}

			if i == direction {
				next.cost = 1
			}
			// single turn
			if utils.Abs(direction-i) == 1 || utils.Abs(direction-i) == 3 {
				next.cost = 1001
				direction = i
			}

			// 2 turns
			if utils.Abs(direction-i) == 2 {
				next.cost = 2001
				direction = i
			}

			alt := distances[curr] + next.cost
			if alt < distances[next] {
				distances[next] = alt
			}
			// add to options if we haven't yet considered it at current cost previously
			step := Step{
				coord: next.coord,
				cost:  next.cost,
			}
			seen := tileSeen[step]
			if !seen {
				tileSeen[step] = true
				options = append(options, next)
			}
		}
	}

	return distances
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
				coord:      coord,
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
	var end *Tile
	tiles := make(map[*Tile]bool)
	// create graph
	for k, v := range register {
		adj := k.Adjacent()
		for i, n := range adj {
			t, ok := register[n]
			// if not in register (e.g. out-of-bounds) OR wall
			if !ok || t.symbol == WALL {
				v.neighbours[i] = nil
				continue
			}

			if t.symbol == START {
				start = t
			}

			if t.symbol == END {
				end = t
			}

			tiles[t] = true
			v.neighbours[i] = t
		}
	}

	var all []*Tile
	for t, v := range tiles {
		if v {
			all = append(all, t)
		}
	}

	return &Input{
		maze: &Maze{
			start: start,
			end:   end,
			all:   all,
		},
	}, nil
}

func Part1(input *Input) int {
	distances := input.maze.Dijkstra()
	return distances[input.maze.start]
}
