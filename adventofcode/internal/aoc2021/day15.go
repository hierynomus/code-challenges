package aoc2021

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type CavePath struct {
	risk     int
	location aoc.Point
}

func (c CavePath) Less(o aoc.QueueItem) bool {
	return c.risk < o.(CavePath).risk
}

func Day15(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadIntGrid(reader, "")
	start := aoc.NewPoint(0, 0)
	end := aoc.NewPoint(len(grid[0])-1, len(grid)-1)

	part1 = SafestCavePath(grid, start, end)

	fullGrid := make([][]int, 0)
	for y := 0; y < len(grid)*5; y++ {
		row := make([]int, len(grid[0])*5)
		for x := 0; x < len(grid[0])*5; x++ {
			row[x] = 1 + (grid[y%len(grid)][x%len(grid[0])]+1*(x/len(grid[0])+y/len(grid))-1)%9
		}
		fullGrid = append(fullGrid, row)
	}

	newEnd := aoc.NewPoint(len(fullGrid[0])-1, len(fullGrid)-1)
	part2 = SafestCavePath(fullGrid, start, newEnd)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func SafestCavePath(grid [][]int, start, end aoc.Point) int {
	pq := aoc.NewPriorityQueue()
	lowestRisk := map[aoc.Point]int{}
	pq.Push(CavePath{0, start})

	for pq.Length() > 0 {
		head := pq.Pop().(CavePath)
		if head.location.X == end.X && head.location.Y == end.Y {
			return head.risk
		}

		for _, p := range head.location.Neighbours4() {
			if p.Y < 0 || p.Y >= len(grid) || p.X < 0 || p.X >= len(grid[0]) {
				continue
			}

			newRisk := head.risk + grid[p.Y][p.X]
			if _, ok := lowestRisk[p]; !ok || newRisk < lowestRisk[p] {
				lowestRisk[p] = newRisk
				pq.Push(CavePath{newRisk, p})
			}
		}
	}
	return 0
}
