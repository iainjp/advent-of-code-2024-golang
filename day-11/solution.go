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
func (sl *StoneLine) BlinkTimes(times int) *StoneLine {
	for range times {
		sl.Blink()
		// fmt.Printf("Blink #%v done\n", time)
	}

	return sl
}

func (sl *StoneLine) Count() int {
	asSlice := sl.ToSlice()
	return len(asSlice)
}

func SumValues(m *map[int]int) int {
	sum := 0
	for _, v := range *m {
		sum += v
	}
	return sum
}

func BlinkTimes(nums []int, times int) int {
	numbersSeen := make(map[int]int)
	for _, n := range nums {
		numbersSeen[n] += 1
	}

	blinks := 0

	for {
		newNumbers := make(map[int]int)
		for k, v := range numbersSeen {
			delete(numbersSeen, k)

			numSlice := strings.Split(strconv.Itoa(k), "")
			lenK := len(numSlice)

			if k == 0 {
				newNumbers[1] += v
			} else if lenK%2 == 0 {
				leftHalfNums := numSlice[0 : lenK/2]
				rightHalfNums := numSlice[lenK/2:]

				leftHalf, _ := strconv.Atoi(strings.Join(leftHalfNums, ""))
				rightHalf, _ := strconv.Atoi(strings.Join(rightHalfNums, ""))
				newNumbers[leftHalf] += v
				newNumbers[rightHalf] += v
			} else {
				newNumbers[k*2024] += v
			}
		}

		for k, v := range newNumbers {
			numbersSeen[k] += v
		}

		blinks++

		if blinks >= times {
			sum := SumValues(&newNumbers)
			return sum
		}
	}

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
			// fmt.Printf("Adding counts for %v, %v times\n", newStone, blinksToGo)
			stoneCountDiff += newStone.BlinkTimes(blinksToGo)
		} else if lenS%2 == 0 {
			stoneCountDiff += 1

			leftHalfNums := numSlice[0 : lenS/2]
			rightHalfNums := numSlice[lenS/2:]

			leftHalf, _ := strconv.Atoi(strings.Join(leftHalfNums, ""))
			rightHalf, _ := strconv.Atoi(strings.Join(rightHalfNums, ""))

			rightStone := Stone{number: rightHalf}
			leftStone := Stone{number: leftHalf}
			// fmt.Printf("Adding counts for %v, %v times\n", rightStone, blinksToGo)
			stoneCountDiff += rightStone.BlinkTimes(blinksToGo)
			// fmt.Printf("Adding counts for %v, %v times\n", leftStone, blinksToGo)
			stoneCountDiff += leftStone.BlinkTimes(blinksToGo)
		} else {
			newStone := Stone{number: current * 2024}
			// fmt.Printf("Adding counts for %v, %v times\n", newStone, blinksToGo)
			stoneCountDiff += newStone.BlinkTimes(blinksToGo)
		}
	}

	return stoneCountDiff
}

// Get the number of stones from blinking `times` times, without keeping LL in memory
func (sl *StoneLine) SimulateBlinkTimes(times int) int {
	return BlinkTimes(sl.GetNumbers(), times)
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
