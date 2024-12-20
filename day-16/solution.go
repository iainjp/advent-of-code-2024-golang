package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"iain.fyi/aoc2024/structure"
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
	from      *Tile
	to        *Tile
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
		if v.to == m.end {
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
		currentTile := current.to
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

	p2Result := Part2(input)
	fmt.Printf("Part 2: got %v\n", p2Result)

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
		from := register[k]
		for dir, coord := range adj {
			to, ok := register[coord]
			// if not in register (e.g. out-of-bounds) OR wall
			if !ok || to.symbol == WALL {
				continue
			}

			if to.symbol == START {
				start = to
			}

			if to.symbol == END {
				end = to
			}

			tiles[to] = true

			vert := Vertex{
				from:      from,
				to:        to,
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

// get vertices pointing to `start` (without doubley-linked graph, ugh)
func GetVerticesPointingTo(to *Tile, distMap map[*Vertex]int) []*Vertex {
	var inbound []*Vertex
	for vert, _ := range distMap {
		if vert.to == to {
			inbound = append(inbound, vert)
		}
	}
	return inbound
}

// return number of tiles within any best path
func Part2(input *Input) int {
	// TODO walk backward through Djikstra
	// - keep map of tiles
	// - at each step, add next node that could have been part of best path
	// -- so -- distances[vertex] IN (distance-1, distance-1001)
	// - then move to each of those nodes and start again.

	distances := input.maze.Dijkstra()
	current := input.maze.end
	end := input.maze.start

	currentCost := Part1(input)

	tiles := structure.NewHashSet[*Tile]()
	tiles.AddAll(current, end)
	// toVisit := GetVerticesPointingTo(current, distances)

	var toVisit []*Vertex
	for current != end {
		inbound := GetVerticesPointingTo(current, distances)
		for _, vert := range inbound {
			dist := distances[vert]

			// if `incomingVert` was in a best path, it cost 1 or 1001 to get to current
			if dist == currentCost || dist == currentCost-1 || dist == currentCost-1001 {
				tiles.Add(vert.from)
				toVisit = append(toVisit, vert)
			}
		}

		if len(toVisit) == 0 {
			break
		}

		current = toVisit[0].from
		currentCost = distances[toVisit[0]]
		toVisit = toVisit[1:]
	}

	return tiles.Size()
}
