package aoc2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day05(reader *bufio.Scanner) (string, string) {
	part1, part2 := 0, 0

	lines := make([]aoc.Line, 0)
	for reader.Scan() {
		line := reader.Text()
		pts := strings.Split(line, " -> ")
		lines = append(lines, aoc.NewLine(aoc.ReadPoint(pts[0]), aoc.ReadPoint(pts[1])))
	}

	intersects := make(map[aoc.Point]int)
	intersects2 := make(map[aoc.Point]int)

	for i := 0; i < len(lines); i++ {
		l1 := lines[i]
		if l1.IsHorizOrVert() {
			pts := l1.Points()
			for _, p := range pts {
				intersects[p]++
				intersects2[p]++
			}
		} else if l1.IsDiag45() {
			pts := l1.Points()
			for _, p := range pts {
				intersects2[p]++
			}
		}
	}

	for _, p := range intersects {
		if p > 1 {
			part1++
		}
	}

	for _, p := range intersects2 {
		if p > 1 {
			part2++
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
