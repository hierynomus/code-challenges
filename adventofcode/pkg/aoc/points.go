package aoc

import (
	"fmt"
	"math"
)

var Origin = Point{0, 0} //nolint:gochecknoglobals

type Point struct {
	X, Y int
}

func ParsePoint(s string) Point {
	var x, y int
	fmt.Sscanf(s, "x=%d, y=%d", &x, &y)
	return Point{X: x, Y: y}
}

func ReadPoint(line string) Point {
	var x, y int
	fmt.Sscanf(line, "%d,%d", &x, &y)
	return Point{X: x, Y: y}
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) Coords() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
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

func (p Point) North() Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func (p Point) NorthEast() Point {
	return Point{X: p.X + 1, Y: p.Y + 1}
}

func (p Point) East() Point {
	return Point{X: p.X + 1, Y: p.Y}
}

func (p Point) SouthEast() Point {
	return Point{X: p.X + 1, Y: p.Y - 1}
}

func (p Point) South() Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func (p Point) SouthWest() Point {
	return Point{X: p.X - 1, Y: p.Y - 1}
}

func (p Point) West() Point {
	return Point{X: p.X - 1, Y: p.Y}
}

func (p Point) NorthWest() Point {
	return Point{X: p.X - 1, Y: p.Y + 1}
}

func (p Point) Neighbours4() []Point {
	return []Point{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
	}
}

func (p Point) Neighbours8() []Point {
	return []Point{
		{X: p.X - 1, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y + 1},
	}
}

func AsGrid(pts []Point, empty, fill rune) [][]rune {
	maxX, maxY := 0, 0
	for _, p := range pts {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	grid := make([][]rune, maxY+1)
	for i := range grid {
		grid[i] = make([]rune, maxX+1)
		for t := range grid[i] {
			grid[i][t] = empty
		}
	}

	for _, p := range pts {
		grid[p.Y][p.X] = fill
	}

	return grid
}

type Point3D struct {
	X, Y, Z int64
}

func (p Point3D) Neighbours6() []Point3D {
	return []Point3D{
		{X: p.X + 1, Y: p.Y, Z: p.Z},
		{X: p.X - 1, Y: p.Y, Z: p.Z},
		{X: p.X, Y: p.Y + 1, Z: p.Z},
		{X: p.X, Y: p.Y - 1, Z: p.Z},
		{X: p.X, Y: p.Y, Z: p.Z + 1},
		{X: p.X, Y: p.Y, Z: p.Z - 1},
	}
}

func (p Point3D) Neighbours26() []Point3D {
	pts := []Point3D{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				pts = append(pts, Point3D{p.X + int64(x), p.Y + int64(y), p.Z + int64(z)})
			}
		}
	}

	return pts
}

func (p Point3D) Add(o Point3D) Point3D {
	return Point3D{X: p.X + o.X, Y: p.Y + o.Y, Z: p.Z + o.Z}
}

func (p Point3D) AddXYZ(x, y, z int64) Point3D {
	return Point3D{X: p.X + x, Y: p.Y + y, Z: p.Z + z}
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

func (p Point) ManhattanPerimeter(dist int) []Point {
	perimeter := []Point{}
	for x := p.X - dist; x <= p.X+dist; x++ {
		left := dist - Abs(p.X-x)
		perimeter = append(perimeter, Point{X: x, Y: p.Y - left})
		perimeter = append(perimeter, Point{X: x, Y: p.Y + left})
	}

	return perimeter
}

func Neighbours4(x, y int) []Point {
	return []Point{
		{X: x - 1, Y: y},
		{X: x + 1, Y: y},
		{X: x, Y: y - 1},
		{X: x, Y: y + 1},
	}
}

func Neighbours8(x, y int) []Point {
	return []Point{
		{X: x - 1, Y: y - 1},
		{X: x, Y: y - 1},
		{X: x + 1, Y: y - 1},
		{X: x - 1, Y: y},
		{X: x + 1, Y: y},
		{X: x - 1, Y: y + 1},
		{X: x, Y: y + 1},
		{X: x + 1, Y: y + 1},
	}
}

func DeltaNeighbours8() []Point {
	return []Point{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: -1, Y: 1},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
	}
}

func FindBoundingBox3D(pts []Point3D) (min, max Point3D) {
	min = Point3D{math.MaxInt64, math.MaxInt64, math.MaxInt64}
	max = Point3D{math.MinInt64, math.MinInt64, math.MinInt64}
	for _, p := range pts {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Z < min.Z {
			min.Z = p.Z
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
		if p.Z > max.Z {
			max.Z = p.Z
		}
	}

	return
}
