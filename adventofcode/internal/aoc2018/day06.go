package aoc2018

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day06(reader *bufio.Scanner) (string, string) {
	i := 0
	points := []aoc.Point{}
	for reader.Scan() {
		c := reader.Text()
		xy := strings.Split(c, ",")
		p := aoc.NewPoint(aoc.ToInt(strings.TrimSpace(xy[0])), aoc.ToInt(strings.TrimSpace(xy[1])))
		points = append(points, p)
		i += 1
	}

	part1 := d6p1(points)
	part2 := d6p2(points)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func d6p1(list []aoc.Point) int {
	grid := map[aoc.Point]int{}
	for i, p := range list {
		grid[p] = i
	}

	frontier := grid
	for i := 0; i < 400; i++ {
		newFrontier := map[aoc.Point]int{}
		for k, v := range frontier {
			if v != -1 {
				for _, p := range k.Neighbours4() {
					if _, ok := grid[p]; !ok {
						_, inFrontier := newFrontier[p]
						if inFrontier && newFrontier[p] != v {
							newFrontier[p] = -1
						} else {
							newFrontier[p] = v
						}
					}
				}
			}
		}
		frontier = newFrontier
		for k, v := range frontier {
			grid[k] = v
		}
	}

	boundary := aoc.IntSet{}
	for _, v := range frontier {
		boundary.Add(v)
	}

	h := aoc.IntHistogram{}
	for _, v := range grid {
		if !boundary.Contains(v) {
			h.Add(v)
		}
	}

	_, sz := h.Max()
	return sz
}

func d6p2(points []aoc.Point) int {
	sz := 0
	for x := 0; x < 320; x++ {
		for y := 0; y < 800; y++ {
			f := aoc.NewPoint(x, y)
			manhattanSum := 0
			for _, p := range points {
				manhattanSum += aoc.Manhattan(f, p)
			}
			if manhattanSum < 10000 {
				sz += 1
			}
		}
	}

	return sz
}
