package aoc

var exists = struct{}{}

type PointSet Set[Point]

func NewPointSet(ps []Point) PointSet {
	s := PointSet{}
	s.Adds(ps)
	return s
}

func (s PointSet) Copy() PointSet {
	ps := PointSet{}
	for p := range s {
		ps[p] = exists
	}

	return ps
}

func (s PointSet) Add(p Point) {
	s[p] = exists
}

func (s PointSet) Adds(ps []Point) {
	for _, p := range ps {
		s[p] = exists
	}
}

// Creates a new union set
func (s PointSet) Union(ps PointSet) PointSet {
	union := PointSet{}
	for p := range s {
		union[p] = exists
	}
	for p := range ps {
		union[p] = exists
	}

	return union
}

func (s PointSet) Contains(p Point) bool {
	_, ok := s[p]
	return ok
}

func (s PointSet) BoundingBox() (min, max Point) {
	for p := range s {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}

	return
}

func (s PointSet) AsGrid() [][]rune {
	min, max := s.BoundingBox()
	grid := make([][]rune, max.Y-min.Y+1)
	for i := range grid {
		grid[i] = make([]rune, max.X-min.X+1)
		for t := range grid[i] {
			grid[i][t] = '.'
		}
	}

	for p := range s {
		grid[p.Y-min.Y][p.X-min.X] = '#'
	}

	return grid
}

func (s PointSet) Delete(p Point) {
	delete(s, p)
}

func (s PointSet) Deletes(ps []Point) {
	for _, p := range ps {
		delete(s, p)
	}
}

func GridToPointSet[T any](grid [][]T, filter func(T) bool) PointSet {
	ps := PointSet{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if filter(grid[y][x]) {
				ps.Add(Point{X: x, Y: y})
			}
		}
	}
	return ps
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items []T) Set[T] {
	s := Set[T]{}
	s.Adds(items)
	return s
}

func (s Set[T]) Add(item T) {
	s[item] = exists
}

func (s Set[T]) Adds(items []T) {
	for _, item := range items {
		s[item] = exists
	}
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Delete(item T) {
	delete(s, item)
}

func (s Set[T]) AsSlice() []T {
	items := []T{}
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s Set[T]) Copy() Set[T] {
	n := Set[T]{}
	for k := range s {
		n[k] = exists
	}

	return n
}

func (s Set[T]) Intersect(other Set[T]) {
	for k := range s {
		if !other.Contains(k) {
			s.Delete(k)
		}
	}
}

func (s Set[T]) IntersectSlice(other []T) {
	for k := range s {
		found := false
		for _, o := range other {
			found = found || k == o
		}

		if !found {
			s.Delete(k)
		}
	}
}
