package main

import (
	"bufio"
	"errors"
	"fmt"
	"iter"
	"maps"
	"os"
	"strings"

	"iain.fyi/aoc2024/utils"
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

func (pm *PlotMap) GetPlotsIter() iter.Seq[*Plot] {
	return maps.Values(pm.plotByCoord)
}

// if this is called a lot, think about processing upfront
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

type Set[T comparable] struct {
	data map[T]bool
}

func (s *Set[T]) Put(t T) {
	s.data[t] = true
}

func (s *Set[T]) PutAll(t []T) {
	for _, tt := range t {
		s.Put(tt)
	}
}

func (s *Set[T]) Exists(t T) bool {
	_, ok := s.data[t]
	return ok
}

func NewSet[T comparable]() Set[T] {
	data := make(map[T]bool)
	return Set[T]{data}
}

type Region struct {
	plots Set[*Plot]
}

func (r *Region) Area() int {
	return len(r.plots.data)
}

func (r *Region) Perimeter() int {
	pCount := 0
	for p, _ := range r.plots.data {
		adjacent := []*Plot{
			p.up, p.right, p.down, p.left,
		}
		for _, adj := range adjacent {
			if adj == nil {
				pCount += 1
			}
		}
	}
	return pCount
}

func NewRegion() Region {
	return Region{plots: NewSet[*Plot]()}
}

func GetRegions(pm PlotMap) []Region {
	seen := NewSet[*Plot]()
	it := pm.GetPlotsIter()

	var walk func(plot *Plot, region *Region)
	walk = func(plot *Plot, region *Region) {
		if !seen.Exists(plot) {
			seen.Put(plot)
			region.plots.Put(plot)

			for _, p := range []*Plot{plot.up, plot.right, plot.down, plot.left} {
				if p != nil && !seen.Exists(p) && !region.plots.Exists(p) {
					walk(p, region)
					region.plots.Put(p)
				}
			}
		}
	}

	var regions []Region
	for plot := range it {
		region := NewRegion()
		walk(plot, &region)
		// if we populated region at all, we saw a new region
		if len(region.plots.data) > 0 {
			plots := utils.IterSeqToSlice(maps.Keys(region.plots.data))
			region.plots.PutAll(plots)
			regions = append(regions, region)
		}
		region = NewRegion()
	}

	return regions
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

}

func Part1(input *Input) int {
	totalCost := 0
	regions := GetRegions(input.plotMap)

	for _, r := range regions {
		cost := r.Area() * r.Perimeter()
		totalCost += cost
	}

	return totalCost
}
