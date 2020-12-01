package aoc2018

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day01(reader *bufio.Scanner) (string, string) {
	var input []int

	for reader.Scan() {
		input = append(input, aoc.ToInt(reader.Text()))
	}

	part1 := 0
	for _, i := range input {
		part1 += i
	}

	seen := aoc.IntSet{}
	part2 := 0
	sum := 0
	for part2 == 0 {
		for _, i := range input {
			sum += i
			if !seen.Contains(sum) {
				seen.Add(sum)
			} else {
				part2 = sum
				break
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
