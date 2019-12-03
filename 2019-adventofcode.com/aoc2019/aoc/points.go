package aoc

import "fmt"

var Origin = Point{0, 0}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func Manhattan(p1 Point, p2 Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}

func ManhattanSort(pts []Point) func(i, j int) bool {
	return func(i, j int) bool {
		p1 := pts[i]
		p2 := pts[j]
		m1, m2 := Manhattan(Origin, p1), Manhattan(Origin, p2)
		if m1 != m2 {
			return m1 < m2
		}
		if p1.X != p2.X {
			return p1.X < p2.X
		}
		return p1.Y < p2.Y
	}
}