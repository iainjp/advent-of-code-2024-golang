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

	trueCount := countOccurences(result)[true]

	fmt.Printf("Part 1 result: %v\n", trueCount)
}

// this is wrong and I can't fix it
// I stole the solution lol
func part2() {
	input, _ := GetInput("input.txt")
	result := ReportSafetyCheckWithTolerance(input.reports)

	trueCount := countOccurences(result)[true]

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
		results = append(results, isSafeWithTolerance(levels))
	}

	return results
}

func isSafeWithTolerance(levels []int) bool {
	var increasing bool
	var toleranceUsed bool = false
	var leftIndex int = 0
	var rightIndex int = 1

	var maxRightIndex int = len(levels) - 1

	for {
		a := levels[leftIndex]
		b := levels[rightIndex]

		// figure out if going up or down, and if they are equal.
		if leftIndex == 0 {
			if a < b {
				increasing = true
			} else if b < a {
				increasing = false
			} else {
				if toleranceUsed {
					return false
				}
				toleranceUsed = true
				rightIndex++
				continue
			}
		}

		if increasing && a > b {
			if toleranceUsed {
				return false
			}
			toleranceUsed = true
			leftIndex--
			continue
		}

		if !increasing && a < b {
			if toleranceUsed {
				return false
			}
			toleranceUsed = true
			leftIndex--
			continue
		}

		diff := utils.Abs(b - a)
		// breaks da rules yo
		if diff < 1 || diff > 3 {
			if toleranceUsed {
				return false
			}
			toleranceUsed = true
			if rightIndex == maxRightIndex {
				break
			}
			rightIndex++
			continue
		} else {
			leftIndex++
			if rightIndex == maxRightIndex {
				break
			}
			rightIndex++
		}
	}
	return true
}

// return []bool of report safety evaluations
func ReportSafetyCheck(reports [][]int) []bool {
	var results []bool

	atLeast1 := func(i int) bool { return i >= 1 }
	atMost3 := func(i int) bool { return i <= 3 }

	for _, levels := range reports {
		distances := GetDistances(levels)
		result := IsIncreasingOrDecreasing(levels) &&
			utils.All(distances, atLeast1) &&
			utils.All(distances, atMost3)
		results = append(results, result)
	}

	return results
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

func countOccurences[T comparable](ts []T) map[T]int {
	counts := make(map[T]int)
	for _, num := range ts {
		counts[num] = counts[num] + 1
	}
	return counts
}
