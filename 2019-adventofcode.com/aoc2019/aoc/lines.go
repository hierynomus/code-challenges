package aoc

type Line struct {
	p1, p2 Point
}

func (l Line) Intersection(o Line) Point {
	x1, y1, x2, y2 := l.p1.X, l.p1.Y, l.p2.X, l.p2.Y
	x3, y3, x4, y4 := o.p1.X, o.p1.Y, o.p2.X, o.p2.Y
	x := ((x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	y := ((x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	return Point{x, y}
}

func (l Line) Contains(p Point) bool {
	if l.p1.X == l.p2.X {
		return p.X == l.p1.X && between(l.p1.Y, p.Y, l.p2.Y)
	}
	if l.p1.Y == l.p2.Y {
		return p.Y == l.p1.Y && between(l.p1.X, p.X, l.p2.X)
	}
	return false // TODO deal with slanted lines
}

// Check whether 'y' is between 'x' and 'z'
func between(x, y, z int) bool {
	if x < z {
		return x <= y && y <= z
	} else {
		return z <= y && y <= x
	}
}
