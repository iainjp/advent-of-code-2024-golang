package main

import (
	"bufio"
	"errors"
	"fmt"
	"maps"
	"os"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

type Coord struct {
	x, y int
}

type Plot struct {
	crop string
	// plots adjacent of same crop type
	up    *Plot
	right *Plot
	down  *Plot
	left  *Plot
}

func (p *Plot) Crop() string {
	return p.crop
}

type PlotMap struct {
	plotByCoord map[Coord]*Plot
}

func (pm *PlotMap) Put(coord Coord, point *Plot) {
	pm.plotByCoord[coord] = point
}

// if this is called a lot, think about processing upfront
func (pm *PlotMap) GetPlots(crop string) []*Plot {
	var plots []*Plot
	for v := range maps.Values(pm.plotByCoord) {
		if v.crop == crop {
			plots = append(plots, v)
		}
	}
	return plots
}

func (pm *PlotMap) GetCoords(crop string) []Coord {
	var cs []Coord
	for k, v := range pm.plotByCoord {
		if v.crop == crop {
			cs = append(cs, k)
		}
	}
	return cs

}

func (pm *PlotMap) Get(coord Coord) *Plot {
	plot := pm.plotByCoord[coord]
	return plot
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
		plotByCoord: make(map[Coord]*Plot),
	}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		for x, c := range split {
			coord := Coord{x, y}
			currentPlot := Plot{crop: c}
			if x > 0 {
				leftCoord := Coord{x - 1, y}
				left := plotMap.Get(leftCoord)
				if left.crop == currentPlot.crop {
					currentPlot.left = left
					left.right = &currentPlot
				}
			}
			if y > 0 {
				upCoord := Coord{x, y - 1}
				up := plotMap.Get(upCoord)
				if up.crop == currentPlot.crop {
					currentPlot.up = up
					up.down = &currentPlot
				}

			}
			plotMap.Put(coord, &currentPlot)
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
	return len(input.plotMap.plotByCoord)
}
