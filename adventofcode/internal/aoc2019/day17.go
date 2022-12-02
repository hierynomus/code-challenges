package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2019/intcode"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

//nolint:funlen
func Day17(scanner *bufio.Scanner) (string, string) {
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
		case <-icm.ClosedCh:
			break loop
		default:
			i := icm.Output.Read()
			c := rune(i)
			switch c {
			case '\n':
				grid = append(grid, []rune{})
				y++
			default:
				grid[y] = append(grid[y], c)
			}
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
			droid.Input.Write(int(m))
		}
	}()

	Dust := 0
loop2:
	for {
		select {
		case <-droid.ClosedCh:
			break loop2
		default:
			i := droid.Output.Read()
			Dust = i
		}
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
		{X: x - 1, Y: y},
		{X: x + 1, Y: y},
		{X: x, Y: y - 1},
		{X: x, Y: y + 1},
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
