package aoc2024

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Antenna struct {
	Loc  aoc.Point
	Freq rune
}

func Day08(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadRuneGrid(reader)

	antennae := []Antenna{}
	freqMap := map[rune][]aoc.Point{}
	for y, row := range grid {
		for x, c := range row {
			if c != '.' {
				antennae = append(antennae, Antenna{Loc: aoc.Point{X: x, Y: y}, Freq: c})
				freqMap[c] = append(freqMap[c], aoc.Point{X: x, Y: y})
			}
		}
	}

	antiNodes := aoc.PointSet{}
	resonantHarmonics := aoc.PointSet{}
	for _, locs := range freqMap {
		combos := aoc.PointCombinations(locs)
		for _, combo := range combos {
			dx, dy := combo[1].X-combo[0].X, combo[1].Y-combo[0].Y
			// Every antanna is also a resonant harmonic antinode
			resonantHarmonics.Add(combo[0])
			resonantHarmonics.Add(combo[1])

			an := combo[0].AddXY(-dx, -dy)
			if aoc.InBounds(an, grid) {
				antiNodes.Add(an)
				for aoc.InBounds(an, grid) {
					resonantHarmonics.Add(an)
					an = an.AddXY(-dx, -dy)
				}
			}

			an = combo[1].AddXY(dx, dy)
			if aoc.InBounds(an, grid) {
				antiNodes.Add(an)
				for aoc.InBounds(an, grid) {
					resonantHarmonics.Add(an)
					an = an.AddXY(dx, dy)
				}
			}
		}
	}

	part1 = len(antiNodes)
	part2 = len(resonantHarmonics)
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
