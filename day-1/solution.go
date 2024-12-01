package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput("input.txt")
	result := SumDistances(input.left, input.right)
	fmt.Printf("Part 1 result: %v\n", result)
}

func part2() {
	input := getInput("input.txt")
	result := SimilarityScore(input.left, input.right)
	fmt.Printf("Part 2 result: %v\n", result)
}

type Input struct {
	left, right []int
}

func getInput(filename string) Input {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "   ")

		leftInt, _ := strconv.Atoi(items[0])
		left = append(left, leftInt)

		rightInt, _ := strconv.Atoi(items[1])
		right = append(right, rightInt)
	}

	return Input{left: left, right: right}

}

func abs(v int) int {
	return max(v, -v)
}

func SumDistances(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for index, leftNum := range left {
		distance := abs(leftNum - right[index])
		total += distance
	}

	return total
}

func SimilarityScore(left, right []int) int {
	counts := make(map[int]int)
	for _, num := range right {
		counts[num] = counts[num] + 1
	}

	score := 0
	for _, num := range left {
		score += num * counts[num]
	}
	return score
}
