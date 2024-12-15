package main

import (
	"bufio"
	"errors"
	"fmt"
	"maps"
	"os"
	"regexp"
	"strconv"
)

var ErrInputFile = errors.New("cannot open input file")

type Position struct {
	x, y int
}

type Velocity struct {
	x, y int
}

type Robot struct {
	position *Position
	velocity *Velocity
}

// non-negative modulo
func mod(a, b int) int {
	return (a%b + b) % b
}

func (r *Robot) Tick(height int, width int) {
	newX := mod(r.position.x+r.velocity.x, width)
	newY := mod(r.position.y+r.velocity.y, height)

	r.position = &Position{x: newX, y: newY}
}

type Input struct {
	robots         []*Robot
	width          int
	height         int
	elapsedSeconds int
}

func (i *Input) PositionMap() map[Position]int {
	posMap := make(map[Position]int)
	for _, robot := range i.robots {
		p := robot.position
		existing := posMap[*p]
		posMap[*p] = existing + 1
	}
	return posMap
}

func (i *Input) Print(height int, width int) {
	var output [][]string

	posMap := i.PositionMap()

	for y := range height {
		var line []string
		for x := range width {
			v, ok := posMap[Position{x, y}]
			if ok {
				line = append(line, strconv.Itoa(v))
			} else {
				line = append(line, ".")
			}
		}
		output = append(output, line)
	}

	for _, line := range output {
		for _, c := range line {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func (i *Input) Tick(height int, width int) {
	for _, robot := range i.robots {
		robot.Tick(height, width)
	}
	i.elapsedSeconds += 1
}

func (i *Input) SafetyFactor(height int, width int) int {
	posMap := i.PositionMap()

	midY := (height - 1) / 2
	midX := (width - 1) / 2

	fmt.Printf("mid-height: %v, mid-width: %v\n", midY, midX)

	q1Count := 0
	q2Count := 0
	q3Count := 0
	q4Count := 0

	for k, count := range posMap {
		if k.x == midX || k.y == midY {
			continue
		}

		if k.y < midY {
			if k.x < midX {
				q1Count += count
			} else {
				q2Count += count
			}
		} else {
			if k.x < midX {
				q3Count += count
			} else {
				q4Count += count
			}
		}
	}

	return q1Count * q2Count * q3Count * q4Count
}

func (i *Input) AllPositionsDistinct() bool {
	posMap := i.PositionMap()

	maxCount := 0
	for v := range maps.Values(posMap) {
		if v > maxCount {
			maxCount = v
		}
	}

	return maxCount < 2
}

func getPosition(line string) *Position {
	posPattern := regexp.MustCompile(`p=(\d+),(\d+)`)
	pos := posPattern.FindStringSubmatch(line)
	posX, _ := strconv.Atoi(pos[1])
	posY, _ := strconv.Atoi(pos[2])

	return &Position{x: posX, y: posY}
}

func getVelocity(line string) *Velocity {
	vPattern := regexp.MustCompile(`v=(-?\d+),(-?\d+)`)
	v := vPattern.FindStringSubmatch(line)
	vX, _ := strconv.Atoi(v[1])
	vY, _ := strconv.Atoi(v[2])

	return &Velocity{x: vX, y: vY}
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := Input{
		robots:         make([]*Robot, 0),
		elapsedSeconds: 0,
	}

	for scanner.Scan() {
		// example line: p=0,4 v=3,-3
		line := scanner.Text()

		pos := getPosition(line)
		v := getVelocity(line)
		robot := Robot{
			position: pos,
			velocity: v,
		}
		input.robots = append(input.robots, &robot)
	}

	return &input, nil
}

func main() {
	input, _ := GetInput("input.txt")

	// set real width and height
	input.width = 101
	input.height = 103

	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

	p2Result := Part2(input)
	fmt.Printf("Part 2: got %v\n", p2Result)

}

func Part1(input *Input) int {
	height := input.height
	width := input.width

	fmt.Println("Initial state:")
	input.Print(height, width)

	for range 100 {
		input.Tick(height, width)
	}

	fmt.Printf("%v second state:\n", input.elapsedSeconds)
	input.Print(height, width)

	return input.SafetyFactor(height, width)
}

// assuming the christmas tree is written in `1`s, all positions must be distinct
// (other there would be >1s)
func Part2(input *Input) int {
	height := input.height
	width := input.width

	fmt.Println("Initial state:")
	input.Print(height, width)

	for {
		input.Tick(height, width)
		if input.AllPositionsDistinct() {
			fmt.Printf("All positions distinct at %v seconds\n\n", input.elapsedSeconds)
			input.Print(height, width)
			break
		}
	}

	return input.elapsedSeconds
}
