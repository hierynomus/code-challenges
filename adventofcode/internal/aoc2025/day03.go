package aoc2025

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day03(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)

	for _, line := range lines {
		ints := aoc.StringToIntSlice(line)
		a, aidx := aoc.MaxIdx(ints, 0, len(ints)-2)
		b, _ := aoc.MaxIdx(ints, aidx+1, len(ints)-1)

		jolt := make([]int, 12)
		idx := 0
		for i := 0; i < len(jolt); i++ {
			m, midx := aoc.MaxIdx(ints, idx, len(ints)-(len(jolt)-i))
			jolt[i] = m
			idx = midx + 1
		}

		part1 += a*10 + b
		part2 += aoc.ToInt(aoc.IntArrayAsString(jolt, ""))
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
