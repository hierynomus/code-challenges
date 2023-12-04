package aoc2023

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var Digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Day01(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)
	part1 = d1p1(lines)
	part2 = d1p2(lines)
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func d1p1(lines []string) int {
	part1 := 0
	for _, line := range lines {
		first, last := -1, -1
		for i := 0; i < len(line); i++ {
			if f, err := strconv.Atoi(string(line[i])); err == nil && first == -1 {
				first = f
			}

			if l, err := strconv.Atoi(string(line[len(line)-i-1])); err == nil && last == -1 {
				last = l
			}

			if first != -1 && last != -1 {
				break
			}
		}

		part1 += (first*10 + last)
	}

	return part1
}

func d1p2(lines []string) int {
	part2 := 0
	for _, line := range lines {
		first, last := -1, -1
		for i := 0; i < len(line); i++ {
			if f, err := strconv.Atoi(string(line[i])); err == nil && first == -1 {
				first = f
			}

			if l, err := strconv.Atoi(string(line[len(line)-i-1])); err == nil && last == -1 {
				last = l
			}

			for k, v := range Digits {
				if strings.HasPrefix(line[i:], k) && first == -1 {
					first = v
				}

				if strings.HasPrefix(line[len(line)-i-1:], k) && last == -1 {
					last = v
				}
			}

			if first != -1 && last != -1 {
				break
			}
		}

		part2 += (first*10 + last)
	}

	return part2
}
