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

type Stone struct {
	prev   *Stone
	next   *Stone
	number int
}

type StoneLine struct {
	head *Stone
}

func (sl *StoneLine) ToSlice() []Stone {
	current := sl.head
	stoneSlice := []Stone{*current}

	for current.next != nil {
		current = current.next
		stoneSlice = append(stoneSlice, *current)
	}

	return stoneSlice
}

func (sl *StoneLine) GetNumbers() []int {
	current := sl.head
	resultSlice := []int{current.number}

	for current.next != nil {
		current = current.next
		resultSlice = append(resultSlice, current.number)
	}

	return resultSlice
}

func (sl *StoneLine) Blink() {
	current := sl.head

	for current != nil {
		numSlice := strings.Split(strconv.Itoa(current.number), "")
		lenS := len(numSlice)

		if current.number == 0 {
			current.number = 1

			current = current.next
		} else if lenS%2 == 0 {
			// re-using current as left stone
			leftStone := current

			leftHalfNums := numSlice[0 : lenS/2]
			rightHalfNums := numSlice[lenS/2:]

			leftHalf, _ := strconv.Atoi(strings.Join(leftHalfNums, ""))
			rightHalf, _ := strconv.Atoi(strings.Join(rightHalfNums, ""))

			rightStone := Stone{number: rightHalf}
			rightStone.prev = leftStone
			rightStone.next = leftStone.next

			leftStone.number = leftHalf
			leftStone.next = &rightStone

			current = rightStone.next
		} else {
			current.number = current.number * 2024

			current = current.next
		}
	}
}

// Helper method to blink X times
func (sl *StoneLine) BlinkTimes(times int) {
	for time := range times {
		sl.Blink()
		fmt.Printf("Blink #%v done\n", time)
	}
}

// build LinkedList of stones, returning head
func BuildStones(ints []int) *StoneLine {
	var head *Stone
	var prev *Stone
	for i, n := range ints {
		stone := &Stone{
			number: n,
		}
		if i == 0 {
			head = stone
		} else {
			stone.prev = prev
			prev.next = stone
		}
		prev = stone
	}

	return &StoneLine{head: head}
}

type Input struct {
	stones []int
}

func GetInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrInputFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	var stones []int
	for _, c := range strings.Split(line, " ") {
		i, _ := strconv.Atoi(c)
		stones = append(stones, i)
	}

	return &Input{stones: stones}, nil
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

	p2Result := Part2(input)
	fmt.Printf("Part 1: got %v\n", p2Result)
}

func Part1(input *Input) int {
	stoneList := BuildStones(input.stones)

	stoneList.BlinkTimes(25)

	return len(stoneList.ToSlice())
}

func Part2(input *Input) int {
	stoneList := BuildStones(input.stones)

	stoneList.BlinkTimes(75)

	return len(stoneList.ToSlice())
}
