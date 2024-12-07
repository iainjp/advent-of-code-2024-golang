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

		var operands []int
		for _, o := range strings.Split(strings.TrimSpace(split[1]), " ") {
			oint, _ := strconv.Atoi(o)
			operands = append(operands, oint)
		}

		equations = append(equations, Equation{targetTotal, operands})
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
	add *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(val uint64) {
	t.InsertRecursive(t.root, val, 0)
}

func (t *Tree) InsertRecursive(node *Node, val uint64, prev uint64) *Node {
	if t.root == nil {
		t.root = &Node{value: val, nil, nil}
		return t.root
	}

	if node.mult == nil {
		node.mult = &Node{prev*val, nil, nil}
	}

	if node.add == nil {
		node.add = &Node{prev+add, nil, nil}
	}

	t.InsertRecursive(node.mult, val, node.value)
	t.InsertRecursive(node.add, val, node.value)
	return node
}



func Part1(input *Input) int {
	// Run through all possible equations, find those that are possible, sum those results.
	// Using a binary tree (since there's 2 options at each operand, + | *)
	// then each child node is the result.
	// Walk the equation and the tree, creating nodes for results. Then on to the next layer and continue.
	// Finally, walk the tree and check if any leaf nodes equal the target total.


	for _, equation := range input.equations {
		start := Tree{
			value: equation.operands[0],
		}

		for i := equation.operands[1:] {
			start.left = Multiple(i, int(start.value))
			start.right = Add(i, int(start.value))

		}

	}
}

func Add(i, i2 int) int {
	return i + i2
}

func Multiple(i, i2 int) int {
	return i * i2
}
