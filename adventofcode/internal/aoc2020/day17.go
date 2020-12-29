package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var active = struct{}{}

type DimensionPoint struct {
	D int
	P aoc.Point3D
}

func (p DimensionPoint) Neighbours80() []DimensionPoint {
	ns := []DimensionPoint{}
	ns = append(ns, DimensionPoint{D: p.D - 1, P: p.P}, DimensionPoint{D: p.D + 1, P: p.P})
	for _, n := range p.P.Neighbours26() {
		ns = append(ns, DimensionPoint{D: p.D - 1, P: n}, DimensionPoint{D: p.D, P: n}, DimensionPoint{D: p.D + 1, P: n})
	}

	return ns
}

func Day17(reader *bufio.Scanner) (string, string) {
	input := aoc.ReadRuneGrid(reader)

	part1 := len(threeDSpace(input))
	part2 := len(DimensionalSpace(input))
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func DimensionalSpace(grid [][]rune) map[DimensionPoint]struct{} {
	activeCubes := map[DimensionPoint]struct{}{}

	for y, row := range grid {
		for x, c := range row {
			if c == '#' {
				activeCubes[DimensionPoint{D: 0, P: aoc.Point3D{X: int64(x), Y: int64(y), Z: 0}}] = active
			}
		}
	}

	round := activeCubes
	for r := 0; r < 6; r++ {
		toScan := map[DimensionPoint]struct{}{}
		for p := range round {
			for _, n := range p.Neighbours80() {
				toScan[n] = active
			}
		}

		next := map[DimensionPoint]struct{}{}
		for x := range toScan {
			_, isActive := round[x]
			activeCount := 0
			for _, n := range x.Neighbours80() {
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

	return round
}

func threeDSpace(grid [][]rune) map[aoc.Point3D]struct{} {
	activeCubes := map[aoc.Point3D]struct{}{}

	for y, row := range grid {
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

	return round
}
