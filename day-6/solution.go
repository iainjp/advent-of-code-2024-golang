package main

import (
	"bufio"
	"errors"
	"fmt"
	"maps"
	"os"
	"strings"

	"iain.fyi/aoc2024/utils"
)

var ErrInputFile = errors.New("cannot open input file")

type Point struct {
	x, y int
}

type Direction int

const (
	None  Direction = iota
	Up              = iota
	Down            = iota
	Left            = iota
	Right           = iota
)

const (
	Obstacle   = "#"
	EmptySpace = "."
	OutOfMap   = ""
)

type Guard struct {
	position  Point
	direction Direction
	path      []Point
	seen      map[Point]int
}

type Map map[Point]string

func NextPoint(curr Point, dir Direction) *Point {
	var next *Point
	switch dir {
	case Up:
		next = &Point{curr.x, curr.y - 1}
	case Right:
		next = &Point{curr.x + 1, curr.y}
	case Down:
		next = &Point{curr.x, curr.y + 1}
	case Left:
		next = &Point{curr.x - 1, curr.y}
	}

	return next
}

func NextDirection(curr Direction) Direction {
	var next Direction
	switch curr {
	case Up:
		next = Right
	case Right:
		next = Down
	case Down:
		next = Left
	case Left:
		next = Up
	}
	return next
}

// Move the guard, return false if still in map, true if out
func (g *Guard) Move(pointMap Map) bool {
	current := g.position

	next := NextPoint(current, g.direction)
	for pointMap[*next] == Obstacle {
		g.direction = NextDirection(g.direction)
		next = NextPoint(current, g.direction)
	}

	newPath := append(g.path, *next)
	g.path = newPath
	g.position = *next

	return pointMap[g.position] == OutOfMap
}

// make copy of guard
func (g Guard) Clone() Guard {
	clone := Guard{
		position:  g.position,
		direction: g.direction,
	}

	clone.path = append(clone.path, g.path...)

	seen := make(map[Point]int)
	for k, v := range g.seen {
		seen[k] = v
	}
	clone.seen = seen

	return clone
}

type Result int

const (
	NotFinished Result = iota
	Finished
	InLoop
)

func (g *Guard) MoveWithLoopDetection(pointMap Map) Result {
	oldPosition := g.position

	next := NextPoint(oldPosition, g.direction)
	for pointMap[*next] == Obstacle {
		g.direction = NextDirection(g.direction)
		next = NextPoint(oldPosition, g.direction)
	}

	g.seen[*next] = g.seen[*next] + 1

	newPath := append(g.path, *next)
	g.path = newPath
	g.position = *next

	seenOld := g.seen[oldPosition]
	seenNext := g.seen[*next]

	if seenOld > 3 && seenNext > 3 {
		return InLoop
	}

	if pointMap[g.position] == OutOfMap {
		return Finished
	}

	return NotFinished
}

type Input struct {
	pointMap Map
	guard    Guard
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var guard Guard
	pointMap := make(Map, 100)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")

		for x, c := range chars {
			point := Point{x, y}
			pointMap[point] = c

			possibleGuard := MakeGuard(c, point)
			if possibleGuard != nil {
				guard = *possibleGuard
			}
		}
		y++
	}

	return &Input{pointMap, guard}, nil
}

func MakeGuard(c string, p Point) *Guard {
	directions := map[string]Direction{
		"^": Up,
		">": Right,
		"<": Left,
		"v": Down,
	}

	if directions[c] == None {
		return nil
	}

	return &Guard{
		position:  p,
		direction: directions[c],
		path:      []Point{p},
		seen:      map[Point]int{p: 1},
	}
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

	p2Result := Part2(input)
	fmt.Printf("Part 2: got %v\n", p2Result)

}

func Part1(input *Input) int {
	// 1. Parse map and place guard
	// 2. Walk guard using rules, until you hit boundary
	// 3. Record the steps in full
	// 4. Count steps, donezo

	pointMap := input.pointMap
	guard := input.guard

	outOfBounds := guard.Move(pointMap)
	for !outOfBounds {
		outOfBounds = guard.Move(pointMap)
	}

	counts := utils.CountOccurences(guard.path)
	uniquePoints := 0
	for range maps.Keys(counts) {
		uniquePoints++
	}

	// -1 to not count the out-of-bounds point
	return uniquePoints - 1
}

func Part2(input *Input) int {
	// TODO
	// 1. Parse map, place guard
	// 2. Create permutation of map, each with one empty space replaced with an obstacle
	// 3. Run the guard while identifying loops (lots of duplicates in guard.path?)
	// 4. Track counter for permutations that can cause loops

	pointMap := input.pointMap
	firstRunGuard := input.guard.Clone()

	// run guard once to get path
	firstRunFinished := firstRunGuard.Move(pointMap)
	for !firstRunFinished {
		firstRunFinished = firstRunGuard.Move(pointMap)
	}

	allOptions := AllMapOptions(pointMap, firstRunGuard)

	var resultSet []Result
	for _, m := range allOptions {
		guard := input.guard.Clone()

		result := guard.MoveWithLoopDetection(m)
		for result == NotFinished {
			result = guard.MoveWithLoopDetection(m)
		}
		resultSet = append(resultSet, result)
	}

	loops := 0
	for _, r := range resultSet {
		if r == InLoop {
			loops += 1
		}
	}

	return loops
}

func AllMapOptions(m Map, guard Guard) []Map {
	var allOptions []Map

	var uniqueEmptySpaceInMap []Point
	pointCounts := utils.CountOccurences(guard.path)
	for p := range maps.Keys(pointCounts) {
		if m[p] == "." {
			uniqueEmptySpaceInMap = append(uniqueEmptySpaceInMap, p)
		}
	}

	for _, p := range uniqueEmptySpaceInMap {
		mNew := make(Map, len(m))
		maps.Copy(mNew, m)
		mNew[p] = Obstacle
		allOptions = append(allOptions, mNew)
	}
	return allOptions
}
