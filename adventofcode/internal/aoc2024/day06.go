package aoc2024

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Guard struct {
	Visited   map[aoc.Point]aoc.Point
	Position  aoc.Point
	Direction aoc.Point
}

func (g *Guard) Clone() *Guard {
	visited := map[aoc.Point]aoc.Point{}
	for k, v := range g.Visited {
		visited[k] = v
	}

	return &Guard{Position: g.Position, Visited: visited, Direction: g.Direction}
}

func (g *Guard) Move(grid [][]rune) bool {
	g.Visited[g.Position] = g.Direction
	newPos := g.Position.Add(g.Direction)
	if newPos.X < 0 || newPos.X >= len(grid[0]) || newPos.Y < 0 || newPos.Y >= len(grid) {
		return false
	} else if grid[newPos.Y][newPos.X] == '#' {
		g.Direction = g.Direction.RotateRight()
		return true
	}

	g.Position = newPos
	return true
}

func (g *Guard) FindLoop(grid [][]rune) bool {
	for {
		m := g.Move(grid)
		if !m {
			return false
		}

		if d, ok := g.Visited[g.Position]; ok {
			if d == g.Direction {
				return true
			}
		}
	}
}

func Day06(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadRuneGrid(reader)
	var guard *Guard
	for r, row := range grid {
		for i, c := range row {
			if c == '^' || c == 'v' || c == '<' || c == '>' {
				guard = &Guard{Position: aoc.Point{X: i, Y: r}, Visited: map[aoc.Point]aoc.Point{}}
				switch c {
				case '^':
					guard.Direction = aoc.North
				case 'v':
					guard.Direction = aoc.South
				case '<':
					guard.Direction = aoc.West
				case '>':
					guard.Direction = aoc.East
				}

				break
			}
		}
	}

	clone := guard.Clone()
	moved := true
	for moved {
		moved = clone.Move(grid)
	}
	part1 = len(clone.Visited)

	// Part 2
	for {
		nextPos := guard.Position.Add(guard.Direction)
		if nextPos.X < 0 || nextPos.X >= len(grid[0]) || nextPos.Y < 0 || nextPos.Y >= len(grid) {
			break
		}
		if grid[nextPos.Y][nextPos.X] == '.' {
			grid[nextPos.Y][nextPos.X] = '#'
			c := guard.Clone()
			if c.FindLoop(grid) {
				part2++
			}
			grid[nextPos.Y][nextPos.X] = '.'
		}
		if !guard.Move(grid) {
			break
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
