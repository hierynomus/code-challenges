package aoc2022

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day01(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	elves := []int{}
	lines := aoc.ReadStringArray(reader)
	elf := 0
	for _, line := range lines {
		if line == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		elf += aoc.ToInt(line)
	}

	elves = append(elves, elf)

	sort.Ints(elves)
	part1 = elves[len(elves)-1]
	part2 = elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
