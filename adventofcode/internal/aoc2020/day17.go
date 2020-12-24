package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var active = struct{}{}

func Day17(reader *bufio.Scanner) (string, string) {
	input := aoc.ReadRuneGrid(reader)

	activeCubes := map[aoc.Point3D]struct{}{}

	for y, row := range input {
		for x, c := range row {
			if c == '#' {
				activeCubes[aoc.Point3D{X: int64(x), Y: int64(y), Z: 0}] = active
			}
		}
	}

	round := activeCubes
	for r := 0; r < 6; r++ {
		toScan := map[aoc.Point3D]struct{}{}
		for p := range round {
			for _, n := range p.Neighbours26() {
				toScan[n] = active
			}
		}

		next := map[aoc.Point3D]struct{}{}
		for x := range toScan {
			_, isActive := round[x]
			activeCount := 0
			for _, n := range x.Neighbours26() {
				if _, ok := round[n]; ok {
					activeCount++
					if activeCount > 3 && isActive {
						break
					}
				}
			}

			if isActive && (activeCount == 2 || activeCount == 3) {
				next[x] = active
			} else if !isActive && activeCount == 3 {
				next[x] = active
			}
		}

		round = next
	}

	part1 := len(round)
	return strconv.Itoa(part1), ""
}
