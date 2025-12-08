package aoc2025

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day07(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadRuneGrid(reader)
	beams := aoc.Set[int]{}
	beams.Add(strings.Index(string(grid[0]), "S"))
	splits := 0
	for y := 1; y < len(grid); y++ {
		newBeams := aoc.Set[int]{}
		for beam := range beams {
			switch grid[y][beam] {
			case '^':
				splits++
				newBeams.Add(beam - 1)
				newBeams.Add(beam + 1)
			case '.':
				newBeams.Add(beam)
			}
		}
		beams = newBeams
	}

	part1 = splits
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
