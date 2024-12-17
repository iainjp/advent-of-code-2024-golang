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

func (c *Coord) GPS() int {
	return (c.y * 100) + c.x
}

type GridPoint struct {
	symbol string
}

type Input struct {
	grid  map[Coord]GridPoint
	moves *[]string
}

func (i *Input) Run() {
	move := i.PopMove()
	for move != "" {
		// do the move lol
		fmt.Printf("Move: %v\n", move)

		move = i.PopMove()
	}
}

func (i *Input) PopMove() string {
	moves := *i.moves
	if len(moves) == 0 {
		return ""
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

	grid := make(map[Coord]GridPoint)

	y := 0
	scanner := bufio.NewScanner(file)
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
			grid[coord] = point
			x += 1
		}
		y += 1
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
		grid:  grid,
		moves: &moves,
	}

	return &input, nil
}

func main() {
	input, _ := GetInput("input_example.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func Part1(input *Input) int {
	return 0
}
