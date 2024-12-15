package main

import (
	"bufio"
	"errors"
	"fmt"
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
	robots []*Robot
}

func (i *Input) Print(height int, width int) {
	var output [][]string

	posMap := make(map[Position]int)
	for _, robot := range i.robots {
		p := robot.position
		existing := posMap[*p]
		posMap[*p] = existing + 1
	}

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
		robots: make([]*Robot, 0),
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
	input, _ := GetInput("input_example.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func Part1(input *Input) int {

	seconds := 100
	fmt.Println("Initial state:")
	input.Print(7, 11)

	for range seconds {
		input.Tick(7, 11)
	}

	fmt.Printf("%v second state:\n", seconds)
	input.Print(7, 11)
	return 0
}
