package aoc2022

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var SandStart = aoc.Point{X: 500, Y: 0}

func Day14(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	cave := aoc.MakeRuneGrid(700, 180, '.')

	maxY := 0
	for reader.Scan() {
		line := reader.Text()
		s := strings.Split(line, " -> ")
		p := aoc.ReadPoint(s[0])
		if p.Y > maxY {
			maxY = p.Y
		}

		for i := 1; i < len(s); i++ {
			q := aoc.ReadPoint(s[i])
			if q.Y > maxY {
				maxY = q.Y
			}

			if p.X == q.X {
				for _, y := range aoc.RangeIncl(p.Y, q.Y) {
					cave[y][p.X] = '#'
				}
			} else {
				for _, x := range aoc.RangeIncl(p.X, q.X) {
					cave[p.Y][x] = '#'
				}
			}
			p = q
		}
	}

	for x := 0; x < len(cave[0]); x++ {
		cave[maxY+2][x] = '#'
	}

	sandAtRest := 0
	sand := SandStart
	for {
		switch {
		case cave[sand.Y+1][sand.X] == '.':
			sand = sand.AddXY(0, 1)
		case cave[sand.Y+1][sand.X-1] == '.':
			sand = sand.AddXY(-1, 1)
		case cave[sand.Y+1][sand.X+1] == '.':
			sand = sand.AddXY(1, 1)
		case sand.Y == SandStart.Y && sand.X == SandStart.X:
			part2 = sandAtRest + 1
			return strconv.Itoa(part1), strconv.Itoa(part2)
		default:
			if sand.Y == maxY+1 && part1 == 0 {
				part1 = sandAtRest
			}
			cave[sand.Y][sand.X] = 'o'
			sandAtRest++
			sand = SandStart
		}
	}
}
