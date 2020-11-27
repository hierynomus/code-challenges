package aoc2019

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

const (
	Asteroid = '#'
	Space    = '.'
)

type AsteroidBelt map[*aoc.Point]map[float64]*LineOfSight

type LineOfSight struct {
	asteroid *aoc.Point
	other    *aoc.Point
	v        *aoc.Vector
}

func Day10(scanner *bufio.Scanner) (string, string) {
	asteroids := AsteroidBelt{}
	asteroids.BuildAsteroidBelt(scanner)
	mostLinesOfSight := asteroids.FindMostLinesOfSight()

	angles := []float64{}
	for a := range asteroids[mostLinesOfSight] {
		angles = append(angles, a)
	}

	sort.Float64s(angles)
	nr200 := asteroids[mostLinesOfSight][angles[199]]

	return strconv.Itoa(len(asteroids[mostLinesOfSight])), strconv.Itoa(nr200.asteroid.X*100 + nr200.asteroid.Y)
}

func (ab AsteroidBelt) BuildAsteroidBelt(scanner *bufio.Scanner) {
	x, y := 0, 0
	for scanner.Scan() {
		x = 0

		for _, r := range scanner.Text() {
			if r == Asteroid {
				a := &aoc.Point{X: x, Y: y}
				lines := CalculateLinesOfSight(a, ab)
				ab[a] = lines
			}
			x++
		}
		y++
	}
}

func (ab AsteroidBelt) FindMostLinesOfSight() *aoc.Point {
	var mostLinesOfSight *aoc.Point = nil

	nrLines := 0

	for a, lines := range ab {
		if mostLinesOfSight == nil || len(lines) > nrLines {
			mostLinesOfSight = a
			nrLines = len(lines)
		}
	}

	return mostLinesOfSight
}

func CalculateLinesOfSight(asteroid *aoc.Point, asteroids AsteroidBelt) map[float64]*LineOfSight {
	linesFromAsteroid := map[float64]*LineOfSight{}

	for a, lines := range asteroids {
		v := aoc.CreateVector(a, asteroid)
		l, ok := lines[v.Radians]

		if ok && l.v.Length > v.Length || !ok {
			lines[v.Radians] = &LineOfSight{asteroid: a, other: asteroid, v: v}
		}

		v2 := aoc.CreateVector(asteroid, a)
		l2, ok := linesFromAsteroid[v2.Radians]

		if ok && l2.v.Length > v2.Length || !ok {
			linesFromAsteroid[v2.Radians] = &LineOfSight{asteroid: a, other: asteroid, v: v2}
		}
	}

	return linesFromAsteroid
}
