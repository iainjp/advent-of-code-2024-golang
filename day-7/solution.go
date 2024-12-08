package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var ErrInputFile = errors.New("cannot open input file")

type Equation struct {
	targetTotal uint64
	operands    []uint64
}

type Input struct {
	equations []Equation
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var equations []Equation
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ":")
		targetTotal, _ := strconv.Atoi(split[0])

		var operands []uint64
		for _, o := range strings.Split(strings.TrimSpace(split[1]), " ") {
			oint, _ := strconv.Atoi(o)
			operands = append(operands, uint64(oint))
		}

		equations = append(equations, Equation{uint64(targetTotal), operands})
	}

	return &Input{equations: equations}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)
}

type Node struct {
	value uint64
	mult  *Node
	add   *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(val uint64) *Tree {
	if t.root == nil {
		t.root = &Node{val, nil, nil}
	} else {
		t.root.InsertRecursive(val, t.root.value)
	}
	return t
}

func (t *Tree) GetLeafNodes() []*Node {

	var leafNodes []*Node
	var appendLeaf = func(n *Node) {
		leafNodes = append(leafNodes, n)
	}

	t.root.GetLeafNodes(appendLeaf)

	return leafNodes
}

func (t *Tree) GetLeadNodeValues() []uint64 {
	var leafNodeValues []uint64
	var appendValue = func(n *Node) {
		leafNodeValues = append(leafNodeValues, n.value)
	}

	t.root.GetLeafNodes(appendValue)
	return leafNodeValues
}

func (n *Node) GetLeafNodes(collector func(*Node)) {
	if n.mult == nil && n.add == nil {
		collector(n)
	} else {
		n.mult.GetLeafNodes(collector)
		n.add.GetLeafNodes(collector)
	}
}

func (n *Node) InsertRecursive(val uint64, prev uint64) {
	if n == nil {
		return
	}

	if n.mult == nil {
		n.mult = &Node{prev * val, nil, nil}
	} else {
		n.mult.InsertRecursive(val, n.mult.value)
	}

	if n.add == nil {
		n.add = &Node{prev + val, nil, nil}
	} else {
		n.add.InsertRecursive(val, n.add.value)
	}
}

func Part1(input *Input) uint64 {
	// Run through all possible equations, find those that are possible, sum those results.
	// Using a binary tree (since there's 2 options at each operand, + | *)
	// then each child node is the result.
	// Walk the equation and the tree, creating nodes for results. Then on to the next layer and continue.
	// Finally, walk the tree and check if any leaf nodes equal the target total.

	sumOfSolvableEquations := uint64(0)

	for _, eq := range input.equations {
		if CanBeSolved(eq) {
			sumOfSolvableEquations += eq.targetTotal
		}
	}

	return sumOfSolvableEquations
}

func CanBeSolved(eq Equation) bool {
	tree := Tree{}
	for _, n := range eq.operands {
		tree.Insert(n)
	}

	possibleResults := tree.GetLeadNodeValues()

	return slices.Contains(possibleResults, eq.targetTotal)
}
