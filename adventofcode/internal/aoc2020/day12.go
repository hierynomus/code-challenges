package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Ship struct {
	Point   aoc.Point
	Bearing rune
}

var Movements map[rune]func(p aoc.Point, d int) aoc.Point = map[rune]func(p aoc.Point, d int) aoc.Point{
	'N': func(p aoc.Point, d int) aoc.Point { return p.AddXY(0, -d) },
	'S': func(p aoc.Point, d int) aoc.Point { return p.AddXY(0, d) },
	'E': func(p aoc.Point, d int) aoc.Point { return p.AddXY(d, 0) },
	'W': func(p aoc.Point, d int) aoc.Point { return p.AddXY(-d, 0) },
}

func Day12(reader *bufio.Scanner) (string, string) {
	lines := aoc.ReadStringArray(reader)

	s := &Ship{aoc.Origin, 'E'}
	for _, l := range lines {
		d := aoc.ToInt(l[1:])
		switch rune(l[0]) {
		case 'F':
			s.Point = Movements[s.Bearing](s.Point, d)
		case 'L':
			s.Bearing = turn(s.Bearing, d, -1)
		case 'R':
			s.Bearing = turn(s.Bearing, d, 1)
		default:
			s.Point = Movements[rune(l[0])](s.Point, d)
		}
	}

	part1 := aoc.Manhattan(aoc.Origin, s.Point)

	wp := aoc.Point{X: 10, Y: -1}
	ship := &Ship{Point: aoc.Origin}
	for _, l := range lines {
		d := aoc.ToInt(l[1:])
		switch rune(l[0]) {
		case 'F':
			ship.Point = ship.Point.AddXY(d*wp.X, d*wp.Y)
		case 'L':
			for i := 0; i < d/90; i++ {
				wp = aoc.Point{X: wp.Y, Y: -wp.X}
			}
		case 'R':
			for i := 0; i < d/90; i++ {
				wp = aoc.Point{X: -wp.Y, Y: wp.X}
			}
		default:
			wp = Movements[rune(l[0])](wp, d)
		}
	}

	part2 := aoc.Manhattan(aoc.Origin, ship.Point)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func turn(currentBearing rune, degrees int, direction int) rune {
	bearings := "NESW"
	idx := strings.IndexRune(bearings, currentBearing)
	idx = (4 + idx + (direction * degrees / 90)) % 4
	return rune(bearings[idx])
}
