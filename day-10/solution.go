package main

import (
	"bufio"
	"errors"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"

	"iain.fyi/aoc2024/utils"
)

var ErrInputFile = errors.New("cannot open input file")

type Node struct {
	height int
	next   []*Node
}

func (n *Node) IsValidNext(on *Node) bool {
	return n.height-on.height == 1
}

type Graph struct {
	trailheads []*Node
}

// Walk from each trailhead, return count of finishes per trailhead
func (g *Graph) Walk() map[*Node]int {

	// collection of trailhead, X occurrences = X finishes
	var trailheadFinished []*Node
	collect := func(n *Node) {
		trailheadFinished = append(trailheadFinished, n)
	}

	var dfs func(head *Node, n *Node, collector func(n *Node))
	dfs = func(head *Node, n *Node, collector func(n *Node)) {
		if n.height == 9 {
			collector(head)
		}
		for _, nn := range n.next {
			dfs(head, nn, collector)
		}
	}

	// run DFS from each trailhead
	for _, th := range g.trailheads {
		dfs(th, th, collect)
	}

	return utils.CountOccurences(trailheadFinished)
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

	// build initial nodeMatrix
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

	// set up neighbours
	for outerIdx, nodeLine := range nodeMatrix {
		for innerIdx, currNode := range nodeLine {
			// left
			if innerIdx > 0 {
				left := nodeLine[innerIdx-1]
				if left.IsValidNext(currNode) {
					currNode.next = append(currNode.next, left)
				}
			}

			// right
			if innerIdx < len(nodeLine)-1 {
				right := nodeLine[innerIdx+1]
				if right.IsValidNext(currNode) {
					currNode.next = append(currNode.next, right)
				}
			}

			// up
			if outerIdx > 0 {
				up := nodeMatrix[outerIdx-1][innerIdx]
				if up.IsValidNext(currNode) {
					currNode.next = append(currNode.next, up)
				}
			}

			// down
			if outerIdx < len(nodeMatrix)-1 {
				down := nodeMatrix[outerIdx+1][innerIdx]
				if down.IsValidNext(currNode) {
					currNode.next = append(currNode.next, down)
				}
			}
		}
	}

	return &Input{graph: Graph{trailheads: trailheads}}, nil
}

func Part1(input *Input) int {
	finishes := input.graph.Walk()

	total := 0
	for v := range maps.Values(finishes) {
		total += v
	}

	return total
}
