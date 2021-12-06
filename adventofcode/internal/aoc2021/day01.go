package aoc2021

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day01(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	depths := aoc.ReadIntArray(reader)
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			part1 += 1
		}
	}

	prevSum := depths[2] + depths[1] + depths[0]
	for i := 3; i < len(depths); i++ {
		sum := depths[i-2] + depths[i-1] + depths[i]
		if sum > prevSum {
			part2 += 1
		}

		prevSum = sum
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
