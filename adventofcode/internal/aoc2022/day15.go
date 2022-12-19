package aoc2022

import (
	"bufio"
	"errors"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

const Part1TargetY = 2000000

type Sensor struct {
	Loc   aoc.Point
	Range int
}

func (s Sensor) InRange(p aoc.Point) bool {
	return aoc.Manhattan(s.Loc, p) <= s.Range
}

func (s Sensor) ScannedPositionsAt(y int) (int, int, error) {
	d := aoc.Abs(y - s.Loc.Y)
	if d > s.Range {
		return 0, 0, errors.New("no positions scanned")
	}
	rangeLeft := s.Range - d
	return s.Loc.X - rangeLeft, s.Loc.X + rangeLeft, nil

}

func InRange(sensors []Sensor, p aoc.Point) bool {
	for _, s := range sensors {
		if s.InRange(p) {
			return true
		}
	}
	return false
}

type Beacon aoc.Point

func FindNonBeaconPositions(sensors []Sensor, beacons []Beacon, targetLine int) int {
	xs := aoc.NewIntSet([]int{})
	for _, s := range sensors {
		minX, maxX, err := s.ScannedPositionsAt(targetLine)
		if err != nil {
			continue
		}

		xs.Adds(aoc.RangeIncl(minX, maxX))
	}

	bs := aoc.IntSet{}
	for _, b := range beacons {
		if b.Y == targetLine {
			bs.Add(b.X)
		}
	}

	return len(xs) - len(bs)
}

func ParseSensorsAndBeacons(reader *bufio.Scanner, sensors *[]Sensor, beacons *[]Beacon) {
	for reader.Scan() {
		line := reader.Text()
		sb := strings.Split(line, ": ")
		sp := aoc.ParsePoint(strings.Split(sb[0], "at ")[1])
		bp := aoc.ParsePoint(strings.Split(sb[1], "at ")[1])
		r := aoc.Manhattan(sp, bp)
		*sensors = append(*sensors, Sensor{Loc: sp, Range: r})
		*beacons = append(*beacons, Beacon(bp))
	}
}

func FindDistressBeacon(sensors []Sensor, min, max aoc.Point) aoc.Point {
	d := 1
	for {
		for _, s := range sensors {
			for _, p := range s.Loc.ManhattanPerimeter(s.Range + d) {
				if p.X < min.X || p.X > max.X || p.Y < min.Y || p.Y > max.Y {
					continue
				}

				if !InRange(sensors, p) {
					return p
				}
			}
		}
		d++
	}
}

func Day15(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	sensors := []Sensor{}
	beacons := []Beacon{}

	ParseSensorsAndBeacons(reader, &sensors, &beacons)

	part1 = FindNonBeaconPositions(sensors, beacons, Part1TargetY)
	distress := FindDistressBeacon(sensors, aoc.Origin, aoc.NewPoint(4000000, 4000000))
	part2 = 4000000*distress.X + distress.Y
	return strconv.Itoa(part1), strconv.Itoa(part2)

}
