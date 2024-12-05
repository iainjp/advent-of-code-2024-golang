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

type OrderingRule struct {
	lower, upper int
}

func (r OrderingRule) Evaluate(pages []int) bool {
	// rule only matters if both sides are in update
	shouldEval := slices.Contains(pages, r.lower) &&
		slices.Contains(pages, r.upper)

	// if it's not applicable, same outcome as if it was and valid
	if !shouldEval {
		return true
	}

	return slices.Index(pages, r.lower) < slices.Index(pages, r.upper)
}

type Update struct {
	pages []int
}

func (u Update) MiddlePageValue() int {
	middlePageIndex := (len(u.pages) - 1) / 2
	return u.pages[middlePageIndex]
}

type Input struct {
	rules   []OrderingRule
	updates []Update
}

const (
	RULES_SEPARATOR = "|"
	PAGE_SEPARATOR  = ","
)

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules []OrderingRule
	var updates []Update

	for scanner.Scan() {
		line := scanner.Text()
		// skip empty line
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if strings.Contains(line, RULES_SEPARATOR) {
			parts := strings.Split(line, RULES_SEPARATOR)
			lower, _ := strconv.Atoi(parts[0])
			upper, _ := strconv.Atoi(parts[1])
			rule := OrderingRule{lower: lower, upper: upper}
			rules = append(rules, rule)
		}

		if strings.Contains(line, PAGE_SEPARATOR) {
			pageList := strings.Split(line, PAGE_SEPARATOR)
			var pages []int
			for _, p := range pageList {
				pint, _ := strconv.Atoi(p)
				pages = append(pages, pint)
			}
			update := Update{pages: pages}
			updates = append(updates, update)
		}
	}

	return &Input{rules: rules, updates: updates}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)
}

func Part1(input *Input) int {
	// TODO
	// 1. parse rules
	// 2. parse pages
	// 3. run pages through rules. if all match, get middle value
	// 4. return sum of middle values of page sets that match

	middlePageTotal := 0

	for _, update := range input.updates {
		var ruleOutcomes []bool
		for _, rule := range input.rules {
			ruleOutcomes = append(ruleOutcomes, rule.Evaluate(update.pages))
		}

		if !slices.Contains(ruleOutcomes, false) {
			middlePageTotal += update.MiddlePageValue()
		}
	}

	return middlePageTotal
}
