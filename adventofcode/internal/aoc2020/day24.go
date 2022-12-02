package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var empty = struct{}{}

var hexMoves = map[string]func(aoc.Point3D) aoc.Point3D{
	"e":  func(p aoc.Point3D) aoc.Point3D { return aoc.Point3D{X: p.X + 1, Y: p.Y - 1, Z: p.Z} },
	"w":  func(p aoc.Point3D) aoc.Point3D { return aoc.Point3D{X: p.X - 1, Y: p.Y + 1, Z: p.Z} },
	"se": func(p aoc.Point3D) aoc.Point3D { return aoc.Point3D{X: p.X, Y: p.Y - 1, Z: p.Z + 1} },
	"ne": func(p aoc.Point3D) aoc.Point3D { return aoc.Point3D{X: p.X + 1, Y: p.Y, Z: p.Z - 1} },
	"nw": func(p aoc.Point3D) aoc.Point3D { return aoc.Point3D{X: p.X, Y: p.Y + 1, Z: p.Z - 1} },
	"sw": func(p aoc.Point3D) aoc.Point3D { return aoc.Point3D{X: p.X - 1, Y: p.Y, Z: p.Z + 1} },
}

func hexNeighbours(p aoc.Point3D) []aoc.Point3D {
	return []aoc.Point3D{
		{X: p.X + 1, Y: p.Y - 1, Z: p.Z},
		{X: p.X - 1, Y: p.Y + 1, Z: p.Z},
		{X: p.X, Y: p.Y - 1, Z: p.Z + 1},
		{X: p.X + 1, Y: p.Y, Z: p.Z - 1},
		{X: p.X, Y: p.Y + 1, Z: p.Z - 1},
		{X: p.X - 1, Y: p.Y, Z: p.Z + 1},
	}
}

//nolint:funlen
func Day24(reader *bufio.Scanner) (string, string) {
	lines := aoc.ReadStringArray(reader)

	grid := map[aoc.Point3D]bool{}
	for _, l := range lines {
		x := 0
		p := aoc.ZeroPoint3D()
		for x < len(l) {
			switch rune(l[x]) {
			case 'e':
				p = hexMoves["e"](p)
				x++
			case 'w':
				p = hexMoves["w"](p)
				x++
			case 'n':
				p = hexMoves[l[x:x+2]](p)
				x += 2
			case 's':
				p = hexMoves[l[x:x+2]](p)
				x += 2
			}
		}

		if _, ok := grid[p]; ok {
			delete(grid, p)
		} else {
			grid[p] = true
		}
	}

	day := map[aoc.Point3D]bool{}
	for k, v := range grid {
		day[k] = v
	}

	for i := 0; i < 100; i++ {
		nextDay := map[aoc.Point3D]bool{}
		toCheck := map[aoc.Point3D]struct{}{}
		// We need to check all points (and their neighbours)
		for k := range day {
			toCheck[k] = empty
			for _, n := range hexNeighbours(k) {
				toCheck[n] = empty
			}
		}

		for p := range toCheck {
			countBlack := 0
			for _, n := range hexNeighbours(p) {
				if day[n] {
					countBlack++
				}
			}

			if day[p] {
				// Flip to white if 0 or more than 2 black neighbours
				nextDay[p] = countBlack > 0 && countBlack <= 2
			} else {
				nextDay[p] = countBlack == 2
			}

			if !nextDay[p] {
				delete(nextDay, p)
			}
		}

		day = nextDay
	}

	return strconv.Itoa(len(grid)), strconv.Itoa(len(day))
}
