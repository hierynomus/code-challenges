package aoc2025

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day04(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadRuneGrid(reader)
	rolls := aoc.GridToPointSet(grid, func(c rune) bool { return c == '@' })

	removable := Removable(&rolls)
	part1 = len(removable)
	part2 = part1
	for len(removable) > 0 {
		rolls.Deletes(removable)
		removable = Removable(&rolls)
		part2 += len(removable)
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func Removable(rolls *aoc.PointSet) []aoc.Point {
	removable := []aoc.Point{}
	for p := range *rolls {
		cnt := 0
		for _, n := range p.Neighbours8() {
			if rolls.Contains(n) {
				cnt++
			}
		}
		if cnt < 4 {
			removable = append(removable, p)
		}
	}
	return removable
}
