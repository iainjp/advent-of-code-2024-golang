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
	// TODO -- try baking the direction into the neighbour more explicitly - e.g a vertex for each (*Tile, direction)
	// TOOD -- can then just use DFS or Djikstra across those vertices
	coord Coord
	cost  int
}

func (a Tile) Equal(o Tile) bool {
	return a.coord == o.coord && a.cost == o.cost
}

type Maze struct {
	start *Tile
	end   *Tile
	all   []*Tile
}

// Useful: https://www.freecodecamp.org/news/dijkstras-shortest-path-algorithm-visual-introduction/
// Returns map of *Tile -> cost to get there
func (m *Maze) Dijkstra() map[*Tile]int {
	facing := EAST
	start := m.start

	distances := make(map[*Tile]int)
	for _, tile := range m.all {
		distances[tile] = int(math.Inf(1))
	}
	distances[start] = 0

	// toVisit := append([]*Tile{}, m.all...)
	toVisit := m.all
	byDistance := func(i, j int) bool {
		return distances[toVisit[i]] < distances[toVisit[j]]
	}

	for len(toVisit) > 0 {
		sort.SliceStable(toVisit, byDistance)

		curr := toVisit[0]
		toVisit = toVisit[1:]

		for direction, next := range curr.neighbours {
			var cost int
			if next == nil {
				continue
			}

			// directions are messed up - this stage should be _considering_ next
			// and marking the cost, not changing directions constantly.
			turns := utils.Abs(facing - direction)

			switch turns {
			case 0:
				// straight ahead
				cost = 1
			case 1:
				// clock-wise
				cost = 1001
				facing = direction
			case 2:
				// go backwards - skip
				cost = 2001
				facing = direction
			case 3:
				// anti-clockwise
				cost = 1001
				facing = direction
			}

			alt := distances[curr] + cost
			existing := distances[next]
			if alt < existing {
				distances[next] = alt
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
