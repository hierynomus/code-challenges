package aoc2022

import (
	"bufio"
	"errors"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

const Part1TargetY = 2000000

func InSensorRange(loc, p aoc.Point, r int) bool {
	return aoc.Manhattan(loc, p) <= r
}

func ScannedPositionsAt(l aoc.Point, r, y int) (int, int, error) {
	d := aoc.Abs(y - l.Y)
	if d > r {
		return 0, 0, errors.New("no positions scanned")
	}
	rangeLeft := r - d
	return l.X - rangeLeft, l.X + rangeLeft, nil
}

func InRange(sensors map[aoc.Point]int, p aoc.Point) bool {
	for l, r := range sensors {
		if InSensorRange(l, p, r) {
			return true
		}
	}
	return false
}

type Beacon aoc.Point

func FindNonBeaconPositions(sensors map[aoc.Point]int, beacons []Beacon, targetLine int) int {
	xs := aoc.NewIntSet([]int{})
	for l, r := range sensors {
		minX, maxX, err := ScannedPositionsAt(l, r, targetLine)
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

func ParseSensorsAndBeacons(reader *bufio.Scanner) (map[aoc.Point]int, []Beacon) {
	sensors := map[aoc.Point]int{}
	beacons := []Beacon{}

	for reader.Scan() {
		line := reader.Text()
		sb := strings.Split(line, ": ")
		sp := aoc.ParsePoint(strings.Split(sb[0], "at ")[1])
		bp := aoc.ParsePoint(strings.Split(sb[1], "at ")[1])
		r := aoc.Manhattan(sp, bp)
		sensors[sp] = r
		beacons = append(beacons, Beacon(bp))
	}

	return sensors, beacons
}

func FindDistressBeacon(sensors map[aoc.Point]int, min, max aoc.Point) aoc.Point {
	d := 1
	for {
		for l, r := range sensors {
			for _, p := range l.ManhattanPerimeter(r + d) {
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

	sensors, beacons := ParseSensorsAndBeacons(reader)

	part1 = FindNonBeaconPositions(sensors, beacons, Part1TargetY)
	distress := FindDistressBeacon(sensors, aoc.Origin, aoc.NewPoint(4000000, 4000000))
	part2 = 4000000*distress.X + distress.Y
	return strconv.Itoa(part1), strconv.Itoa(part2)

}
