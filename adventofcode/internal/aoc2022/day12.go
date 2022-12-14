package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type HillPath struct {
	Position aoc.Point
	Length   int
}

func (h HillPath) Less(item aoc.QueueItem) bool {
	return h.Length < item.(HillPath).Length
}

func FindShortestPath(elevations [][]rune, start aoc.Point, goal func(hp HillPath) bool, validClimb func(elevation, other rune) bool) int {
	pq := aoc.NewPriorityQueue()
	pq.Push(HillPath{start, 0})
	visited := map[aoc.Point]bool{}
	for pq.Length() > 0 {
		head := pq.Pop().(HillPath)
		elevation := elevations[head.Position.Y][head.Position.X]

		if goal(head) {
			return head.Length
		}

		for _, p := range head.Position.Neighbours4() {
			if p.Y < 0 || p.Y >= len(elevations) || p.X < 0 || p.X >= len(elevations[0]) {
				continue
			}

			otherElevation := elevations[p.Y][p.X]

			if !validClimb(elevation, otherElevation) {
				continue
			}

			if _, ok := visited[p]; !ok {
				visited[p] = true
				pq.Push(HillPath{p, head.Length + 1})
			}
		}
	}

	return -1
}

func Day12(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	elevations := aoc.ReadRuneGrid(reader)
	var start, end aoc.Point
	for y := 0; y < len(elevations); y++ {
		for x := 0; x < len(elevations[y]); x++ {
			if elevations[y][x] == 'S' {
				start = aoc.NewPoint(x, y)
				elevations[y][x] = 'a' // Start-point is at 'a'
			} else if elevations[y][x] == 'E' {
				end = aoc.NewPoint(x, y)
				elevations[y][x] = 'z' // End-point is at 'z'
			}
		}
	}

	// Part 1
	part1 = FindShortestPath(elevations, start, func(hp HillPath) bool {
		return hp.Position.X == end.X && hp.Position.Y == end.Y
	}, func(elevation, other rune) bool {
		return other-elevation <= 1
	})

	// Find the shortest path from the end to any lowest elevation
	part2 = FindShortestPath(elevations, end, func(hp HillPath) bool {
		return elevations[hp.Position.Y][hp.Position.X] == 'a'
	}, func(elevation, other rune) bool {
		return elevation-other <= 1
	})

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
