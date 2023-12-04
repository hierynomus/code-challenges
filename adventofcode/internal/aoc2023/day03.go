package aoc2023

import (
	"bufio"
	"strconv"
	"unicode"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type PartNr int

func Day03(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadRuneGrid(reader)

	partNrs := map[aoc.Point]PartNr{}
	symbols := map[aoc.Point]rune{}

	nr := 0
	nrPoints := []aoc.Point{}
	for y := 0; y < len(grid); y++ {
		// Leftover point from end of line
		if nr > 0 {
			// Append to partNrs
			pn := PartNr(nr)
			for _, p := range nrPoints {
				partNrs[p] = pn
			}
			nr = 0
			nrPoints = []aoc.Point{}
		}

		for x := 0; x < len(grid[y]); x++ {
			if unicode.IsDigit(grid[y][x]) {
				nr = nr*10 + int(grid[y][x]-'0')
				nrPoints = append(nrPoints, aoc.Point{X: x, Y: y})
			} else if nr > 0 {
				// Append to partNrs
				pn := PartNr(nr)
				for _, p := range nrPoints {
					partNrs[p] = pn
				}
				nr = 0
				nrPoints = []aoc.Point{}
			}

			if !unicode.IsDigit(grid[y][x]) && grid[y][x] != '.' {
				symbols[aoc.Point{X: x, Y: y}] = grid[y][x]
			}
		}
	}

	for p, s := range symbols {
		parts := aoc.IntSet{}
		for _, n := range p.Neighbours8() {
			if p, ok := partNrs[n]; ok {
				parts.Add(int(p))
			}
		}

		for p := range parts {
			part1 += p
		}

		if s == '*' && len(parts) == 2 {
			slice := parts.AsSlice()
			part2 += slice[0] * slice[1]
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
