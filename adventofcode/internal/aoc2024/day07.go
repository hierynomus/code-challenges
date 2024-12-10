package aoc2024

import (
	"bufio"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day07(reader *bufio.Scanner) (string, string) {
	var part1, part2 int64

	lines := aoc.ReadStringArray(reader)

	for _, line := range lines {
		// Parse the line
		ps := strings.Split(line, " ")
		outcome := aoc.ToInt64(ps[0][:len(ps[0])-1])
		nrs := []int64{}
		for _, p := range ps[1:] {
			nrs = append(nrs, aoc.ToInt64(p))
		}

		if solveCalibration(outcome, nrs[0], nrs[1:], false) {
			part1 += outcome
		}

		if solveCalibration(outcome, nrs[0], nrs[1:], true) {
			part2 += outcome
		}
	}

	return aoc.Int64ToString(part1), aoc.Int64ToString(part2)
}

func solveCalibration(outcome, nr int64, rest []int64, concat bool) bool {
	if len(rest) == 0 {
		return outcome == nr
	}

	if solveCalibration(outcome, nr+rest[0], rest[1:], concat) {
		return true
	}

	if solveCalibration(outcome, nr*rest[0], rest[1:], concat) {
		return true
	}

	if concat {
		nn := aoc.ToInt64(aoc.Int64ToString(nr) + aoc.Int64ToString(rest[0]))
		if solveCalibration(outcome, nn, rest[1:], concat) {
			return true
		}
	}

	return false
}
