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
	SPACE = "."
	WALL  = "#"
	BOX   = "O"
	ROBOT = "@"
)

type Coord struct {
	x, y int
}

func (c *Coord) GPS() int {
	return (c.y * 100) + c.x
}

type GridPoint struct {
	symbol string
}

type Grid struct {
	cgMap map[Coord]GridPoint
	maxX  int
	maxY  int
}

func (g *Grid) GetRobot() Coord {
	var rc Coord
	for k, v := range g.cgMap {
		if v.symbol == ROBOT {
			rc = k
		}
	}
	return rc
}

func (g *Grid) GetBoxes() []Coord {
	var coords []Coord
	for k, v := range g.cgMap {
		if v.symbol == BOX {
			coords = append(coords, k)
		}
	}
	return coords
}

func (g *Grid) Row(row int) []string {
	var elements []string
	for x := range g.maxX {
		coord := Coord{x: x, y: row}
		elements = append(elements, g.cgMap[coord].symbol)
	}
	return elements
}

func (g *Grid) Column(col int) []string {
	var elements []string
	for y := range g.maxY {
		coord := Coord{x: col, y: y}
		elements = append(elements, g.cgMap[coord].symbol)
	}
	return elements
}

func (g *Grid) Left() {
	robot := g.GetRobot()
	row := g.Row(robot.y)

	// assume we can't move
	firstSpace := -1
	for i := robot.x; i >= 0; i-- {
		if row[i] == SPACE {
			firstSpace = i
			break
		}
		if row[i] == WALL {
			break
		}
	}

	// can move
	if firstSpace > -1 {
		for i := firstSpace; i < robot.x; i++ {
			currCoord := Coord{x: i, y: robot.y}
			rightCoord := Coord{x: i + 1, y: robot.y}
			g.cgMap[currCoord] = g.cgMap[rightCoord]
		}

		// set robot old position to "."
		g.cgMap[robot] = GridPoint{symbol: SPACE}
	}
}

func (g *Grid) Right() {
	robot := g.GetRobot()
	row := g.Row(robot.y)

	// assume we can't move
	firstSpace := -1
	for i := robot.x; i <= g.maxX; i++ {
		if row[i] == SPACE {
			firstSpace = i
			break
		}
		if row[i] == WALL {
			break
		}
	}

	// can move
	if firstSpace > -1 {
		for i := firstSpace; i > robot.x; i-- {
			currCoord := Coord{x: i, y: robot.y}
			left := Coord{x: i - 1, y: robot.y}
			g.cgMap[currCoord] = g.cgMap[left]
		}

		// set robot old position to "."
		g.cgMap[robot] = GridPoint{symbol: SPACE}
	}
}

func (g *Grid) Down() {
	robot := g.GetRobot()
	col := g.Column(robot.x)

	// assume we can't move
	firstSpace := -1
	for i := robot.y; i <= g.maxY; i++ {
		if col[i] == SPACE {
			firstSpace = i
			break
		}
		if col[i] == WALL {
			break
		}
	}

	// can move
	if firstSpace > -1 {
		for i := firstSpace; i > robot.y; i-- {
			currCoord := Coord{x: robot.x, y: i}
			up := Coord{x: robot.x, y: i - 1}
			g.cgMap[currCoord] = g.cgMap[up]
		}

		// set robot old position to "."
		g.cgMap[robot] = GridPoint{symbol: SPACE}
	}
}

func (g *Grid) Up() {
	robot := g.GetRobot()
	col := g.Column(robot.x)

	// assume we can't move
	firstSpace := -1
	for i := robot.y; i >= 0; i-- {
		if col[i] == SPACE {
			firstSpace = i
			break
		}
		if col[i] == WALL {
			break
		}
	}

	// can move
	if firstSpace > -1 {
		for i := firstSpace; i < robot.y; i++ {
			currCoord := Coord{x: robot.x, y: i}
			down := Coord{x: robot.x, y: i + 1}
			g.cgMap[currCoord] = g.cgMap[down]
		}

		// set robot old position to "."
		g.cgMap[robot] = GridPoint{symbol: SPACE}
	}
}

type Input struct {
	grid  Grid
	moves *[]string
}

func (i *Input) Run() {
	move := i.PopMove()
	for move != "" {
		switch move {
		case "^":
			i.grid.Up()
		case ">":
			i.grid.Right()
		case "v":
			i.grid.Down()
		case "<":
			i.grid.Left()
		}

		move = i.PopMove()
	}
}

func (i *Input) PopMove() string {
	const EMPTY = ""
	moves := *i.moves
	if len(moves) == 0 {
		return EMPTY
	}
	move := moves[0]
	poppedMoves := moves[1:]
	i.moves = &poppedMoves
	return move
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	m := make(map[Coord]GridPoint)

	y := 0
	scanner := bufio.NewScanner(file)
	maxX := 0
	for scanner.Scan() {
		x := 0
		line := strings.Split(scanner.Text(), "")

		if len(line) == 0 {
			break
		}

		for _, c := range line {
			coord := Coord{x: x, y: y}
			point := GridPoint{
				symbol: c,
			}
			m[coord] = point
			x += 1
		}
		y += 1
		if x > maxX {
			maxX = x
		}
	}

	var moves []string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		if len(line) == 0 {
			break
		}
		moves = append(moves, line...)
	}

	input := Input{
		grid: Grid{
			cgMap: m,
			maxX:  maxX,
			maxY:  y,
		},
		moves: &moves,
	}

	return &input, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func Part1(input *Input) int {
	input.Run()
	boxes := input.grid.GetBoxes()

	gpsSum := 0
	for _, c := range boxes {
		gpsSum += c.GPS()
	}

	return gpsSum
}
