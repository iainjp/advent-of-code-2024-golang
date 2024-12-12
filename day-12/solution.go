package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"iain.fyi/aoc2024/utils"
)

var ErrInputFile = errors.New("cannot open input file")

type Coord struct {
	x, y int
}

type Plot struct {
	crop  string
	up    *Plot
	right *Plot
	down  *Plot
	left  *Plot
}

func (p *Plot) Crop() string {
	return p.crop
}

type PlotMap struct {
	m map[Coord]Plot
	r map[string][]*Coord
}

func (pm *PlotMap) Put(coord Coord, point Plot) {
	pm.m[coord] = point
	pm.r[point.Crop()] = append(pm.r[point.Crop()], &coord)
}

func (pm *PlotMap) GetCoords(crop string) []Coord {
	return utils.Map(pm.r[crop], func(c *Coord) Coord { return *c })
}

func (pm *PlotMap) Get(coord Coord) Plot {
	return pm.m[coord]
}

type Input struct {
	plotMap PlotMap
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	plotMap := PlotMap{
		m: make(map[Coord]Plot),
		r: make(map[string][]*Coord),
	}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		for x, c := range split {
			coord := Coord{x, y}
			plot := Plot{crop: c}
			if x > 0 {
				left := plotMap.Get(Coord{x - 1, y})
				plot.left = &left
			}
			if x < len(split)-1 {
				right := plotMap.Get(Coord{x + 1, y})
				plot.right = &right
			}
			if y > 0 {
				up := plotMap.Get(Coord{x, y - 1})
				plot.up = &up

				up.down = &plot
			}
			plotMap.Put(coord, plot)
		}
		y += 1
	}

	return &Input{plotMap}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func Part1(input *Input) int {
	return len(input.plotMap.m)
}
