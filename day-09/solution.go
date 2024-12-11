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

var ErrInputFile = errors.New("cannot open input file")

type Input struct {
	diskmap Diskmap
}

type Diskmap struct {
	series []int
}

type Block interface {
	Symbol() string
}

type FileBlock struct {
	id int
}

func (f FileBlock) Symbol() string {
	return strconv.Itoa(f.id)
}

type Space struct{}

func (s Space) Symbol() string {
	return "."
}

type BlocksOnDisk struct {
	blocks []Block
}

func (d *Diskmap) ToBlocks() BlocksOnDisk {
	var blocks []Block

	id := 0
	// for each digit
	for idx, n := range d.series {
		// it goes [file, space, file, space, ...]
		if idx%2 == 0 {
			// n files, each with same ID
			for range n {
				block := FileBlock{id}
				blocks = append(blocks, block)
			}
			id += 1
		} else {
			for range n {
				block := Space{}
				blocks = append(blocks, block)
			}
		}
	}

	return BlocksOnDisk{blocks: blocks}
}

func (bod BlocksOnDisk) Print() string {
	var chars []string
	for _, b := range bod.blocks {
		chars = append(chars, b.Symbol())
	}
	return strings.Join(chars, "")
}

// return new compacted BlocksOnDisk
func (bod BlocksOnDisk) Compact() BlocksOnDisk {
	cloned := make([]Block, len(bod.blocks))
	copy(cloned, bod.blocks)

	tipIdx := 0
	tailIdx := len(cloned) - 1

	// until both indexes meet or surpass each other
	for tipIdx < tailIdx {
		tip := cloned[tipIdx]
		// skip until we find an empty space
		for tip.Symbol() != "." {
			tipIdx += 1
			tip = cloned[tipIdx]
		}

		tail := cloned[tailIdx]
		// skip until we hit a file block
		for tail.Symbol() == "." {
			tailIdx -= 1
			tail = cloned[tailIdx]
		}

		// Swap tip and tail, move pointers by 1
		Swap(cloned, tipIdx, tailIdx)
		tipIdx += 1
		tailIdx -= 1
	}

	return BlocksOnDisk{blocks: cloned}
}

func (bod BlocksOnDisk) UniqueFileIds() []int {
	var ids []int
	for _, b := range bod.blocks {
		fb, ok := b.(FileBlock)
		if ok {
			ids = append(ids, fb.id)
		}
	}
	idCounts := utils.CountOccurences(ids)
	var uniqueIds []int
	for k, _ := range idCounts {
		uniqueIds = append(uniqueIds, k)
	}
	return uniqueIds
}

type SliceRange struct {
	start, end int
}

func (sr *SliceRange) Length() int {
	return sr.end - sr.start + 1
}

// return map of ID -> slice for each ID in blocks
func (bod BlocksOnDisk) BuildIndex() map[int]SliceRange {
	var indexMap = make(map[int]SliceRange, 0)

	startIdx := 0
	endIdx := 1

	maxIdx := len(bod.blocks) - 1

	for startIdx < maxIdx {
		// find start of range of files
		start := bod.blocks[startIdx]
		for start.Symbol() == "." && startIdx < maxIdx {
			startIdx += 1
			start = bod.blocks[startIdx]
		}

		endIdx = startIdx

		// find end of range (e.g. last matching ID), or end of sequence
		end := bod.blocks[endIdx]

		for end.Symbol() == start.Symbol() {
			endIdx += 1
			if endIdx > maxIdx {
				break
			}
			end = bod.blocks[endIdx]
		}

		fb, ok := start.(FileBlock)
		if ok {
			indexMap[fb.id] = SliceRange{start: startIdx, end: endIdx}
		}

		// nudge right to restart finding
		startIdx = endIdx
		endIdx += 1
	}

	return indexMap
}

// Get space spans left-to-right
func (bod BlocksOnDisk) OrderedSpaceSpans() []SliceRange {
	var emptySpans []SliceRange
	startIdx := 0
	endIdx := 1

	maxIdx := len(bod.blocks) - 1

	for startIdx < maxIdx && endIdx < maxIdx {
		// find start of range of files
		start := bod.blocks[startIdx]
		for start.Symbol() != "." && startIdx < maxIdx {
			startIdx += 1
			start = bod.blocks[startIdx]
		}

		endIdx = startIdx

		// find end of span, or end of sequence
		end := bod.blocks[endIdx]
		for end.Symbol() == start.Symbol() && endIdx < maxIdx {
			endIdx += 1
			end = bod.blocks[endIdx]
		}

		_, ok := start.(Space)
		if ok {
			r := SliceRange{start: startIdx, end: endIdx}
			emptySpans = append(emptySpans, r)
		}

		// nudge right and loop again
		startIdx = endIdx
		endIdx += 1
	}

	return emptySpans
}

func (bod BlocksOnDisk) CompactContiguousFiles() BlocksOnDisk {
	cloned := make([]Block, len(bod.blocks))
	copy(cloned, bod.blocks)
	newBod := BlocksOnDisk{blocks: cloned}

	fileIds := newBod.UniqueFileIds()
	slices.Sort(fileIds)
	slices.Reverse(fileIds)

	rangeById := newBod.BuildIndex()

	// for each ID in descending order:
	// - get left-most span,
	// - check if it's long enough. Yes? Swap em. No, check next span.
	// - if no empty spans available, move to next ID
	for _, id := range fileIds {
		fileRange := rangeById[id]
		fileSlice := newBod.blocks[fileRange.start:fileRange.end]

		spaceRanges := newBod.OrderedSpaceSpans()
		for _, spaceRange := range spaceRanges {
			// break if spaceRange is not to left of fileRange
			if fileRange.start < spaceRange.start {
				break
			}

			spaceSlice := newBod.blocks[spaceRange.start:spaceRange.end]
			if len(spaceSlice) >= len(fileSlice) {
				for n := range len(fileSlice) {
					SwapBetweenSlices(spaceSlice, n, fileSlice, n)
				}
				break
			}
		}

		// rangeById = newBod.BuildIndex()
	}

	return newBod
}

// Swap two entries in a slice
func Swap[S ~[]E, E comparable](slice S, idx1 int, idx2 int) {
	slice[idx1], slice[idx2] = slice[idx2], slice[idx1]
}

func SwapBetweenSlices[S ~[]E, E comparable](slice1 S, idx1 int, slice2 S, idx2 int) {
	slice1[idx1], slice2[idx2] = slice2[idx2], slice1[idx1]
}

func (bod BlocksOnDisk) Checksum() int {
	checksum := 0

	for pos, block := range bod.blocks {
		// type assertion
		file, ok := block.(FileBlock)
		if ok {
			checksum += pos * file.id
		}
	}

	return checksum
}

func main() {
	input, _ := GetInput("input.txt")
	p1Result := Part1(input)
	fmt.Printf("Part 1: got %v\n", p1Result)

	p2Result := Part2(input)
	fmt.Printf("Part 2: got %v\n", p2Result)
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

	var series []int
	for _, c := range strings.Split(line, "") {
		i, _ := strconv.Atoi(c)
		series = append(series, i)
	}
	diskmap := Diskmap{series: series}

	return &Input{diskmap: diskmap}, nil
}

func Part1(input *Input) int {
	blocks := input.diskmap.ToBlocks()
	compacted := blocks.Compact()
	checksum := compacted.Checksum()

	return checksum
}

func Part2(input *Input) int {
	blocks := input.diskmap.ToBlocks()
	contiguous := blocks.CompactContiguousFiles()
	checksum := contiguous.Checksum()

	return checksum
}
