package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day09(reader *bufio.Scanner) (string, string) {
	xmas := aoc.ReadIntArray(reader)

	var part1 int
	for i := 25; i < len(xmas); i++ {
		if !validXmas(i, xmas) {
			part1 = xmas[i]
		}
	}

	start, end := findEncryptionRange(xmas, part1)
	part2 := aoc.Max(xmas[start:end+1]) + aoc.Min(xmas[start:end+1])

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func validXmas(i int, xmas []int) bool {
	x := xmas[i]
	for a := i - 25; a < i-1; a++ {
		for b := a + 1; b < i; b++ {
			if xmas[a]+xmas[b] == x {
				return true
			}
		}
	}

	return false
}

func findEncryptionRange(xmas []int, weakness int) (int, int) {
	s, e := 1, 2
	sum := xmas[s] + xmas[e]
	for true {
		if sum < weakness {
			e += 1
			sum += xmas[e]
		} else if sum > weakness {
			sum -= xmas[s]
			s += 1
		} else {
			break
		}
	}

	return s, e
}
