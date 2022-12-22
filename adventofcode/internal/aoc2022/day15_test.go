package aoc2022

import (
	"bufio"
	"strings"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

const A2022D15Sample = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

func TestDay15_Sample(t *testing.T) {
	sensors, beacons := ParseSensorsAndBeacons(bufio.NewScanner(strings.NewReader(A2022D15Sample)))
	assert.Equal(t, 26, FindNonBeaconPositions(sensors, beacons, 10))
	assert.Equal(t, aoc.NewPoint(14, 11), FindDistressBeacon(sensors, aoc.Origin, aoc.NewPoint(20, 20)))
}
