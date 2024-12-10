package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

type Node struct {
	height int
	next   map[int]*Node
}

type Graph struct {
	trailheads []*Node
}

type Input struct {
	graph Graph
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
	var trailheads []*Node

	var nodeMatrix [][]*Node

	for scanner.Scan() {
		line := scanner.Text()
		var nodeLine []*Node
		for _, c := range strings.Split(line, "") {
			h, _ := strconv.Atoi(c)
			n := &Node{height: h}
			if n.height == 0 {
				trailheads = append(trailheads, n)
			}
			nodeLine = append(nodeLine, n)
		}
		nodeMatrix = append(nodeMatrix, nodeLine)
	}

	// TODO other loop to set next field values

	return &Input{graph: Graph{trailheads: trailheads}}, nil
}

func Part1(input *Input) int {
	return len(input.graph.trailheads)
}
