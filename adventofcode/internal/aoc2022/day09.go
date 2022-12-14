package aoc2022

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Rope struct {
	TailVisited aoc.PointSet
	Knots       []aoc.Point
}

func NewRope(length int) *Rope {
	r := &Rope{
		Knots:       []aoc.Point{},
		TailVisited: aoc.NewPointSet([]aoc.Point{aoc.Origin}),
	}

	for i := 0; i < length; i++ {
		r.Knots = append(r.Knots, aoc.Origin)
	}

	return r
}

func (r *Rope) MoveIfNeeded(idx int) {
	head := idx - 1
	switch {
	case r.Knots[head].X == r.Knots[idx].X && aoc.Abs(r.Knots[head].Y-r.Knots[idx].Y) > 1:
		r.Knots[idx] = r.Knots[idx].AddXY(0, aoc.Sign(r.Knots[head].Y-r.Knots[idx].Y))
	case r.Knots[head].Y == r.Knots[idx].Y && aoc.Abs(r.Knots[head].X-r.Knots[idx].X) > 1:
		r.Knots[idx] = r.Knots[idx].AddXY(aoc.Sign(r.Knots[head].X-r.Knots[idx].X), 0)
	case aoc.Abs(r.Knots[head].X-r.Knots[idx].X) > 1 || aoc.Abs(r.Knots[head].Y-r.Knots[idx].Y) > 1:
		r.Knots[idx] = r.Knots[idx].AddXY(aoc.Sign(r.Knots[head].X-r.Knots[idx].X), aoc.Sign(r.Knots[head].Y-r.Knots[idx].Y))
	}

	if idx == len(r.Knots)-1 {
		r.TailVisited.Add(r.Knots[idx])
	}
}

func (r *Rope) MoveHead(steps int, dx, dy int) {
	for i := 0; i < steps; i++ {
		r.Knots[0] = r.Knots[0].AddXY(dx, dy)
		for idx := 1; idx < len(r.Knots); idx++ {
			r.MoveIfNeeded(idx)
		}
	}
}

func Day09(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	r := NewRope(2)
	r2 := NewRope(10)
	for reader.Scan() {
		l := reader.Text()
		s := strings.Split(l, " ")

		switch s[0] {
		case "U":
			r.MoveHead(aoc.ToInt(s[1]), 0, -1)
			r2.MoveHead(aoc.ToInt(s[1]), 0, -1)
		case "D":
			r.MoveHead(aoc.ToInt(s[1]), 0, 1)
			r2.MoveHead(aoc.ToInt(s[1]), 0, 1)
		case "L":
			r.MoveHead(aoc.ToInt(s[1]), -1, 0)
			r2.MoveHead(aoc.ToInt(s[1]), -1, 0)
		case "R":
			r.MoveHead(aoc.ToInt(s[1]), 1, 0)
			r2.MoveHead(aoc.ToInt(s[1]), 1, 0)
		}

		part1 = len(r.TailVisited)
		part2 = len(r2.TailVisited)
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
