package days

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
)

const (
	Asteroid = '#'
	Space    = '.'
)

type AsteroidBelt map[*aoc.Point]map[float64]*LineOfSight
type Day10 struct {
	asteroids AsteroidBelt
}
type LineOfSight struct {
	asteroid *aoc.Point
	other    *aoc.Point
	v        *aoc.Vector
}

func (d *Day10) Solve(scanner *bufio.Scanner) (string, string) {
	d.asteroids = AsteroidBelt{}
	d.BuildAsteroidBelt(scanner)
	mostLinesOfSight := d.FindMostLinesOfSight()

	angles := []float64{}
	for a := range d.asteroids[mostLinesOfSight] {
		angles = append(angles, a)
	}

	sort.Float64s(angles)
	nr200 := d.asteroids[mostLinesOfSight][angles[199]]

	return strconv.Itoa(len(d.asteroids[mostLinesOfSight])), strconv.Itoa(nr200.asteroid.X*100 + nr200.asteroid.Y)
}

func (d *Day10) BuildAsteroidBelt(scanner *bufio.Scanner) {
	x, y := 0, 0
	for scanner.Scan() {
		x = 0

		for _, r := range scanner.Text() {
			if r == Asteroid {
				a := &aoc.Point{X: x, Y: y}
				lines := CalculateLinesOfSight(a, d.asteroids)
				d.asteroids[a] = lines
			}
			x++
		}
		y++
	}
}

func (d *Day10) FindMostLinesOfSight() *aoc.Point {
	var mostLinesOfSight *aoc.Point = nil

	nrLines := 0

	for a, lines := range d.asteroids {
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
