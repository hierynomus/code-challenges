package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day01(reader *bufio.Scanner) (string, string) {
	expense := []int{}
	for reader.Scan() {
		expense = append(expense, aoc.ToInt(reader.Text()))
	}

	var part1 int

	for c := range aoc.IntCombinationsN(expense, 2) {
		if aoc.Sum(c) == 2020 {
			part1 = c[0] * c[1]
		}
	}

	part2 := 0
	for c := range aoc.IntCombinationsN(expense, 3) {
		if aoc.Sum(c) == 2020 {
			part2 = c[0] * c[1] * c[2]
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
