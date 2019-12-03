package days

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
)

type Wire []aoc.Point

type Day03 struct{}

func (d *Day03) Solve(reader *bufio.Scanner) (string, string) {
	var wires []Wire
	for reader.Scan() {
		wireDirections, _ := aoc.AsStringArray(reader.Text())
		wire := wireToPoints(wireDirections)
		wires = append(wires, wire)
	}

	intersects := wires[0].Intersections(wires[1])

	sort.Slice(intersects, aoc.ManhattanSort(intersects))

	part1 := fmt.Sprintf("%d", aoc.Manhattan(aoc.Origin, intersects[0]))

	sort.Slice(intersects, func(i, j int) bool {
		i1, i2 := intersects[i], intersects[j]
		d1 := wires[0].IndexOf(i1) + wires[1].IndexOf(i1)
		d2 := wires[0].IndexOf(i2) + wires[1].IndexOf(i2)
		return d1 < d2
	})

	part2 := fmt.Sprintf("%d", wires[0].IndexOf(intersects[0])+wires[1].IndexOf(intersects[0])+2)
	return part1, part2
}

func (w Wire) Intersections(o Wire) []aoc.Point {
	pointMap := map[aoc.Point]bool{}
	for _, p := range w {
		pointMap[p] = true
	}

	var intersects []aoc.Point
	for _, p := range o {
		if pointMap[p] {
			intersects = append(intersects, p)
		}
	}
	return intersects
}

func (w Wire) IndexOf(p aoc.Point) int {
	for i, x := range w {
		if x == p {
			return i
		}
	}
	return -1
}

func wireToPoints(wire []string) Wire {
	x, y := 0, 0
	var points []aoc.Point
	for _, move := range wire {
		dir := move[0]
		dist, err := strconv.Atoi(move[1:])
		if err != nil {
			panic(err)
		}
		var dx, dy int
		switch dir {
		case 'U':
			dx, dy = 0, 1
		case 'D':
			dx, dy = 0, -1
		case 'L':
			dx, dy = -1, 0
		case 'R':
			dx, dy = 1, 0
		}
		nx, ny, pts := makePoints(x, y, dx, dy, dist)
		x, y = nx, ny
		points = append(points, pts...)
	}
	return points
}

func makePoints(x, y, dx, dy, nr int) (nx, ny int, np []aoc.Point) {
	nx, ny = x, y
	for i := 0; i < nr; i++ {
		nx, ny = nx+dx, ny+dy
		np = append(np, aoc.Point{nx, ny})
	}
	return
}
