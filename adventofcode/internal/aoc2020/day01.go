package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day01(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	expense := aoc.IntSet{}

	for reader.Scan() {
		i := aoc.ToInt(reader.Text())
		if expense.Contains(2020 - i) {
			part1 = i * (2020 - i)
		}

		for j := range expense {
			if expense.Contains(2020 - i - j) {
				part2 = i * j * (2020 - i - j)
			}
		}

		expense.Add(i)
	}

	// for c := range aoc.IntCombinationsN(expense, 2) {
	// 	if aoc.Sum(c) == 2020 {
	// 		part1 = c[0] * c[1]
	// 	}
	// }

	// part2 := 0
	// for c := range aoc.IntCombinationsN(expense, 3) {
	// 	if aoc.Sum(c) == 2020 {
	// 		part2 = c[0] * c[1] * c[2]
	// 	}
	// }

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
