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
		panic(fmt.Errorf("boom"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	grid := [][]rune{}
	grid = append(grid, []rune{})

	go icm.Run()

	y := 0

loop:
	for {
		select {
		case i := <-icm.IO.Output:
			c := rune(i)
			switch c {
			case '\n':
				grid = append(grid, []rune{})
				y++
			default:
				grid[y] = append(grid[y], c)
			}
		case <-icm.ClosedCh:
			break loop
		}
	}

	intersections := FindIntersections(grid)

	sum := 0
	for _, i := range intersections {
		sum += i.X * i.Y
	}

	program[0] = 2
	droid := intcode.NewIntCodeMachine(program)

	go droid.Run()

	// Calculated by hand
	// 	Print(grid)
	//Main: A,B,A,C,B,C,B,C,A,C
	//A: L,10,R,12,R,12
	//B: R,6,R,10,L,10
	//C: R,10,L,10,L,12,R,6
	go func() {
		movements := "A,B,A,C,B,C,B,C,A,C\nL,10,R,12,R,12\nR,6,R,10,L,10\nR,10,L,10,L,12,R,6\nn\n"
		for _, m := range movements {
			droid.IO.Input <- int(m)
		}
	}()

	Dust := 0
	for i := range droid.IO.Output {
		Dust = i
	}

	return strconv.Itoa(sum), strconv.Itoa(Dust)
}

func FindIntersections(grid [][]rune) []aoc.Point {
	points := []aoc.Point{}

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
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
		aoc.Point{X: x - 1, Y: y}, //nolint:gofmt
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
