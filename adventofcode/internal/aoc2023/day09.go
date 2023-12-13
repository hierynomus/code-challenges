package aoc2023

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day09(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)
	for _, line := range lines {
		nrs := aoc.AsIntArraySpace(line)
		next := predictNext(nrs)
		prev := predictPrevious(nrs)
		part1 += next
		part2 += prev
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func predictNext(nrs []int) int {
	reduced := make([]int, len(nrs)-1)
	allzero := true
	for i := 0; i < len(nrs)-1; i++ {
		reduced[i] = nrs[i+1] - nrs[i]
		if reduced[i] != 0 {
			allzero = false
		}
	}

	if allzero {
		return nrs[len(nrs)-1]
	}

	return nrs[len(nrs)-1] + predictNext(reduced)
}

func predictPrevious(nrs []int) int {
	reduced := make([]int, len(nrs)-1)
	allzero := true
	for i := 0; i < len(nrs)-1; i++ {
		reduced[i] = nrs[i+1] - nrs[i]
		if reduced[i] != 0 {
			allzero = false
		}
	}

	if allzero {
		return nrs[0]
	}

	return nrs[0] - predictPrevious(reduced)
}
