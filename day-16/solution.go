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
	symbol   string
	coord    Coord
	vertices []*Vertex
}

type Maze struct {
	start       *Tile
	end         *Tile
	all         []*Tile
	allVertices []*Vertex
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
		// go backwards - later skipped
		cost = math.MaxInt32
	case 3:
		// anti-clockwise
		cost = 1001
	}

	return cost
}

func (m *Maze) EndVertices() []*Vertex {
	var endVertices []*Vertex
	for _, v := range m.allVertices {
		if v.tile == m.end {
			endVertices = append(endVertices, v)
		}
	}
	return endVertices
}

// Returns map of *Vertex -> cost to get there from start
func (m *Maze) Dijkstra() map[*Vertex]int {
	start := m.start

	// try tracking cost by vertex (includes direction)
	vertexCosts := make(map[*Vertex]int)
	for _, vertex := range m.allVertices {
		vertexCosts[vertex] = math.MaxInt32
	}

	toVisit := start.vertices
	// set initial distances
	for _, v := range toVisit {
		vertexCosts[v] = Cost(EAST, v.direction)
	}

	vertexByDistance := func(i, j int) bool {
		return vertexCosts[toVisit[i]] < vertexCosts[toVisit[j]]
	}

	traversed := make(map[Vertex]bool)

	for len(toVisit) > 0 {
		sort.SliceStable(toVisit, vertexByDistance)

		current := toVisit[0]
		currentTile := current.tile
		facing := current.direction
		toVisit = toVisit[1:]

		for _, vertex := range currentTile.vertices {
			cost := Cost(facing, vertex.direction)

			// calculate possible vertex cost
			altVCost := vertexCosts[current] + cost
			existingVCost := vertexCosts[vertex]

			// if less, found new best cost
			if altVCost < existingVCost {
				vertexCosts[vertex] = altVCost
			}

			// if we haven't traversed the vertex yet, at it to queue
			previouslyTraversed := traversed[*vertex]
			if !previouslyTraversed {
				traversed[*vertex] = true
				toVisit = append(toVisit, vertex)
			}
		}
	}

	return vertexCosts
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
	input, _ := GetInput("input.txt")
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

	var allVertices []*Vertex

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
			allVertices = append(allVertices, &vert)
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
			start:       start,
			end:         end,
			all:         all,
			allVertices: allVertices,
		},
	}, nil
}

// return cost of the best path to end
func Part1(input *Input) int {
	distances := input.maze.Dijkstra()

	endVertices := input.maze.EndVertices()
	minEndDistance := math.MaxInt32
	for _, ev := range endVertices {
		d := distances[ev]
		if d < minEndDistance {
			minEndDistance = d
		}
	}

	return minEndDistance
}

// return number of tiles within any best path
func Part2(input *Input) int {
	// TODO walk backward through Djikstra
	// - keep map of tiles
	// - at each step, add next node that could have been part of best path
	// -- so -- distances[vertex] IN (distance-1, distance-1001)
	// - then move to each of those nodes and start again.

	return 0
}
