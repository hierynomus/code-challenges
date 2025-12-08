package aoc2025

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2025D08Sample = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

func TestDay08_Sample(t *testing.T) {
	d := day.TestDay(t, ConnectJunctionBoxes(10))
	d.WithInput(A2025D08Sample, "40", "25272")
}
