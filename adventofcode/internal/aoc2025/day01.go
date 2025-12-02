package aoc2025

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day01(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)

	idx := 50
	for _, line := range lines {
		dir, num := line[0], aoc.ToInt(line[1:])
		switch dir {
		case 'R':
			part2 += num / 100
			idx = idx + num%100
			if idx > 100 {
				part2 += 1
			}
			idx = idx % 100
		case 'L':
			part2 += num / 100
			nidx := idx - num%100
			if idx != 0 && nidx < 0 {
				part2 += 1
			}
			idx = (nidx + 100) % 100
		}

		if idx == 0 {
			part1 += 1
			part2 += 1
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
