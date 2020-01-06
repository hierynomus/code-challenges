package aoc

import "fmt"

var Origin = Point{0, 0} //nolint:gochecknoglobals

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) Add(o Point) Point {
	return Point{X: p.X + o.X, Y: p.Y + o.Y}
}

func (p Point) AddXY(x, y int) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}

type Point3D struct {
	X, Y, Z int64
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func ZeroPoint3D() Point3D {
	return Point3D{X: 0, Y: 0, Z: 0}
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

func Neighbours(x, y int) []Point {
	return []Point{
		Point{X: x - 1, Y: y}, //nolint:gofmt
		Point{X: x + 1, Y: y},
		Point{X: x, Y: y - 1},
		Point{X: x, Y: y + 1},
	}
}
