package aoc2023

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var PipeSymbols = map[rune][]aoc.Point{
	'|': {aoc.North, aoc.South},
	'-': {aoc.East, aoc.West},
	'F': {aoc.South, aoc.East},
	'7': {aoc.West, aoc.South},
	'J': {aoc.North, aoc.West},
	'L': {aoc.East, aoc.North},
}

func Day10(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)
	var start aoc.Point
	var loop aoc.PointSet
	var loopStart aoc.Point

	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				continue
			}

			if c == 'S' {
				start = aoc.NewPoint(x, y)
				loop = aoc.NewPointSet([]aoc.Point{start})
				found := false
				for _, n := range aoc.NewPoint(x, y).Neighbours4() {
					r := rune(lines[n.Y][n.X])
					if _, ok := PipeSymbols[r]; !ok {
						continue
					}

					for _, d := range PipeSymbols[r] {
						if n.Add(d) == start {
							loopStart = n
							found = true
						}
					}

					if found {
						break
					}
				}
			}
		}
	}

	nxt := loopStart
	for {
		if loop.Contains(nxt) {
			break
		}

		loop.Add(nxt)
		for _, d := range PipeSymbols[rune(lines[nxt.Y][nxt.X])] {
			n := nxt.Add(d)
			if loop.Contains(n) {
				continue
			}

			nxt = n
			break
		}
	}

	part1 = len(loop) / 2
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
