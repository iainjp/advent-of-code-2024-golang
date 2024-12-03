package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"iain.fyi/aoc2024/utils"
)

func main() {
	part1()
	part2()
}

func part1() {
	input, _ := GetInput("input.txt")
	result := ReportSafetyCheck(input.reports)

	trueCount := utils.CountOccurences(result)[true]

	// should be 369 for full input
	fmt.Printf("Part 1 result: %v\n", trueCount)
}

// I went back and did it for real!
func part2() {
	input, _ := GetInput("input.txt")
	result := ReportSafetyCheckWithTolerance(input.reports)

	trueCount := utils.CountOccurences(result)[true]

	// should be 428 for full input
	fmt.Printf("Part 2 result: %v\n", trueCount)
}

type Input struct {
	reports [][]int
}

var ErrInputFile = errors.New("cannot open input file")

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		var levels []int
		line := scanner.Text()
		items := strings.Split(line, " ")

		for _, itemString := range items {
			item, _ := strconv.Atoi(itemString)
			levels = append(levels, item)
		}

		reports = append(reports, levels)
	}

	return &Input{reports: reports}, nil

}

func ReportSafetyCheckWithTolerance(reports [][]int) []bool {
	var results []bool

	for _, levels := range reports {
		result := isSafe(levels)
		if !result {
			permuts := GetPermutationsWithOneMissing(levels)
			for _, permut := range permuts {
				if isSafe(permut) {
					result = true
					break
				}
			}
		}

		results = append(results, result)
	}

	return results
}

// return []bool of report safety evaluations
func ReportSafetyCheck(reports [][]int) []bool {
	var results []bool

	for _, levels := range reports {
		result := isSafe(levels)
		results = append(results, result)
	}

	return results
}

var atLeast1 = func(i int) bool { return i >= 1 }
var atMost3 = func(i int) bool { return i <= 3 }

func isSafe(levels []int) bool {
	distances := GetDistances(levels)
	return IsIncreasingOrDecreasing(levels) &&
		utils.All(distances, atLeast1) &&
		utils.All(distances, atMost3)
}

func GetPermutationsWithOneMissing(levels []int) [][]int {
	var permuts [][]int
	for index, _ := range levels {
		slice := make([]int, 0)
		slice = append(slice, levels[:index]...)
		slice = append(slice, levels[index+1:]...)
		permuts = append(permuts, slice)
	}
	return permuts
}

func GetDistances(series []int) []int {
	var distances []int
	for i := range len(series) - 1 {
		distance := utils.Abs(series[i] - series[i+1])
		distances = append(distances, distance)
	}
	return distances
}

func IsIncreasingOrDecreasing(levels []int) bool {
	isIncreasing := slices.IsSorted(levels)
	if isIncreasing {
		return true
	}

	slices.Reverse(levels)
	isDecreasing := slices.IsSorted(levels)
	return isDecreasing
}
