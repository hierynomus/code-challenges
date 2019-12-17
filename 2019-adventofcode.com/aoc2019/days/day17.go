package days

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

type Day17 struct{}

func (d *Day17) Solve(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("Boom!"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	grid := [][]rune{}
	grid = append(grid, []rune{})

	go icm.Run()

	y := 0
	for !icm.Closed {
		c := rune(<-icm.IO.Output)
		switch c {
		case '\n':
			y += 1
			grid = append(grid, []rune{})
		default:
			grid[y] = append(grid[y], c)
		}
	}

	intersections := FindIntersections(grid)
	fmt.Print(intersections)
	Print(grid)

	sum := 0
	for _, i := range intersections {
		sum += i.X * i.Y
	}
	return strconv.Itoa(sum), ""
}

func FindIntersections(grid [][]rune) []aoc.Point {
	points := []aoc.Point{}
	xLen := len(grid[0])
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < xLen-1; x++ {
			if grid[y][x] == '#' {
				found := true
				for _, n := range getNeighbours(x, y) {
					if grid[n.Y][n.X] != '#' {
						found = false
						break
					}
				}
				if found {
					points = append(points, aoc.Point{X: x, Y: y})
				}
			}
		}
	}
	return points
}

func getNeighbours(x, y int) []aoc.Point {
	return []aoc.Point{
		aoc.Point{X: x - 1, Y: y},
		aoc.Point{X: x + 1, Y: y},
		aoc.Point{X: x, Y: y - 1},
		aoc.Point{X: x, Y: y + 1},
	}
}

func Print(grid [][]rune) {
	for y := 0; y < len(grid); y++ {
		line := grid[y]
		for x := 0; x < len(line); x++ {
			fmt.Printf("%c", line[x])
		}
		fmt.Print("\n")
	}
}
