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
	NORTH = iota // 1
	EAST  = iota // 2
	SOUTH = iota // 3
	WEST  = iota // 4
)

type Vertex struct {
	tile      *Tile
	direction int
}

type Tile struct {
	symbol string
	coord  Coord
	// TODO -- try baking the direction into the neighbour more explicitly - e.g a vertex for each (*Tile, direction)
	// TOOD -- can then just use DFS or Djikstra across those vertices
	vertices []*Vertex
}

type Maze struct {
	start *Tile
	end   *Tile
	all   []*Tile
}

func Cost(facing, direction int) int {
	var cost int
	turns := utils.Abs(facing - direction)
	switch turns {
	case 0:
		// straight ahead
		cost = 1
	case 1:
		// clock-wise
		cost = 1001
	case 2:
		// go backwards - making prohibitively expensive
		cost = -1
	case 3:
		// anti-clockwise
		cost = 1001
	}

	return cost
}

// Useful: https://www.freecodecamp.org/news/dijkstras-shortest-path-algorithm-visual-introduction/
// Returns map of *Tile -> cost to get there
func (m *Maze) Dijkstra() map[*Tile]int {
	start := m.start

	distances := make(map[*Tile]int)
	for _, tile := range m.all {
		distances[tile] = math.MaxInt32
	}
	distances[start] = 0

	toVisit := start.vertices
	// set initial distances
	for _, v := range toVisit {
		distances[v.tile] = Cost(EAST, v.direction)
	}

	vertexByDistance := func(i, j int) bool {
		return distances[toVisit[i].tile] < distances[toVisit[j].tile]
	}

	traversed := make(map[Vertex]bool)

	for len(toVisit) > 0 {
		sort.SliceStable(toVisit, vertexByDistance)

		current := toVisit[0]
		currentTile := current.tile
		facing := current.direction
		toVisit = toVisit[1:]

		for _, vertex := range currentTile.vertices {
			direction := vertex.direction
			nextTile := vertex.tile
			cost := Cost(facing, direction)
			if cost < 0 {
				continue
			}

			altCost := distances[currentTile] + cost
			existingCost := distances[nextTile]
			if altCost < existingCost {
				distances[nextTile] = altCost
			}

			previouslyTraversed := traversed[*vertex]
			if !previouslyTraversed {
				traversed[*vertex] = true
				toVisit = append(toVisit, vertex)
			}

			if nextTile.symbol == END {
				fmt.Printf("Hit END! AltCode: %v; Existing: %v\n", altCost, existingCost)
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
				symbol: c,
				coord:  coord,
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
		for dir, coord := range adj {
			t, ok := register[coord]
			// if not in register (e.g. out-of-bounds) OR wall
			if !ok || t.symbol == WALL {
				continue
			}

			if t.symbol == START {
				start = t
			}

			if t.symbol == END {
				end = t
			}

			tiles[t] = true

			vert := Vertex{
				tile:      t,
				direction: dir,
			}
			v.vertices = append(v.vertices, &vert)
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
	return distances[input.maze.end]
}
