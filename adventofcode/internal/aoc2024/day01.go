package aoc2024

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day01(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadIntGrid(reader, "   ")
	l1, l2 := make([]int, len(grid)), make([]int, len(grid))
	for i, line := range grid {
		l1[i] = line[0]
		l2[i] = line[1]
	}

	part1 = d1p1(l1, l2)
	part2 = d1p2(l1, l2)
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func d1p1(l1, l2 []int) int {

	sort.Ints(l1)
	sort.Ints(l2)

	s := 0
	for i := 0; i < len(l1); i++ {
		s += aoc.Abs(l1[i] - l2[i])
	}

	return s
}

func d1p2(l1, l2 []int) int {
	h := aoc.MakeIntHistogram(l2)

	sim := 0
	for _, l := range l1 {
		sim += l * h[l]
	}

	return sim
}
