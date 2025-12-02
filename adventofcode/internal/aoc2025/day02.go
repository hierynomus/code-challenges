package aoc2025

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day02(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines, err := aoc.AsStringArray(aoc.Read(reader))
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		se := aoc.AsIntArrayS(line, "-")
		start, end := se[0], se[1]
		for i := start; i <= end; i++ {
			if isInvalidIdPart1(i) {
				part1 += i
			}
			if isInvalidIdPart2(i) {
				part2 += i
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func isInvalidIdPart1(id int) bool {
	s := strconv.Itoa(id)
	mid := len(s) / 2
	left, right := s[:mid], s[mid:]

	return left == right
}

func isInvalidIdPart2(id int) bool {
	s := strconv.Itoa(id)
	mid := len(s) / 2
	for i := 1; i <= mid; i++ {
		prefix := s[:i]
		for j := i; j <= len(s)-i; j += i {
			if s[j:j+i] != prefix {
				break
			}
			if j+i == len(s) {
				return true
			}
		}
	}

	return false
}
