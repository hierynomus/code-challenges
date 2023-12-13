package aoc2023

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day11(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)

	galaxies1 := parseGalaxies(lines, 2)
	galaxies2 := parseGalaxies(lines, 1000000)

	for g := 0; g < len(galaxies1); g++ {
		for o := g + 1; o < len(galaxies1); o++ {
			part1 += aoc.Manhattan(galaxies1[g], galaxies1[o])
			part2 += aoc.Manhattan(galaxies2[g], galaxies2[o])
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parseGalaxies(lines []string, expansion int) []aoc.Point {
	galaxies := []aoc.Point{}
	dy := 0
	for y, line := range lines {
		dx := 0
		expands := true
		for x, c := range line {
			xexpands := true
			for yy := 0; yy < len(lines); yy++ {
				if lines[yy][x] == '#' {
					xexpands = false
					break
				}
			}

			if xexpands {
				dx += expansion - 1
			}

			if c == '.' {
				continue
			} else if c == '#' {
				galaxies = append(galaxies, aoc.NewPoint(x+dx, y+dy))
				expands = false
			}
		}

		if expands {
			dy += expansion - 1
		}
	}

	return galaxies
}
