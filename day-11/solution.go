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
	for range times {
		sl.Blink()
		// fmt.Printf("Blink #%v done\n", time)
	}
}

// Blink `times` times for `num` (without newing up Stone structs), return diff of number of stones
// TODO add memoization

type CacheKey struct {
	number int
	times  int
}

var cache = make(map[CacheKey]int)

func BlinkTimes(num int, times int) int {
	if times == 0 {
		return 0
	}

	// hacky memoization - think it's borked tbh
	cacheKey := CacheKey{number: num, times: times}
	v, ok := cache[cacheKey]
	if ok {
		return v
	}

	stoneCountDiff := 0

	numSlice := strings.Split(strconv.Itoa(num), "")
	lenS := len(numSlice)

	for i := range times {
		// got to add 2 not 1, since it would start blinking on _next_ blink, not current.
		blinksForNewStones := times - (i + 2)

		var result int
		cacheKey = CacheKey{num, blinksForNewStones}

		if num == 0 {
			result = BlinkTimes(1, blinksForNewStones)
		} else if lenS%2 == 0 {
			stoneCountDiff += 1

			leftHalfNums := numSlice[0 : lenS/2]
			rightHalfNums := numSlice[lenS/2:]

			leftHalf, _ := strconv.Atoi(strings.Join(leftHalfNums, ""))
			rightHalf, _ := strconv.Atoi(strings.Join(rightHalfNums, ""))

			leftResult := BlinkTimes(leftHalf, blinksForNewStones)
			rightResult := BlinkTimes(rightHalf, blinksForNewStones)
			result = leftResult + rightResult

		} else {
			newCurr := num * 2024
			result = BlinkTimes(newCurr, blinksForNewStones)
		}
		cache[cacheKey] = result
		stoneCountDiff += result
	}

	return stoneCountDiff
}

// Blink `times` times, return diff of numbers of stones
func (s *Stone) BlinkTimes(times int) int {
	if times == 0 {
		return 0
	}

	stoneCountDiff := 0
	current := s.number

	numSlice := strings.Split(strconv.Itoa(current), "")
	lenS := len(numSlice)

	for i := range times {
		// got to add 2 not 1, since it would start blinking on _next_ blink, not current.
		blinksToGo := times - (i + 2)

		if current == 0 {
			newStone := Stone{number: 1}
			fmt.Printf("Adding counts for %v, %v times\n", newStone, blinksToGo)
			stoneCountDiff += newStone.BlinkTimes(blinksToGo)
		} else if lenS%2 == 0 {
			stoneCountDiff += 1

			leftHalfNums := numSlice[0 : lenS/2]
			rightHalfNums := numSlice[lenS/2:]

			leftHalf, _ := strconv.Atoi(strings.Join(leftHalfNums, ""))
			rightHalf, _ := strconv.Atoi(strings.Join(rightHalfNums, ""))

			rightStone := Stone{number: rightHalf}
			leftStone := Stone{number: leftHalf}
			fmt.Printf("Adding counts for %v, %v times\n", rightStone, blinksToGo)
			stoneCountDiff += rightStone.BlinkTimes(blinksToGo)
			fmt.Printf("Adding counts for %v, %v times\n", leftStone, blinksToGo)
			stoneCountDiff += leftStone.BlinkTimes(blinksToGo)
		} else {
			newStone := Stone{number: current * 2024}
			fmt.Printf("Adding counts for %v, %v times\n", newStone, blinksToGo)
			stoneCountDiff += newStone.BlinkTimes(blinksToGo)
		}
	}

	return stoneCountDiff
}

// Get the number of stones from blinking `times` times, without keeping LL in memory
func (sl *StoneLine) SimulateBlinkTimes(times int) int {
	count := len(sl.GetNumbers())

	stones := sl.ToSlice()
	for _, stone := range stones {
		num := stone.number
		count += BlinkTimes(num, times)
	}

	return count
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
	return stoneList.SimulateBlinkTimes(75)
}
