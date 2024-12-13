package aoc2024

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func paintRegion(garden map[aoc.Point]rune, p aoc.Point, region, painted aoc.PointSet) {
	region.Add(p)
	painted.Add(p)
	for _, n := range p.Neighbours4() {
		if garden[n] == garden[p] && !painted.Contains(n) {
			paintRegion(garden, n, region, painted)
		}
	}
}

func countCorners(region aoc.PointSet) int {
	corners := 0
	for p := range region {
		for _, n := range []aoc.Point{aoc.NewPoint(-1, -1), aoc.NewPoint(-1, 1), aoc.NewPoint(1, -1), aoc.NewPoint(1, 1)} {
			// Concave corner
			// .X?    ?X.    ???     ???
			// XX? or ?XX or ?XX. or XX?
			// ???    ???    ?X.     .X?
			if region.Contains(p.AddXY(n.X, 0)) && region.Contains(p.AddXY(0, n.Y)) && !region.Contains(p.Add(n)) {
				corners++
			}

			// Convex corner
			// ..?    ?..    ???    ???
			// .X? or ?X. or ?X. or .X?
			// ???   ???     ?..    ..?
			if !region.Contains(p.AddXY(n.X, 0)) && !region.Contains(p.AddXY(0, n.Y)) {
				corners++
			}
		}
	}

	return corners
}

func Day12(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	inp := aoc.ReadRuneGrid(reader)
	garden := make(map[aoc.Point]rune)
	for y, row := range inp {
		for x, plot := range row {
			garden[aoc.Point{X: x, Y: y}] = plot
		}
	}

	painted := aoc.PointSet{}

	regions := []aoc.PointSet{}
	for p := range garden {
		if !painted.Contains(p) {
			region := aoc.PointSet{}
			paintRegion(garden, p, region, painted)
			regions = append(regions, region)
		}
	}

	for _, region := range regions {
		part1 += len(region) * aoc.Perimeter(region)
		part2 += len(region) * countCorners(region)
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
